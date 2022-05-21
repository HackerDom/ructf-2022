#!/usr/bin/env python3

import requests
import uuid
import logging
import random
import string

from gornilo import CheckRequest, Verdict, PutRequest, GetRequest, VulnChecker, NewChecker
from http.client import HTTPConnection

from gornilo.models.verdict.verdict_codes import *

from logging import getLogger

HTTPConnection.debuglevel = 1
logging.basicConfig(
    format='%(asctime)s - %(name)s - %(levelname)s - %(message)s'
)
logging.getLogger().setLevel(logging.DEBUG)
log = logging.getLogger("requests.packages.urllib3")
log.setLevel(logging.DEBUG)
log.propagate = True

PORT = 15345

print = getLogger().info

checker = NewChecker()

flag_alphabet = string.digits + string.ascii_uppercase


def generate_some_flag():
    # [0-9A-Z]{31}=
    flag = []
    for i in range(31):
        flag.append(random.choice(flag_alphabet))
    flag.append('=')

    return ''.join(flag)


def select_random_demo():
    pass


@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    try:
        with requests.get(f'http://{request.hostname}:{PORT}/api/demo/list?page_num=0&page_size=100') as r:
            if r.status_code != 200:
                return Verdict.MUMBLE('service mumble')

        return Verdict.OK()

    except Exception as e:
        log.error(e)
        return Verdict.DOWN('service down')


@checker.define_vuln('base64 corruption')
class Base64Vuln(VulnChecker):
    @staticmethod
    def put(r: PutRequest) -> Verdict:
        try:
            name = 'demo-' + str(uuid.uuid4())
            author = str(uuid.uuid4())
            secret = r.flag
            demo = select_random_demo()

            with requests.put(f'http://{r.hostname}:{PORT}/api/demos') as r:
                pass

            return Verdict.OK_WITH_FLAG_ID(name, key)
        except Exception as e:
            log.error(e)
            return Verdict.DOWN('service down')

    @staticmethod
    def get(r: GetRequest) -> Verdict:
        r.fla


if __name__ == '__main__':
    checker.run()
