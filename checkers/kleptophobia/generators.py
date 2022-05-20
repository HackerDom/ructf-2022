import random
import string

ALPHA = string.ascii_letters + string.digits


def gen_string(a=20, b=20, alpha=ALPHA):
    return ''.join(random.choice(alpha) for _ in range(random.randint(a, b)))


def gen_int():
    return random.randint(0, 10000)
