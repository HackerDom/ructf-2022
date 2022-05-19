import random
import string

ALPHA = string.ascii_lowercase + string.digits


def gen_string(a=20, b=20):
    return ''.join(random.choice(ALPHA) for _ in range(random.randint(a, b)))


def gen_int():
    return random.randint(0, 10000)
