#!/usr/bin/env python3

import json
import random
import secrets
from typing import Any, Tuple, Sequence, Coroutine

import api
import gornilo


port = 17171
checker = gornilo.NewChecker()


def generate_username() -> str:
    return secrets.token_hex(8)


def generate_disease_type() -> str:
    return secrets.token_hex(4)


def generate_disease_phase() -> str:
    return secrets.token_urlsafe(8)


def generate_disease_symptoms() -> Sequence[str]:
    length = random.randint(5, 10)
    symptoms = [secrets.token_hex(3) for _ in range(length)]

    return symptoms


def generate_disease(name: str) -> Tuple[str, Sequence[str]]:
    type = random.choice(['mental', 'infectious', 'other'])

    if type == 'mental':
        phase = generate_disease_phase()
        serialized = f'{name} (mental), {phase} phase'

        return serialized, (type, name, phase)
    elif type == 'infectious':
        symptoms = generate_disease_symptoms()
        serialized = f'{name} (infectious); symptoms: {", ".join(symptoms)}'

        return serialized, (type, name, ' '.join(symptoms))
    else:
        type = generate_disease_type()
        serialized = f'{name} ({type})'

        return serialized, (type, name)


async def wrap_exceptions(
        action: Coroutine[gornilo.CheckRequest, Any, gornilo.Verdict],
        request: gornilo.CheckRequest,
) -> gornilo.Verdict:
    try:
        return await action(request)
    except api.ProtocolError as e:
        return gornilo.Verdict.MUMBLE(f'protocol error: {e}')
    except EOFError:
        return gornilo.Verdict.MUMBLE('received EOF')
    except TimeoutError:
        return gornilo.Verdict.DOWN('timeout error')
    except ConnectionError:
        return gornilo.Verdict.DOWN('connection error')


async def do_put(request: gornilo.PutRequest) -> gornilo.Verdict:
    async with api.Ambulance.connect(request.hostname, port) as io:
        await io.read_banner()

        username = generate_username()
        response, (password, _) = await io.register(username)

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to register')

        disease, params = generate_disease(request.flag)
        response = await io.update_disease(*params)

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to update disease')

        response = await io.user_exit()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to exit')

        return gornilo.Verdict.OK_WITH_FLAG_ID(
            username,
            json.dumps([password, disease]),
        )


async def do_get(request: gornilo.GetRequest) -> gornilo.Verdict:
    async with api.Ambulance.connect(request.hostname, port) as io:
        await io.read_banner()

        username = request.public_flag_id
        password, expected_disease = json.loads(request.flag_id)
        response = await io.login(username, password)

        if response is not api.Response.OK:
            return gornilo.Verdict.CORRUPT('failed to login')

        response, (name, disease) = await io.print_info()

        if response is not api.Response.OK:
            return gornilo.Verdict.CORRUPT('failed to print info')

        if name != username or disease != expected_disease:
            return gornilo.Verdict.CORRUPT('invalid info')

        return gornilo.Verdict.OK()


async def do_check(request: gornilo.CheckRequest) -> gornilo.Verdict:
    async with api.Ambulance.connect(request.hostname, port) as io:
        await io.read_banner()

        return gornilo.Verdict.OK()


@checker.define_vuln('flag_id is username')
class AmbulanceChecker(gornilo.VulnChecker):
    @staticmethod
    async def put(request: gornilo.PutRequest) -> gornilo.Verdict:
        return await wrap_exceptions(do_put, request)

    @staticmethod
    async def get(request: gornilo.GetRequest) -> gornilo.Verdict:
        return await wrap_exceptions(do_get, request)

    @checker.define_check
    async def check(request: gornilo.CheckRequest) -> gornilo.Verdict:
        return await wrap_exceptions(do_check, request)


if __name__ == '__main__':
    checker.run()
