#!/usr/bin/env python3

import enum
import asyncio
import contextlib

from typing import Tuple, AsyncGenerator


class Response(enum.Enum):
    OK = enum.auto()
    BAD_USERNAME = enum.auto()
    BAD_PASSWORD = enum.auto()
    BAD_RECOVERY_KEY = enum.auto()


class ProtocolError(Exception):
    pass


class Ambulance:
    def __init__(self, reader: asyncio.StreamReader, writer: asyncio.StreamWriter) -> None:
        self._reader = reader
        self._writer = writer

    @contextlib.asynccontextmanager
    @staticmethod
    async def connect(host: str, port: int) -> AsyncGenerator['Ambulance', None]:
        reader, writer = await asyncio.open_connection(host, port)

        try:
            yield Ambulance(reader, writer)
        finally:
            writer.close()

    async def readline(self) -> bytes:
        return await self._reader.readline()

    async def writeline(self, data: bytes) -> None:
        self._writer.write(data + b'\n')
        await self._writer.drain()

    async def read_banner(self) -> None:
        with open('banner.txt', 'rb') as file:
            for line in file:
                if line != await self.readline():
                    raise ProtocolError('invalid banner')

        line = b'=== Ambulance Database ===\n'
        if line != await self.readline():
            raise ProtocolError('invalid banner')

    async def read_anonymous_menu(self) -> None:
        lines = [
            b'\n',
            b'1) login\n',
            b'2) register\n',
            b'3) exit\n',
        ]

        for line in lines:
            if line != await self.readline():
                raise ProtocolError('invalid menu')

    async def read_user_menu(self) -> None:
        lines = [
            b'\n',
            b'1) print info\n',
            b'2) change recovery key\n',
            b'3) update disease\n',
            b'4) logout\n',
            b'5) exit\n',
        ]

        for line in lines:
            if line != await self.readline():
                raise ProtocolError('invalid menu')

    async def read_prompt(self) -> None:
        data = b'> '
        response = await self._reader.readexactly(2)

        if data != response:
            raise ProtocolError('invalid prompt')

    async def login(self, username: str, password: str) -> Response:
        await self.read_anonymous_menu()

        await self.read_prompt()
        await self.writeline(b'1')

        line = b'[*] Please, enter username:\n'
        if line != await self.readline():
            raise ProtocolError('invalid login interface')

        await self.read_prompt()
        await self.writeline(username.encode())

        response = await self.readline()

        line = b'[-] Sorry, user does not exist.\n'
        if line == response:
            return Response.BAD_USERNAME

        line = b'[*] Please, enter password:\n'
        if line != response:
            raise ProtocolError('invalid login interface')

        await self.read_prompt()
        await self.writeline(password.encode())

        response = await self.readline()

        line = b'[-] Sorry, password is incorrect.\n'
        if line == response:
            return Response.BAD_PASSWORD

        line = f'[+] Welcome, {username}!\n'.encode()
        if line == response:
            return Response.OK

        raise ProtocolError('invalid login interface')

    async def register(self, username: str) -> Tuple[Response, Tuple[str, str]]:
        await self.read_anonymous_menu()

        await self.read_prompt()
        await self.writeline(b'2')

        line = b'[*] Please, enter username:\n'
        if line != await self.readline():
            raise ProtocolError('invalid register interface')

        await self.read_prompt()
        await self.writeline(username.encode())

        response = await self.readline()

        line = b'[-] Sorry, user already exists.\n'
        if line == response:
            return Response.BAD_USERNAME, None

        line = f'[+] Success! Nice to meet you, {username}!\n'.encode()
        if line != response:
            raise ProtocolError('invalid register interface')

        line = b'[!] Here is your password:\n'
        if line != await self.readline():
            raise ProtocolError('invalid register interface')

        password = (await self.readline()).strip(b'\n').decode()

        line = b'[!] Here is your recovery key:\n'
        if line == await self.readline():
            response = await self.readline()
            recovery_key = response.strip(b'\n').decode()
    
            return Response.OK, (password, recovery_key)

        raise ProtocolError('invalid register interface')

    async def anonymous_exit(self) -> Response:
        await self.read_anonymous_menu()

        await self.read_prompt()
        await self.writeline(b'3')

        line = b'[*] Bye.\n'
        if line == await self.readline():
            return Response.OK

        raise ProtocolError('invalid exit interface')

    async def print_info(self) -> Tuple[Response, Tuple[str, str]]:
        await self.read_user_menu()

        await self.read_prompt()
        await self.writeline(b'1')

        name = await self.readline()
        disease = await self.readline()

        prefix = b'['

        if name.startswith(prefix) and disease.startswith(prefix):
            return Response.OK, (
                name.strip().decode(), disease.strip().decode(),
            )

        raise ProtocolError('invalid print info interface')

    async def change_recovery_key(self, password: str, recovery_key: str) -> Response:
        await self.read_user_menu()

        await self.read_prompt()
        await self.writeline(b'2')

        line = b'[*] Please, enter password:\n'
        if line != await self.readline():
            raise ProtocolError('invalid change recovery key interface')

        await self.read_prompt()
        await self.writeline(password.encode())

        response = await self.readline()

        line = b'[-] Sorry, password is incorrect.\n'
        if line == response:
            return Response.BAD_PASSWORD

        line = b'[*] Please, enter new recovery key:\n'
        if line != response:
            raise ProtocolError('invalid change recovery key interface')

        await self.read_prompt()
        await self.writeline(recovery_key.encode())

        response = await self.readline()

        line = b'[-] Sorry, recovery key is not valid.\n'
        if line == response:
            return Response.BAD_RECOVERY_KEY

        line = b'[+] Success, recovery key has been changed.\n'
        if line == response:
            return Response.OK

        raise ProtocolError('invalid change recovery key interface')

    async def update_disease(self, type: str, name: str, additional: str = None) -> Response:
        await self.read_user_menu()

        await self.read_prompt()
        await self.writeline(b'3')

        line = b'[*] Please, enter disease type:\n'
        if line != await self.readline():
            raise ProtocolError('invalid update disease interface')

        await self.read_prompt()
        await self.writeline(type.encode())

        line = b'[*] Please, enter disease name:\n'
        if line != await self.readline():
            raise ProtocolError('invalid update disease interface')

        await self.read_prompt()
        await self.writeline(name.encode())

        if type == 'mental':
            line = b'[*] Please, enter disease phase:\n'
            if line != await self.readline():
                raise ProtocolError('invalid update disease interface')

            await self.read_prompt()
            await self.writeline(additional.encode())
        elif type == 'infectious':
            line = b'[*] Please, enter disease symptoms:\n'
            if line != await self.readline():
                raise ProtocolError('invalid update disease interface')

            await self.read_prompt()
            await self.writeline(additional.encode())

        line = b'[+] Success, disease has been updated.\n'
        if line == await self.readline():
            return Response.OK

        raise ProtocolError('invalid update disease interface')

    async def logout(self) -> Response:
        await self.read_user_menu()

        await self.read_prompt()
        await self.writeline(b'4')

        return Response.OK

    async def user_exit(self) -> Response:
        await self.read_user_menu()

        await self.read_prompt()
        await self.writeline(b'5')

        line = b'[*] Bye.\n'
        if line == await self.readline():
            return Response.OK

        raise ProtocolError('invalid exit interface')
