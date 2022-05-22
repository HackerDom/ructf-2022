#!/usr/bin/env python3.8

import sys
import traceback

import pg8000
import uuid

import requests


import argparse

PG_PORT = 5432
DOCTOR_PORT = 18181


class PgConn:
    def __init__(self, hostname):
        self.hostname = hostname

    def __enter__(self):
        self.conn = pg8000.connect("svcuser", host=self.hostname, password="svcpass", database="postgres", port=PG_PORT)
        return self

    def __exit__(self, type, value, traceback):
        if self.conn:
            self.conn.close()


class PgErrorChecker:
    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
        if exc_type in {pg8000.ProgrammingError}:
            if 'DUPLICATE' in str(exc_value):
                print("Question already exists. Try another one?")
                exit(0)

        if exc_type:
            print(exc_type)
            print(exc_value)
            traceback.print_tb(exc_traceback, file=sys.stdout)
            print("Registry request failed.")
            exit(0)


def main(question, hostname):
    user_id = "user_" + str(uuid.uuid4())

    with PgErrorChecker(), PgConn(hostname) as pgc:
        cursor = pgc.conn.cursor()
        cursor.execute("SELECT add_job('%s', '%s')" % (question, user_id))
        pgc.conn.commit()
        rows = cursor.fetchall()

        if not rows or len(rows) != 1 or len(rows[0]) != 1:
            raise Exception("unexpected registry service response")

        ticket = rows[0][0]
        cursor.close()

    url = f"http://{hostname}:{DOCTOR_PORT}/api/v1/jobs"
    payload = {'id': ticket}
    headers = {}

    resp = requests.put(
        url,
        headers=headers,
        data=payload
    )
    if resp is None:
        raise Exception("corrupt response from doctor")

    resp_json = resp.json()

    if "data" not in resp_json or not resp_json["data"]:
        raise Exception("corrupt response from doctor")

    with PgErrorChecker(), PgConn(hostname) as pgc:
        cursor = pgc.conn.cursor()
        cursor.execute("SELECT * FROM medical_history('%s', '%s')" % (user_id, resp_json["data"]))
        pgc.conn.commit()
        rows = cursor.fetchall()

        if not rows or len(rows) != 1 or len(rows[0]) != 4:
            raise Exception("unexpected registry service response format")

        question, user, status, response = rows[0]
        print("Response: %s\n" % response)


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Ask a question')
    parser.add_argument('question', metavar='Q', type=str, help='your question')
    parser.add_argument('hostname', metavar='H', type=str, help='hostname')

    args = parser.parse_args()
    main(args.question, args.hostname)
