import random
import string

from gornilo import Checker, CheckRequest, GetRequest, PutRequest, Verdict

checker = Checker()


def get_random_str():
    return "".join(random.choices(string.digits + string.ascii_letters, k=random.randint(1,11)))


@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    pass


@checker.define_put(vuln_rate=1, vuln_num=1)
async def put_flag(request: PutRequest):
    pass


@checker.define_get(vuln_num=1)
async def get_flag(request: GetRequest):
    pass


if __name__ == '__main__':
    checker.run()
