import random
import string
import websockets
import json
from gornilo.http_clients import requests_with_retries
from crypto import exponentiation
from gornilo import NewChecker, CheckRequest, GetRequest, PutRequest, Verdict, VulnChecker
from math import gcd
from solver import solve, Snake, Direction

checker = NewChecker()
init_perm = [157, 79, 170, 8, 108, 234, 163, 16, 251, 181, 23, 148, 55, 162, 211, 186, 194, 222, 152, 207, 57, 97,
                 87, 45, 245, 141, 142, 40, 13, 92, 89, 64, 191, 102, 247, 178, 28, 138, 118, 68, 226, 24, 151, 103, 15,
                 139, 154, 244, 180, 83, 82, 196, 171, 167, 31, 155, 63, 246, 38, 200, 228, 120, 218, 204, 10, 238, 47,
                 56, 146, 185, 172, 158, 133, 53, 117, 42, 193, 241, 206, 86, 161, 0, 77, 243, 149, 239, 121, 129, 2,
                 85, 159, 59, 96, 164, 81, 220, 114, 18, 214, 65, 60, 125, 188, 201, 104, 174, 153, 75, 240, 223, 126,
                 35, 189, 113, 27, 236, 122, 143, 124, 73, 227, 43, 49, 67, 187, 48, 99, 250, 39, 20, 165, 115, 1, 177,
                 93, 232, 202, 249, 116, 54, 6, 242, 252, 69, 255, 22, 176, 197, 110, 5, 61, 169, 254, 183, 19, 229,
                 109, 150, 111, 131, 156, 253, 208, 145, 58, 179, 76, 7, 91, 78, 37, 233, 212, 9, 215, 192, 62, 209, 33,
                 32, 198, 168, 17, 195, 136, 166, 98, 130, 71, 248, 90, 217, 25, 30, 112, 34, 231, 3, 237, 21, 80, 224,
                 100, 66, 52, 84, 106, 4, 101, 205, 26, 105, 128, 225, 210, 135, 137, 175, 95, 70, 132, 203, 182, 29,
                 219, 190, 199, 44, 235, 140, 147, 74, 144, 46, 123, 216, 221, 14, 94, 127, 119, 36, 184, 88, 107, 12,
                 41, 134, 213, 72, 173, 160, 50, 51, 11, 230]


def get_random_str():
    return "".join(random.choices(string.digits + string.ascii_letters, k=random.randint(1,11)))


def create_game(prize, url):
    order = 8933296680
    power = random.randint(1, order - 1)
    while gcd(order, power) != 1:
        power = random.randint(1, order - 1)
    secret = str(random.randint(10000000, 89332966))
    body = {
        "secret": secret,
        "init": exponentiation(init_perm, power),
        "flag": prize
    }

    res = requests_with_retries().post(f"http://{url}:5051/create", json=body)
    return res.json()["id"], secret, power


@checker.define_check
async def check_service(request: CheckRequest) -> Verdict:
    prize = get_random_str()
    try:
        game_id, secret, power = create_game(prize, request.hostname)
    except Exception:
        return Verdict.DOWN("could not create game")
    try:
        async with websockets.connect(f"ws://{request.hostname}:5051/play") as ws:
            await ws.send(json.dumps({"id": game_id}))
            resp = await ws.recv()
            resp = json.loads(resp)
            while not resp.get("gameResult"):
                await ws.send(str(json.dumps({"direction": "w",
                                              "closeGame": False,
                                              "newGame": False})))
                resp = await ws.recv()
                resp = json.loads(resp)
            game_map = resp["permutation"]
            counter = resp["counter"]
    except Exception:
        return Verdict.DOWN("could not play and lose")

    field = exponentiation(exponentiation(init_perm, power), int.from_bytes(secret.encode("utf-8"), "big") ^ counter)
    if field != game_map:
        return Verdict.MUMBLE("wrong permutation returned after lose")
    try:
        async with websockets.connect(f"ws://{request.hostname}:5051/play") as ws:
            await ws.send(json.dumps({"id": game_id}))
            resp = await ws.recv()
            resp = json.loads(resp)
            snake = Snake([(4, 1), (3, 1), (2, 1), (1, 1)], Direction.RIGHT)
            new_field = exponentiation(exponentiation(init_perm, power), int.from_bytes(secret.encode("utf-8"), "big") ^ int(resp["counter"]))
            directions = solve(new_field, snake)
            for direction in directions:
                if direction == Direction.UP:
                    move = "w"
                if direction == Direction.DOWN:
                    move = "s"
                if direction == Direction.LEFT:
                    move = "a"
                if direction == Direction.RIGHT:
                    move = "d"
                await ws.send(str(json.dumps({
                    "direction": move,
                    "closeGame": False,
                    "newGame": False
                })))
                resp = await ws.recv()
                resp = json.loads(resp)
            print(resp)
            if resp.get("gameResult") != "win":
                return Verdict.MUMBLE("could not win the game")
            if resp.get("prize") != prize:
                return Verdict.MUMBLE("wrong prize")
            if resp.get("permutation") != new_field:
                return Verdict.MUMBLE("wrong permutation returned after win")
    except Exception:
        return Verdict.DOWN("could not play the game")

    return Verdict.OK()


@checker.define_vuln("flag_id is game id")
class GameChecker(VulnChecker):
    @staticmethod
    def put(request: PutRequest) -> Verdict:
        try:
            game_id, secret, power = create_game(request.flag, request.hostname)
        except Exception:
            return Verdict.DOWN("could not create game")
        if not game_id:
            return Verdict.MUMBLE("could not create game")
        flag_id = json.dumps({
            "game_id": game_id,
            "secret": secret,
            "power": power
        })

        return Verdict.OK_WITH_FLAG_ID(game_id, flag_id)

    @staticmethod
    async def get(request: GetRequest) -> Verdict:
        flag_id = json.loads(request.flag_id)
        secret = flag_id["secret"]
        power = flag_id["power"]
        game_id = flag_id["game_id"]
        try:
            async with websockets.connect(f"ws://{request.hostname}:5051/play") as ws:
                await ws.send(json.dumps({"id": game_id}))
                resp = await ws.recv()
                resp = json.loads(resp)
                snake = Snake([(4, 1), (3, 1), (2, 1), (1, 1)], Direction.RIGHT)
                new_field = exponentiation(exponentiation(init_perm, power),
                                           int.from_bytes(secret.encode("utf-8"), "big") ^ int(resp["counter"]))
                directions = solve(new_field, snake)
                for direction in directions:
                    if direction == Direction.UP:
                        move = "w"
                    if direction == Direction.DOWN:
                        move = "s"
                    if direction == Direction.LEFT:
                        move = "a"
                    if direction == Direction.RIGHT:
                        move = "d"
                    await ws.send(str(json.dumps({
                        "direction": move,
                        "closeGame": False,
                        "newGame": False
                    })))
                    resp = await ws.recv()
                    resp = json.loads(resp)
                print(resp)
                if resp.get("gameResult") != "win":
                    return Verdict.MUMBLE("could not win the game")
                if resp.get("prize") != request.flag:
                    return Verdict.CORRUPT("wrong flag")
                if resp.get("permutation") != new_field:
                    return Verdict.MUMBLE("wrong permutation returned after win")
        except Exception:
            return Verdict.DOWN("could not play game")
        return Verdict.OK()


if __name__ == '__main__':
    checker.run()
