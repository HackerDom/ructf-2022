#!/usr/bin/env python3

import requests

from gornilo import CheckRequest, Verdict, Checker, PutRequest, GetRequest
from gornilo.models.verdict.verdict_codes import *

from logging import getLogger

print = getLogger().info

checker = Checker()


@checker.define_check
def check_service(request: CheckRequest) -> Verdict:
    return Verdict.DOWN('not implemented')


@checker.define_get(vuln_num=1)
def get_flag(request: GetRequest) -> Verdict:
    return Verdict.DOWN('not implemented')


@checker.define_put(vuln_num=1, vuln_rate=1)
def put_flag(request: PutRequest) -> Verdict:
    return Verdict.DOWN('not implemented')


if __name__ == '__main__':
    checker.run()
