from itertools import product
from string import ascii_uppercase, digits
from hashlib import md5
from sys import argv

import grpc
from google.protobuf.message import DecodeError

import generators
import models.models_pb2 as pb2
import models.grpc_pb2_grpc as pb2_grpc

from crypto.cipher import Cipher, BLOCK_SIZE
from crypto.crypto_utils import pad, unpad, xor, DecodingError

from crack import crack, key_from_solver, get_solver


PORT = 50051
HOST = argv[1]
ALPHA = (ascii_uppercase + digits).encode()


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
    middle_name = generators.gen_name(26, 26)
    second_name = generators.gen_name(6, 7)
    room = generators.gen_int()
    flag = generators.gen_string(31, 31) + '='
    return pb2.PrivatePerson(
        first_name=first_name,
        middle_name=middle_name,
        second_name=second_name,
        room=room,
        diagnosis=flag.upper()
    ), username


def hack(unknown_len, iv, ct_block):
    pt_block = xor(pad(bytes([0]*unknown_len + [ord('=')]), BLOCK_SIZE), iv)
    states, key_bits, pt_bits = crack(unknown_len*8, pt_block[unknown_len:], ct_block)

    solver = get_solver()
    for x in product(ALPHA, repeat=unknown_len):
        x = xor(bytes(x), iv[:unknown_len])
        x_bits = list(map(int, bin(int(x.hex(), 16))[2:].zfill(unknown_len*8)))
        new_states = [x_bit == pt_bit for x_bit, pt_bit in zip(x_bits, pt_bits)]
        yield key_from_solver(key_bits, states + new_states, solver)
        solver.reset()


if __name__ == '__main__':
    stub = get_stub()
    ping(stub)

    if len(argv) == 3:
        username = argv[2]
    else:
        print('FAKE check')
        password = generators.gen_string()
        private_person, username = gen_person()
        register(stub, username, password, private_person)

    public_person = get_public_info(stub, username)
    print(public_person)
    fake_public_person = public_to_private_person(public_person)

    test = fake_public_person.SerializeToString()
    test_len, test_padded = len(test)%16, pad(test, BLOCK_SIZE)

    enc_msg = get_encrypted_full_info(stub, username)

    print(f'HACKING: {test_len}')
    for maybe_key in hack(test_len-1, enc_msg[-32:-16], enc_msg[-16:]):
        try:
            dec_msg = Cipher(maybe_key).decrypt(enc_msg)
        except DecodingError:
            continue
        try:
            maybe_person = pb2.PrivatePerson()
            maybe_person.ParseFromString(dec_msg)
            if (
                maybe_person.first_name == public_person.first_name and
                maybe_person.second_name == public_person.second_name
            ):
                print('FLAG?', maybe_person.diagnosis)
                break
        except DecodeError:
            continue