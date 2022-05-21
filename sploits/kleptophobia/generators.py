import random
import string

ALPHA = string.ascii_letters + string.digits


def gen_string(min_size=20, max_size=20, alpha=ALPHA):
    return ''.join(random.choice(alpha) for _ in range(random.randint(min_size, max_size)))


def gen_int():
    return random.randint(0, 10000)


def gen_name(min_size=6, max_size=30, alpha=string.ascii_lowercase):
    s = gen_string(min_size, max_size, alpha)
    return s[0].upper() + s[1:]