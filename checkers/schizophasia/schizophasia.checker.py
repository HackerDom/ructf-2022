#!/usr/bin/env python3
import sys
import traceback

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
PG_CONN_STRING = "host=%s port=%d user=svcuser dbname=postgres password=svcpass"
DOCTOR_PORT = 18181
GET_RETRIES_COUNT = 15
GET_RETRY_DELAY = 2


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
        # if exc_type in {InterfaceError, DatabaseError, DataError, OperationalError, IntegrityError, InternalError, ProgrammingError, NotSupportedError}:
        #     self.verdict = Verdict.DOWN("psycopg2 error")

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


@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    with PgErrorChecker() as registry_ec, PgConn(request.hostname) as pgc:
        cursor = pgc.conn.cursor()
        cursor.execute("SELECT 1")
        cursor.close()

    return registry_ec.verdict


@checker.define_vuln("flag_id is a username")
class OracleChecker(VulnChecker):
    @staticmethod
    def put(request: PutRequest) -> Verdict:
        priv_flag_id = None
        public_flag_id = "user_" + str(uuid.uuid4())
        with PgErrorChecker() as registry_ec, PgConn(request.hostname) as pgc:
            cursor = pgc.conn.cursor()
            cursor.execute("SELECT add_job('%s', '%s')" % (public_flag_id, request.flag))
            pgc.conn.commit()
            rows = cursor.fetchall()

            if not rows or len(rows) != 1 or len(rows[0]) != 1:
                return Verdict.MUMBLE("unexpected registry service response")

            priv_flag_id = rows[0][0]
            cursor.close()

        if registry_ec.verdict._code != OK:
            return registry_ec.verdict

        print("public flag id:%s, priv flag id:%s\n" % (public_flag_id, priv_flag_id))
        with ErrorChecker() as doctor_ec:
            url = f"http://{request.hostname}:{DOCTOR_PORT}/api/v1/jobs"
            payload = {'id': priv_flag_id}
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
            if "data" not in resp_json:
                return Verdict.CORRUPT("corrupt response from doctor")

        return Verdict.OK_WITH_FLAG_ID(public_flag_id, priv_flag_id)

    @staticmethod
    def get(request: GetRequest) -> Verdict:
        return Verdict.OK()


if __name__ == '__main__':
    checker.run()
