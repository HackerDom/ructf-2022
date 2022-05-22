#!/usr/bin/env python3

import json
import random
import secrets
from typing import Any, Tuple, Sequence, Coroutine

import api
import utils
import crypto
import gornilo


port = 17171
checker = gornilo.NewChecker()


def generate_username() -> str:
    return secrets.token_hex(8)


def generate_disease_name() -> str:
    return secrets.token_hex(8)


def generate_disease_type() -> str:
    return secrets.token_hex(4)


def generate_disease_phase() -> str:
    return secrets.token_hex(4)


def generate_disease_symptoms() -> Sequence[str]:
    length = random.randint(3, 6)
    symptoms = [secrets.token_hex(3) for _ in range(length)]

    return symptoms


def validate_password(username: str, password: str, recovery_key: str) -> bool:
    try:
        rs = utils.deserialize_numbers_sequence(password)

        if len(rs) != 2:
            return False

        public_key = crypto.get_public_key(recovery_key)

        return crypto.verify(username.encode(), public_key, password)
    except (utils.SerializationError, crypto.CryptoError):
        return False


def validate_recovery_key(recovery_key: str) -> bool:
    try:
        ds = utils.deserialize_numbers_sequence(recovery_key)

        return len(ds) == 1
    except utils.SerializationError:
        return False


def generate_new_password(password: str) -> str:
    r, s = utils.deserialize_numbers_sequence(password)
    k1, k2 = random.randint(0, 1000), random.randint(0, 1000)
    r_new, s_new = r, k2 * crypto.SecureCurve.q + s

    return utils.serialize_numbers_sequence(r_new, s_new)


def generate_new_auth_pair(username: str) -> Tuple[str, str]:
    recovery_key, _ = crypto.generate_keypair()
    password = crypto.sign(username.encode(), recovery_key)

    return password, recovery_key


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
            return gornilo.Verdict.MUMBLE('failed to print info')

        if name != username or disease != expected_disease:
            return gornilo.Verdict.CORRUPT('invalid info')

        return gornilo.Verdict.OK()


async def do_check(request: gornilo.CheckRequest) -> gornilo.Verdict:
    async with api.Ambulance.connect(request.hostname, port) as io:
        await io.read_banner()

        username = generate_username()
        response, (password, recovery_key) = await io.register(username)

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to register')

        if not validate_recovery_key(recovery_key):
            return gornilo.Verdict.MUMBLE('invalid recovery key')

        if not validate_password(username, password, recovery_key):
            return gornilo.Verdict.MUMBLE('invalid password')

        response, (name, disease) = await io.print_info()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to print info')

        if name != username or disease is not None:
            return gornilo.Verdict.MUMBLE('invalid info')

        response = await io.logout()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to logout')

        response = await io.login(username, password)

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to login')

        expected_disease, params = generate_disease(generate_disease_name())
        response = await io.update_disease(*params)

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to update disease')

        response, (name, disease) = await io.print_info()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to print info')

        if name != username or disease != expected_disease:
            return gornilo.Verdict.MUMBLE('invalid info')

        response = await io.user_exit()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to exit')

    async with api.Ambulance.connect(request.hostname, port) as io:
        await io.read_banner()

        response = await io.login(username, generate_new_password(password))

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to login')

        new_password, new_recovery_key = generate_new_auth_pair(username)
        response = await io.change_recovery_key(
            generate_new_password(password), new_recovery_key,
        )

        response, (name, disease) = await io.print_info()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to print info')

        if name != username or disease != expected_disease:
            return gornilo.Verdict.MUMBLE('invalid info')

        response = await io.logout()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to logout')

        response = await io.login(username, new_password)

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to login')

        response = await io.user_exit()

        if response is not api.Response.OK:
            return gornilo.Verdict.MUMBLE('failed to exit')

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
