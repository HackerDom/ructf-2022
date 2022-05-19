#!/usr/bin/env python3.9
import json
import sys
import traceback

import google
import grpc
from grpc._channel import _InactiveRpcError

import models_pb2 as pb2
import models_pb2_grpc as pb2_grpc
from gornilo import CheckRequest, Verdict, PutRequest, GetRequest, VulnChecker, NewChecker

from generators import gen_string, gen_int
import crypto

checker = NewChecker()
PORT = 50051


def get_stub(host):
    channel = grpc.insecure_channel('{}:{}'.format(host, PORT))
    return pb2_grpc.KleptophobiaStub(channel)


class ErrorChecker:
    def __init__(self):
        self.verdict = Verdict.OK()

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc_value, exc_traceback):
        if exc_type in {_InactiveRpcError}:
            if exc_value.code() == grpc.StatusCode.UNAVAILABLE:
                print(exc_value.__dict__['_state'].__dict__)
                self.verdict = Verdict.DOWN("Service is down")

        if exc_type:
            print(exc_type)
            print(exc_value.__dict__)
            traceback.print_tb(exc_traceback, file=sys.stdout)
        return True


@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    with ErrorChecker() as ec:
        stub = get_stub(request.hostname)
        message = gen_string()
        resp = stub.Ping(pb2.PingBody(message=message))
        if resp.message != message:
            print(f"Different ping message: {message} != {resp.message}")
            return Verdict.MUMBLE("Different ping message")
    return ec.verdict


@checker.define_vuln("flag_id is an username")
class CryptoChecker(VulnChecker):
    @staticmethod
    def put(request: PutRequest) -> Verdict:
        with ErrorChecker() as ec:
            stub = get_stub(request.hostname)

            username = gen_string()
            first_name = gen_string()
            second_name = gen_string()
            room = gen_int()

            private_person = pb2.PrivatePerson(
                first_name=first_name,
                second_name=second_name,
                username=username,
                room=room,
                diagnosis=request.flag,
            )

            password = gen_string()
            register_request = pb2.RegisterReq(
                person=private_person,
                password=password,
            )

            register_response = stub.Register(register_request)
            if register_response.status != pb2.RegisterRsp.Status.OK:
                message = f"Not OK response status: {register_response.message}"
                print(message)
                return Verdict.MUMBLE(message)

            flag_id = json.dumps({
                'username': username,
                'password': password,
                'first_name': first_name,
                'second_name': second_name,
                'room': room,
            })

            ec.verdict = Verdict.OK_WITH_FLAG_ID(username, flag_id)
        return ec.verdict

    @staticmethod
    def get(request: GetRequest) -> Verdict:
        with ErrorChecker() as ec:
            flag_id_json = json.loads(request.flag_id)

            username = flag_id_json['username']
            password = flag_id_json['password']

            stub = get_stub(request.hostname)

            get_public_info_rsp = stub.GetPublicInfo(pb2.GetByUsernameReq(username=username))
            if get_public_info_rsp.status != pb2.GetPublicInfoRsp.Status.OK:
                message = f"Not OK response status: {get_public_info_rsp.message}"
                print(message)
                ec.verdict = Verdict.MUMBLE(message)
                return ec.verdict

            for field in ['username', 'first_name', 'second_name', 'room']:
                expected_value = flag_id_json[field]
                real_value = getattr(get_public_info_rsp.person, field)

                if expected_value != real_value:
                    print(f"expected and real values for field {field} are different: {real_value} != {expected_value}")
                    print(f"person: {get_public_info_rsp.person}")
                    ec.verdict = Verdict.MUMBLE('Wrong public info')
                    return ec.verdict

            get_encrypted_full_info_response = stub.GetEncryptedFullInfo(pb2.GetByUsernameReq(username=username))

            if get_encrypted_full_info_response.status != pb2.GetEncryptedFullInfoRsp.Status.OK:
                message = f"Not OK response status: {get_encrypted_full_info_response.message}"
                print(message)
                ec.verdict = Verdict.MUMBLE(message)
                return ec.verdict

            password_hash = crypto.get_hash(password.encode())
            raw_private_person = crypto.decrypt(get_encrypted_full_info_response.encryptedFullInfo, password_hash)

            private_person = pb2.PrivatePerson()
            try:
                private_person.ParseFromString(raw_private_person)
                if private_person.diagnosis != request.flag:
                    print(f"Wrong flag: {private_person.diagnosis} != {request.flag}")
                    ec.verdict = Verdict.CORRUPT('Wrong flag')

            except google.protobuf.message.DecodeError:
                print(f'Incorrect encrypted data: {raw_private_person}')
                ec.verdict = Verdict.MUMBLE('Incorrect encrypted data')

        return ec.verdict


if __name__ == '__main__':
    checker.run()