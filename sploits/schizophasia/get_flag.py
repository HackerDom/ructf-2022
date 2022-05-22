#!/usr/bin/env python3.8

import hashlib
import sys
import traceback

import pg8000

PG_PORT = 5432
DOCTOR_PORT = 18181

class PgErrorChecker:
    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
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


def get_flag(hostname, uid, hash_str):
    with PgErrorChecker() as registry_ec, PgConn(hostname) as pgc:
        cursor = pgc.conn.cursor()
        print("SELECT * FROM medical_history('%s', '%s')\n" % (uid, hash_str))
        cursor.execute("SELECT * FROM medical_history('%s', '%s')" % (uid, hash_str))
        rows = cursor.fetchall()

        if not rows or len(rows) != 1 or len(rows[0]) != 4:
            raise Exception("unexpected registry service response format")

        question, user, status, response = rows[0]
        print("Q: %s, U: %s, S: %s, R: %s\n" % (question, user, status, response))
        cursor.close()

def main(hostname, uid, token):
    sha_1 = hashlib.sha1()
    sha_1.update((uid + token).encode('utf-8'))
    hash_str = sha_1.hexdigest().upper()
    print("SHA1: %s\n", hash_str)
    print(get_flag(hostname, uid, hash_str))


if __name__ == '__main__':
    if len(sys.argv) != 4:
        print("usage: get_flag.py hostname uid token")
        exit(-1)
    main(sys.argv[1], sys.argv[2], sys.argv[3])
