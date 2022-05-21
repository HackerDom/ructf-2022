import z3
from os import urandom

from crypto.cipher import Cipher, P, ROUNDS, BLOCK_SIZE
from crypto.crypto_utils import pad, xor


def S(x):
    return [
        x[2]^x[5],
        x[3]^1,
        x[1]^x[2]^1,
        x[5]^1,
        x[6]^x[7]^1,
        x[6],
        x[0],
        x[3]^x[4]
    ]


def xor_bytes(a, b):
    return [xor_bits(x, y) for x,y in zip(a, b)]


def xor_bits(a, b):
    return [z3.simplify(x^y) for x,y in zip(a, b)]


def expand_key(key):
    def key_to_blocks(key):
        return [key[i:i+4] for i in range(0, len(key), 4)]

    def shift(blocks):
        return blocks[1:] + [gen_last_block(blocks[0])]

    def gen_last_block(block):
        return list(map(S, block))

    def blocks_to_key(blocks):
        res = []
        for block in blocks:
            res.extend(block)
        return res

    blocks = shift(key_to_blocks(key))
    return blocks_to_key(blocks)


def gen_round_keys(key_bytes):
    round_keys = []
    for _ in range(ROUNDS):
        new_key = expand_key(key_bytes)
        new_key = [[z3.simplify(y) for y in x] for x in new_key]
        round_keys.append(new_key)
        key_bytes = new_key
    return round_keys


def encrypt_one_round(block, round_key):
    block = [S(x) for x in block]
    block = [block[P[i]] for i in range(len(block))]
    return xor_bytes(block, round_key)


def crack(pt, ct):
    pt_first_block = list(map(int, ''.join([bin(x)[2:].zfill(8) for x in pt])))
    pt_bits = [z3.BitVecVal(x, 1) for x in pt_first_block]
    pt_bytes = [pt_bits[i:i+8] for i in range(0, len(pt_bits), 8)]

    ct_first_block = list(map(int, ''.join([bin(x)[2:].zfill(8) for x in ct])))
    ct_bits = [z3.BitVecVal(x, 1) for x in ct_first_block]
    ct_bytes = [ct_bits[i:i+8] for i in range(0, len(ct_bits), 8)]

    key_bits = z3.BitVecs([f'k{i}' for i in range(8*BLOCK_SIZE)], 1)
    key_bytes = [key_bits[i:i+8] for i in range(0, len(key_bits), 8)]

    round_keys = gen_round_keys(key_bytes)

    block = pt_bytes
    for i in range(ROUNDS):
        block = encrypt_one_round(block, round_keys[i])

    s = z3.Solver()
    for p_byte, c_byte in zip(block, ct_bytes):
        for p_bit, c_bit in zip(p_byte, c_byte):
            a = z3.simplify(p_bit ^ c_bit)
            s.add(a == 0)
    res = s.check()
    if res != z3.sat:
        return None
    model = s.model()
    return int(''.join(
        str(model[x].as_long())
        for x in key_bits
    ), 2)


def test():
    pt = b'abcdefg'*7 + b'='
    key = urandom(16)
    cipher = Cipher(key)
    ct = cipher.encrypt(pt)
    assert cipher.decrypt(ct) == pt
    print(ct.hex())

    pt_block = pad(pt, 16)[-16:]
    print(pt_block)
    ct_block = ct[-16:]
    found_key = crack(xor(pt_block, ct[-32:-16]), ct_block)
    found_key = bytes.fromhex(hex(found_key)[2:].zfill(32))
    print(key.hex())
    print(found_key.hex())
    is_good = Cipher(found_key).decrypt(ct) == pt
    print(is_good)


if __name__ == '__main__':
    test()