import sys

from crypto.crypto_utils import expand_key, inverse, pad, permute, substitute, unpad, xor

BLOCK_SIZE = 16
ROUNDS = 13
S = [120, 112, 116, 124, 232, 224, 228, 236, 121, 113, 117, 125, 233, 225, 229, 237, 57, 49, 53, 61, 169, 161, 165, 173, 56, 48, 52, 60, 168, 160, 164, 172, 216, 208, 212, 220, 72, 64, 68, 76, 217, 209, 213, 221, 73, 65, 69, 77, 153, 145, 149, 157, 9, 1, 5, 13, 152, 144, 148, 156, 8, 0, 4, 12, 88, 80, 84, 92, 200, 192, 196, 204, 89, 81, 85, 93, 201, 193, 197, 205, 25, 17, 21, 29, 137, 129, 133, 141, 24, 16, 20, 28, 136, 128, 132, 140, 248, 240, 244, 252, 104, 96, 100, 108, 249, 241, 245, 253, 105, 97, 101, 109, 185, 177, 181, 189, 41, 33, 37, 45, 184, 176, 180, 188, 40, 32, 36, 44, 122, 114, 118, 126, 234, 226, 230, 238, 123, 115, 119, 127, 235, 227, 231, 239, 59, 51, 55, 63, 171, 163, 167, 175, 58, 50, 54, 62, 170, 162, 166, 174, 218, 210, 214, 222, 74, 66, 70, 78, 219, 211, 215, 223, 75, 67, 71, 79, 155, 147, 151, 159, 11, 3, 7, 15, 154, 146, 150, 158, 10, 2, 6, 14, 90, 82, 86, 94, 202, 194, 198, 206, 91, 83, 87, 95, 203, 195, 199, 207, 27, 19, 23, 31, 139, 131, 135, 143, 26, 18, 22, 30, 138, 130, 134, 142, 250, 242, 246, 254, 106, 98, 102, 110, 251, 243, 247, 255, 107, 99, 103, 111, 187, 179, 183, 191, 43, 35, 39, 47, 186, 178, 182, 190, 42, 34, 38, 46]
S_inv = inverse(S)
P = [12, 14, 13, 15, 0, 2, 1, 3, 4, 6, 5, 7, 8, 10, 9, 11]
P_inv = inverse(P)


class Cipher:
    def __init__(self, key: bytes) -> None:
        assert len(key) == BLOCK_SIZE
        self._round_keys = self._gen_round_keys(key)

    def _gen_round_keys(self, key):
        keys = []
        for _ in range(ROUNDS):
            keys.append(expand_key(key, S))
            key = keys[-1]
        return keys

    def _encrypt_block(self, block, round_key):
        block = substitute(block, S)
        block = permute(block, P)
        return xor(block, round_key)
    
    def _decrypt_block(self, block, round_key):
        block = xor(block, round_key)
        block = permute(block, P_inv)
        return substitute(block, S_inv)
    
    def encrypt(self, pt):
        pt = pad(pt, BLOCK_SIZE)
        blocks = [pt[i : i+BLOCK_SIZE] for i in range(0, len(pt), BLOCK_SIZE)]
        for round in range(ROUNDS):
            round_key = self._round_keys[round]
            blocks = [self._encrypt_block(block, round_key) for block in blocks]
        return b''.join(blocks)

    def decrypt(self, ct):
        blocks = [ct[i : i+BLOCK_SIZE] for i in range(0, len(ct), BLOCK_SIZE)]
        for round in range(ROUNDS-1, -1, -1):
            round_key = self._round_keys[round]
            blocks = [self._decrypt_block(block, round_key) for block in blocks]
        return unpad(b''.join(blocks))
