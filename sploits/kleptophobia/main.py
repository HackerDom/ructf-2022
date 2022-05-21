from hashlib import md5
from sys import argv

import grpc
from google.protobuf.message import DecodeError

import generators
import models.models_pb2 as pb2
import models.grpc_pb2_grpc as pb2_grpc

from crypto.cipher import Cipher, BLOCK_SIZE
from crypto.crypto_utils import pad, unpad, xor

from crack import crack


PORT = 50051
HOST = argv[1]


def get_stub():
    channel = grpc.insecure_channel('{}:{}'.format(HOST, PORT))
    return pb2_grpc.KleptophobiaStub(channel)


def ping(stub):
    message="ping"
    ping_rsp = stub.Ping(pb2.PingBody(message=message))
    assert ping_rsp.message == message


def register(stub, username, password, private_person):
    register_req = pb2.RegisterReq(
        username=username,
        password=password,
        person=private_person, 
    )
    reqister_rsp = stub.Register(register_req)
    assert reqister_rsp.status == pb2.RegisterRsp.Status.OK


def get_public_info(stub, username):
    get_public_info_rsp = stub.GetPublicInfo(pb2.GetByUsernameReq(
        username=username
    ))
    assert get_public_info_rsp.status == pb2.GetPublicInfoRsp.Status.OK
    return get_public_info_rsp.person


def get_encrypted_full_info(stub, username):
    get_encrypted_full_info_rsp = stub.GetEncryptedFullInfo(pb2.GetByUsernameReq(
        username=username
    ))
    assert get_encrypted_full_info_rsp.status == pb2.GetEncryptedFullInfoRsp.Status.OK
    return get_encrypted_full_info_rsp.encryptedFullInfo


def public_to_private_person(public_person):
    return pb2.PrivatePerson(
        first_name=public_person.first_name,
        middle_name=public_person.middle_name_restricted,
        second_name=public_person.second_name,
        room=public_person.room,
        diagnosis='*'*31 + '=',
    )


def gen_person():
    username = generators.gen_string(10, 10)
    first_name = generators.gen_name(7, 7)
    middle_name = generators.gen_name(16, 16)
    second_name = generators.gen_name(16, 16)
    room = generators.gen_int()
    flag = generators.gen_string(31, 31) + '='
    return pb2.PrivatePerson(
        first_name=first_name,
        middle_name=middle_name,
        second_name=second_name,
        room=room,
        diagnosis=flag
    ), username


def hack(iv, ct_block):
    for i in range(0x100):
        print(i)
        pt_block = xor(pad(bytes([i, ord('=')]), BLOCK_SIZE), iv)
        yield bytes.fromhex(hex(crack(pt_block, ct_block))[2:].zfill(32))


if __name__ == '__main__':
    stub = get_stub()
    ping(stub)

    password = generators.gen_string()
    private_person, username = gen_person()
    register(stub, username, password, private_person)

    public_person = get_public_info(stub, username)
    fake_public_person = public_to_private_person(public_person)

    test = fake_public_person.SerializeToString()
    test_len, test_padded = len(test)%16, pad(test, BLOCK_SIZE)

    blocks = [test_padded[i : i + BLOCK_SIZE] for i in range(0, len(test_padded), BLOCK_SIZE)]
    for block in blocks:
        print(block)

    enc_msg = get_encrypted_full_info(stub, username)
    real_key = md5(password.encode()).digest()

    for maybe_key in hack(enc_msg[-32:-16], enc_msg[-16:]):
        dec_msg = Cipher(maybe_key).decrypt(enc_msg)
        try:
            maybe_person = pb2.PrivatePerson()
            maybe_person.ParseFromString(dec_msg)
            if maybe_person.diagnosis.endswith('='):
                print(maybe_person.diagnosis, maybe_person.diagnosis == private_person.diagnosis)
                break
        except DecodeError:
            continue