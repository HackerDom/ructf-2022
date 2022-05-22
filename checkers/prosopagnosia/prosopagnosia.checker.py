#!/usr/bin/env python3

import uuid
import random
import string
import base64
import os

from gornilo import CheckRequest, Verdict, PutRequest, GetRequest, VulnChecker, NewChecker
from gornilo.http_clients import requests_with_retries

from logging import getLogger

PORT = 15345
NAME_HEADER = 'X-Svm-Name'
KEY_HEADER = 'X-Svm-Key'
AUTHOR_HEADER = 'X-Svm-Author'
SECRET_HEADER = 'X-Svm-Secret'

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
    dir = os.path.dirname(os.path.realpath(__file__))
    filename = random.choice(os.listdir(os.path.join(dir, 'roms')))

    with open(os.path.join(dir, 'roms', filename), 'rb') as f:
        return f.read()


@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    try:
        with requests_with_retries().get(
                f'http://{request.hostname}:{PORT}/api/demo/list?page_num=0&page_size=100') as r:
            if r.status_code != 200:
                return Verdict.MUMBLE('service mumble')

        return Verdict.OK()

    except Exception as e:
        print(e)
        return Verdict.MUMBLE('service down')


@checker.define_vuln('flags container')
class Base64Vuln(VulnChecker):
    @staticmethod
    def put(req: PutRequest) -> Verdict:
        try:
            name = base64.b64encode(('demo-' + str(uuid.uuid4())).encode()).decode()
            author = base64.b64encode(str(uuid.uuid4()).encode()).decode()
            secret = base64.b64encode(req.flag.encode()).decode()
            demo = select_random_demo()

            print(f'{name}, {author}, {secret}')

            with requests_with_retries().post(f'http://{req.hostname}:{PORT}/api/demo',
                                              headers={NAME_HEADER: name, SECRET_HEADER: secret, AUTHOR_HEADER: author},
                                              files={name: demo}) as r:
                if r.status_code != 200:
                    return Verdict.MUMBLE('wrong http code')

                key = r.headers.get(KEY_HEADER)

                print(f'name {name} and key {key}')

                if key == '' or key is None:
                    return Verdict.MUMBLE('wrong answer')

            with requests_with_retries().get(f'http://{req.hostname}:{PORT}/api/demo',
                                             headers={KEY_HEADER: key, NAME_HEADER: name}) as r:
                if r.status_code != 200:
                    return Verdict.MUMBLE('wrong http code')

                print(r.content)

                j = r.json()

                if j['secret'] != secret or j['author'] != author or j['name'] != name:
                    print(f'{r.content} {name} {author} {secret}')
                    return Verdict.CORRUPT('wrong flag')

                rom_path = j['rom_path']

            with requests_with_retries().get(f'http://{req.hostname}:{PORT}/{rom_path}') as r:
                if r.status_code != 200:
                    return Verdict.MUMBLE('wrong http code')

                if r.content != demo:
                    print(demo)
                    print(r.content)
                    return Verdict.CORRUPT('wrong rom')

            return Verdict.OK_WITH_FLAG_ID(name, key)
        except Exception as e:
            print(e)
            return Verdict.MUMBLE('service error')

    @staticmethod
    def get(req: GetRequest) -> Verdict:
        try:
            name = req.public_flag_id
            key = req.flag_id

            with requests_with_retries().get(f'http://{req.hostname}:{PORT}/api/demo',
                                             headers={KEY_HEADER: key, NAME_HEADER: name}) as r:
                if r.status_code == 401:
                    return Verdict.CORRUPT('wrong flag')

                if r.status_code != 200:
                    return Verdict.MUMBLE('wrong http code')

                j = r.json()

                if j['secret'] != base64.b64encode(req.flag.encode()).decode() or j['name'] != name:
                    print(f'differs: {j["secret"]} {base64.b64encode(req.flag.encode()).decode()} {req.flag}')
                    return Verdict.CORRUPT('wrong flag')

            return Verdict.OK()
        except Exception as e:
            print(e)
            return Verdict.MUMBLE('service error')


if __name__ == '__main__':
    checker.run()
