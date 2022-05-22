from hashlib import md5


class DecodingError(Exception):
    pass


def expand_key(key, S):
    def key_to_blocks(key):
        return [key[i:i+4] for i in range(0, len(key), 4)]

    def blocks_to_key(blocks):
        return b''.join(blocks)

    def shift(blocks):
        return blocks[1:] + [updateLastBlock(blocks[0])]

    def updateLastBlock(block):
        return bytes(S[x] for x in block)

    return blocks_to_key(shift(key_to_blocks(key)))


def pad(text, block_size):
    pad_size = block_size - (len(text) % block_size)
    return text + bytes([pad_size] * pad_size)


def unpad(text):
    padding = text[-1]
    if padding >= 16:
        raise DecodingError("wrong padding")
    if set(text[-padding:-1]) != set(text[-1:]):
        raise DecodingError("wrong padding")
    return text[:-padding]


def inverse(arr):
    return list(map(arr.index, range(len(arr))))


def substitute(text, S):
    return bytes(S[x] for x in text)


def permute(text, P):
    return bytes(text[P[i]] for i in range(len(text)))


def xor(a, b):
    return bytes(x^y for x,y in zip(a, b))


def get_hash(data):
    return md5(data).digest()
