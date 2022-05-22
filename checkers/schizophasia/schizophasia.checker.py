#!/usr/bin/env python3.8
import json
import os
import random
import sys
import traceback
from os.path import exists
from time import time

import gornilo.models.verdict
import pg8000
import uuid

import requests
from requests.exceptions import Timeout
from requests.adapters import HTTPAdapter
from urllib3.util import Retry
from gornilo import CheckRequest, Verdict, PutRequest, GetRequest, VulnChecker, NewChecker
from gornilo.http_clients import requests_with_retries
from gornilo.models.verdict import OK
checker = NewChecker()
PG_PORT = 5432
DOCTOR_PORT = 18181
GET_RETRIES_COUNT = 15
GET_RETRY_DELAY = 2

LOOK_BACK_WINDOW_MINUTES = 31

class ErrorChecker:
    def __init__(self):
        self.verdict = Verdict.OK()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
        if exc_type in {requests.exceptions.ConnectionError, ConnectionError, ConnectionAbortedError,
                        ConnectionRefusedError, ConnectionResetError}:
            self.verdict = Verdict.DOWN("Service is down")
        if exc_type in {requests.exceptions.HTTPError}:
            self.verdict = Verdict.MUMBLE(f"Incorrect http code")

        if exc_type:
            print(exc_type)
            print(exc_value.__dict__)
            traceback.print_tb(exc_traceback, file=sys.stdout)
        return True


class PgErrorChecker:
    def __init__(self):
        self.verdict = Verdict.OK()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
        if exc_type in {pg8000.ProgrammingError}:
            self.verdict = Verdict.DOWN("postgres request failure")
            print(exc_type)
            print(exc_value)
            traceback.print_tb(exc_traceback, file=sys.stdout)

        if exc_type:
            print(exc_type)
            print(exc_value)
            traceback.print_tb(exc_traceback, file=sys.stdout)
        return True


class PgConn:
    def __init__(self, hostname):
        self.hostname = hostname

    def __enter__(self):
        self.conn = pg8000.connect("svcuser", host=self.hostname, password="svcpass", database="postgres", port=PG_PORT)
        return self

    def __exit__(self, type, value, traceback):
        if self.conn:
            self.conn.close()


def get_pg_conn(hostname):
    return


def check_cant_read_old(cursor):
    state = read_state()
    if state is None or len(state) == 0:
        cursor.close()
        return True

    sample = next((entry for entry in state if entry['ts'] < int(time()) - (LOOK_BACK_WINDOW_MINUTES * 60)), None)
    if sample is None:
        print("no good entries found\n")
        return True

    try:
        cursor.execute("SELECT * FROM medical_history('%s', '%s')" % (sample['public_flag_id'], sample['flag_id']))
        cursor.close()

    except Exception as e:
        print("[OK] can't read old data: %s" % e)
        # we want exception on reading an old data
        return True

    return False


def functions_defs_unchanged(cursor):
    check_funcs = {
        'public.add_job(text,text)': """CREATE OR REPLACE FUNCTION public.add_job(text, text)
 RETURNS text
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
BEGIN
IF (SELECT 1 FROM jobqueue WHERE "question"=$1 AND "userid"=$2) THEN
    RAISE 'DUPLICATE';
ELSE
    INSERT INTO jobqueue SELECT $1, $2, false, NULL, now(), t FROM load_token() t;
    RETURN (SELECT meta from create_meta($1, $2, load_token()) meta);
END IF;
END
$function$""",
        'public.medical_history(text, text)': """CREATE OR REPLACE FUNCTION public.medical_history(text, text)
 RETURNS TABLE(q text, u text, s boolean, r text)
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
BEGIN
    IF (SELECT verify($1, t.token, $2) AND authorize(t.token) FROM (SELECT token from jobqueue where userid=$1) t) THEN
        RETURN QUERY SELECT "question", "userid", "status", "result" from jobqueue where "userid" = $1;
    ELSE
        RAISE 'UNAUTHORIZED';
    END IF;
END
$function$""",
        'public.finish_job(text, text, text, text)': """CREATE OR REPLACE FUNCTION public.finish_job(text, text, text, text)
 RETURNS void
 LANGUAGE plpgsql
 SECURITY DEFINER
AS $function$
BEGIN
    IF (SELECT authorize($1)) THEN
        UPDATE jobqueue SET result = $4, status = TRUE WHERE question = $2 AND userid = $3;
ELSE
        RAISE 'UNAUTHORIZED';
END IF;
END
$function$""",
    }

    for func_name, func_def in check_funcs.items():
        if not check_func(cursor, func_name, func_def):
            print("fail")
            return False

    return True


def check_func(cursor, func_name, expected_def):
    cursor.execute("SELECT pg_get_functiondef('%s'::regprocedure);" % func_name)
    rows = cursor.fetchall()

    if not rows or len(rows) != 1 or len(rows[0]) != 1:
        return Verdict.CORRUPT("unexpected registry service add_job func definition")

    actual_def = rows[0][0].strip()

    if actual_def != expected_def:
        print("func:%s expected: %s, returned: %s\n" %(func_name, expected_def, actual_def))
        return False

    return True

@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    with PgErrorChecker() as registry_ec, PgConn(request.hostname) as pgc:
        cursor = pgc.conn.cursor()
        cursor.execute("SELECT 1")

        if not functions_defs_unchanged(cursor):
            return Verdict.CORRUPT("Function definitions changed")

        if not check_cant_read_old(cursor):
            return Verdict.CORRUPT("Can read old data")

    return registry_ec.verdict


def read_state():
    try:
        dir = os.path.dirname(os.path.realpath(__file__))
        state_path = os.path.join(dir, 'state.json')
        if not exists(state_path):
            return None

        with open(state_path, 'r') as sf:
            json_data = json.load(sf)
            return json_data

    except Exception:
        print("warning: failed to read check state file")
        return None


def write_state(data):
    try:
        dir = os.path.dirname(os.path.realpath(__file__))
        state_path = os.path.join(dir, 'state.json')

        with open(state_path, 'w+') as sf:
            json.dump(data, sf)

    except Exception:
        print("warning: failed to write check state file")
        return None


@checker.define_vuln("flag_id is a username")
class OracleChecker(VulnChecker):
    @staticmethod
    def put(request: PutRequest) -> Verdict:
        meta = None
        public_flag_id = "user_" + str(uuid.uuid4())
        with PgErrorChecker() as registry_ec, PgConn(request.hostname) as pgc:
            cursor = pgc.conn.cursor()
            cursor.execute("SELECT add_job('%s', '%s')" % (request.flag, public_flag_id))
            pgc.conn.commit()
            rows = cursor.fetchall()

            if not rows or len(rows) != 1 or len(rows[0]) != 1:
                return Verdict.MUMBLE("unexpected registry service response")

            meta = rows[0][0]
            cursor.close()

        if registry_ec.verdict._code != OK:
            return registry_ec.verdict

        priv_flag_id = None

        print("public flag id:%s, meta:%s\n" % (public_flag_id, meta))
        with ErrorChecker() as doctor_ec:
            url = f"http://{request.hostname}:{DOCTOR_PORT}/api/v1/jobs"
            payload = {'id': meta}
            files = []
            headers = {}

            resp = requests_with_retries().put(
                url,
                headers=headers,
                data=payload
            )
            if resp is None:
                return Verdict.CORRUPT("corrupt response from doctor")

            resp_json = resp.json()
            if "data" not in resp_json or not resp_json["data"]:
                return Verdict.CORRUPT("corrupt response from doctor")

            print("hash: %s\n" % resp_json["data"])
            priv_flag_id = resp_json["data"]

        return Verdict.OK_WITH_FLAG_ID(public_flag_id, priv_flag_id)

    @staticmethod
    def get(request: GetRequest) -> Verdict:
        with PgErrorChecker() as registry_ec, PgConn(request.hostname) as pgc:
            cursor = pgc.conn.cursor()
            print("SELECT * FROM medical_history('%s', '%s')\n" % (request.public_flag_id, request.flag_id))
            cursor.execute("SELECT * FROM medical_history('%s', '%s')" % (request.public_flag_id, request.flag_id))
            pgc.conn.commit()
            rows = cursor.fetchall()

            if not rows or len(rows) != 1 or len(rows[0]) != 4:
                return Verdict.MUMBLE("unexpected registry service response format")

            question, user, status, response = rows[0]
            print("Q: %s, U: %s, S: %s, R: %s\n" % (question, user, status, response))

            if question != request.flag or user != request.public_flag_id or status != True or len(response) == 0:
                return Verdict.MUMBLE("registry service response contents mismatch")

            cursor.close()
            try:
                if registry_ec.verdict._code == OK:
                    print("saving state...")
                    state = read_state()
                    if state is None or len(state) > 120:
                        state = []

                    state.append(
                        {'ts': int(time()), 'public_flag_id': request.public_flag_id, 'flag_id': request.flag_id})
                    write_state(state)
            except Exception as e:
                print("state update failed: %s\n", e)

            return Verdict.OK()

        return registry_ec.verdict


if __name__ == '__main__':
    checker.run()
