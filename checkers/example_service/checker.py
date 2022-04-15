#!/usr/bin/env python3.9
from gornilo import CheckRequest, Verdict, Checker, PutRequest, GetRequest

checker = Checker()

@checker.define_check
def check_service(request: CheckRequest) -> Verdict:
    return Verdict.OK()


@checker.define_put(vuln_num=1, vuln_rate=1)
def put_flag(request: PutRequest) -> Verdict:
    return Verdict.OK()


@checker.define_get(vuln_num=1)
def get_flag(request: GetRequest) -> Verdict:
    return Verdict.OK()

if __name__ == '__main__':
    checker.run()
