#!/usr/bin/env python3.8
import sys
import traceback
import pg8000
import uuid

import requests
from requests.exceptions import Timeout
from requests.adapters import HTTPAdapter
from urllib3.util import Retry
from gornilo import CheckRequest, Verdict, PutRequest, GetRequest, VulnChecker, NewChecker
from gornilo.http_clients import requests_with_retries

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
    with PgErrorChecker() as registry_ec:
        with PgConn(request.hostname) as pgc:
            cursor = pgc.conn.cursor()
            cursor.execute("SELECT 1")
            pgc.conn.commit()

    return registry_ec.verdict


@checker.define_vuln("flag_id is a metadata")
class OracleChecker(VulnChecker):
    @staticmethod
    def put(request: PutRequest) -> Verdict:
        # with ErrorChecker() as doctor_ec:
        #     # check doctor is alive
        #     url = f"http://{request.hostname}:{DOCTOR_PORT}/api/v1/jobs"
        #     payload = {
        #         'id': '4C3443283C07611254AECD1FDDFA99452AF996EC5819784B69C053B66F4F7F5EA22B7E27B33F7C93C492421D21DBD29FCB0819B4B1AFD9D2548CC831C47C57FCBBFF1CE70E3468E692B7355C02C85EB7812F93F7A6A0905C6730904E2CE27F34A8A018DA1025968453BC73047BC43EC90B7FFE323060CA3D5E1A28D9BB8AA337'}
        #     files = []
        #     headers = {}
        #
        #     resp = requests_with_retries().put(
        #         url,
        #         headers=headers,
        #         data=payload
        #     )
        #     if resp is None:
        #         return Verdict.CORRUPT("corrupt response from doctor")
        #
        #     resp_json = resp.json()
        #     if "data" not in resp_json:
        #         doctor_ec.verdict = Verdict.CORRUPT("corrupt response from doctor")
        #     else:
        #         doctor_ec.verdict = Verdict.OK()
        #
        # with PgErrorChecker() as registry_ec:
        #     conn = pg8000.connect("svcuser", host=request.hostname, password="svcpass", database="postgres",
        #                           port=PG_PORT)
        #     cursor = conn.cursor()
        #     cursor.execute("SELECT add_job('%s')" % str(uuid.uuid4()))
        #     conn.commit()
        #     conn.close()
        #
        # if registry_ec.verdict._code > doctor_ec.verdict._code:
        #     return registry_ec.verdict
        #
        # return doctor_ec.verdict

        return Verdict.OK()

    @staticmethod
    def get(request: GetRequest) -> Verdict:
        return Verdict.OK()


if __name__ == '__main__':
    checker.run()
