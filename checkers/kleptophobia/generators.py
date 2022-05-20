import random
import string

import names

ALPHA = string.ascii_letters + string.digits


def gen_string(a=20, b=20, alpha=ALPHA):
    return ''.join(random.choice(alpha) for _ in range(random.randint(a, b)))


def gen_int():
    return random.randint(0, 10000)


def gen_username():
    def _rnd_case(c):
        if random.random() > 0.5:
            return c.upper()
        return c.lower()
    username = ''.join(_rnd_case(c) for c in names.get_first_name())
    if 6 <= len(username) <= 8:
        return username

    return gen_string(6, 8)
