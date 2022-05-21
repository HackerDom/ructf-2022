import  itertools
import collections
import websockets
import json
import asyncio
import random
from math import gcd
import sys
import enum
from typing import List, Tuple, Set, Iterator, Sequence, Optional

URL = "0.0.0.0:5051"

S = SymmetricGroup(256)


Coordinates = Tuple[int, int]


class Direction(enum.Enum):
    LEFT = enum.auto()
    RIGHT = enum.auto()
    UP = enum.auto()
    DOWN = enum.auto()
    
    def opposite(self) -> 'Direction':
        if self is Direction.LEFT:
            return Direction.RIGHT

        if self is Direction.RIGHT:
            return Direction.LEFT

        if self is Direction.UP:
            return Direction.DOWN

        if self is Direction.DOWN:
            return Direction.UP
    
    def move(self, coordinates: Coordinates) -> Coordinates:
        x, y = coordinates

        if self is Direction.LEFT:
            return x - 1, y

        if self is Direction.RIGHT:
            return x + 1, y

        if self is Direction.UP:
            return x, y - 1

        if self is Direction.DOWN:
            return x, y + 1


class Snake:
    def __init__(self, initial: List[Coordinates], direction: Direction) -> None:
        self.deque = collections.deque(initial)
        self.direction = direction
        self.tail = self.deque[-1]
            
    @property
    def head(self) -> Coordinates:
        return self.deque[0]

    def copy(self) -> 'Snake':
        snake = Snake(self.deque, self.direction)
        snake.tail = self.tail

        return snake

    def has_intersection(self) -> bool:
        return self.deque.count(self.head) > 1

    def grow(self) -> None:
        self.deque.append(self.tail)
        
    def move(self, direction: Direction) -> None:
        if direction is not self.direction.opposite():
            self.direction = direction

        head = self.direction.move(self.head)
        self.deque.appendleft(head)

        self.tail = self.deque.pop()


MAX_STEPS = 256
LEVEL_WIDTH = 16
LEVEL_HEIGHT = 16
FOOD_COUNT = 8
FOOD_STEPS = LEVEL_WIDTH * LEVEL_HEIGHT // FOOD_COUNT


def get_possible_directions(snake: Snake) -> Iterator[Direction]:
    directions = [
        Direction.LEFT,
        Direction.RIGHT,
        Direction.UP,
        Direction.DOWN,
    ]

    for direction in directions:
        snake2 = snake.copy()

        if direction is snake.direction.opposite():
            continue

        snake2.move(direction)

        if snake2.has_intersection():
            continue

        x, y = snake2.head

        if not (0 <= x < LEVEL_WIDTH and 0 <= y < LEVEL_HEIGHT):
            continue

        yield direction


State = Tuple[Snake, List[Direction], int, int]


def find_path(
        state: State,
        target: Coordinates,
        visibility: Set[int],
        max_iterations: int,
        random_bound: int
) -> Optional[State]:
    queue: Sequence[State] = collections.deque([state])

    for _ in range(max_iterations):
        if len(queue) == 0:
            return

        current_snake, current_moves, current_index, steps = queue.pop()

        if steps > max(visibility):
            continue

        for direction in get_possible_directions(current_snake):
            if random.randint(0, random_bound) == 0:
                continue

            next_snake = current_snake.copy()
            next_snake.move(direction)

            next_moves = current_moves.copy()
            next_moves.append(direction)

            next_index = current_index

            if next_snake.head == target and steps in visibility:
                next_snake.grow()
                next_index += 1

                return next_snake, next_moves, next_index, steps + 1

            queue.append((next_snake, next_moves, next_index, steps + 1))


def solve(field: List[int], snake: Snake) -> List[Direction]:
    food: List[Coordinates] = [(0, 0)] * FOOD_COUNT

    for i, element in enumerate(field):
        x, y = i % LEVEL_WIDTH, i // LEVEL_HEIGHT

        if element % FOOD_STEPS == 0:
            position = element // FOOD_STEPS
            food[position] = (x, y)

    food_visibility: List[Set[int]] = []

    for i in range(len(food)):
        food_visibility.append(
            set([(i + 1) * FOOD_STEPS - 2, (i + 1) * FOOD_STEPS - 1]),
        )

    food = food
    food_visibility = food_visibility
    state: State = (snake, [], 0, 0)
    while True:
        try:
            for target, visibility in zip(food, food_visibility):
                for _ in range(1000):
                    next_state = find_path(state, target, visibility, 1_000, 2)

                    if next_state is not None:
                        state = next_state
                        break
                else:
                    raise Exception("tooo long")    
            return state[1]
        except Exception:
            state: State = (snake, [], 0, 0)
                


def find_max_order(fields):
    return max(field.order() for field in fields)


def get_fields_by_counter(sign, log, offset, max_order, secret, init, counter):
    answer = None

    for t in range(-1000, 1000):
        if ((sign * (log + t) - offset * max_order) ^^ counter) % max_order == secret:
            answer = sign * (log + t)
            break

    return Permutation(init ^ (-1 * answer)), Permutation(init ^ answer)

def exploit(data, game_counter):
    counters = [counter for counter, _ in data]
    fields = [S(field) for _, field in data]
    
    max_order = find_max_order(fields)
    print(max_order)

    db = collections.defaultdict(list)
    bits = 7

    for f1, f2 in itertools.combinations(fields, 2):
        if f1.order() != max_order and f2.order() != max_order:
            continue

        try:
            log = discrete_log(f2, f1)
        except Exception:
            continue

        g, inv, _ = xgcd(log - 1, max_order)
        if g != 1:
            continue

        for t1 in range(1, 1_000):
            s = t1 * inv % max_order
            s_msb = (s >> bits) << bits

            db[s_msb].append(s)

        for t2 in range(1, 1_000):
            s = (-log * t2) * inv % max_order
            s_msb = (s >> bits) << bits

            db[s_msb].append(s)
    max_freq = max(len(ks) for ks in db.values())
    s_candidates = [s for s, ks in db.items() if len(ks) == max_freq]
    print(max_freq, s_candidates)

    init = None

    for f1, f2 in itertools.combinations(fields, 2):
        if f1.order() != max_order and f2.order() != max_order:
            continue

        try:
            log = discrete_log(f2, f1)
        except Exception:
            continue

        for s_msb in s_candidates:
            for t1, t2 in itertools.product(range(1, 2 ^ bits), repeat=2):
                if (s_msb + t1) * log % max_order == (s_msb + t2) % max_order:
                    t_diff = t1 - t2
                    g, t_diff_inv, _ = xgcd(t_diff, max_order)

                    if g > 1:
                        continue

                    init = (f1 / f2) ^ t_diff_inv
                    break

    print("INIT", Permutation(init))

    logs = []
    new_counters = []

    for COUNTER, FIELD in zip(counters, fields):
        try:
            log = discrete_log(FIELD, init)
        except Exception:
            continue

        logs.append(log)
        new_counters.append(COUNTER)

    print(len(logs), len(new_counters))

    secrets = set()

    for sign, offset in itertools.product([-1, 1], range(1, 2_000)):
        secrets.clear()

        for log, counter in zip(logs, new_counters):
            secrets.add(((sign * log - offset * max_order) ^^ counter) % max_order)

        if len(secrets) == 1:
            secret = secrets.pop()
            break

    print(offset, sign, secret)
    field1, field2 = get_fields_by_counter(sign, log, offset, max_order, secret, init, data[0][0])
    checkfield = Permutation(data[0][1])
    if checkfield == field1:
        valid_init = init
    elif checkfield == field2:
        valid_init = init ^ (-1)
    else:
        raise Exception("No valid init")

    next_field = get_fields_by_counter(sign, log, offset, max_order, secret, valid_init, game_counter)
    return next_field[0]

def multiply(left_perm, right_perm):
    new_perm = [0 for _ in range(len(left_perm))]
    for i, element in enumerate(right_perm):
        new_perm[i] = left_perm[element]
    
    return new_perm


def invert(perm):
    new_perm = [0 for _ in range(len(perm))]
    for i in range(len(perm)):
        new_perm[perm[i]] =  i
    
    return new_perm


def exponentiation(perm, power):
    result = [i for i in range(len(perm))]

    if power < 0:
        perm = invert(perm)
        power = -power

    while power > 0:
        if power & 1 == 1:
            result = multiply(result, perm)
        
        perm = multiply(perm, perm)
        power  = power >> 1
    
    return result

def print_map(map):
    for el in map:
        print(" ".join(el))

async def sploit(id):
    counters_fields = []
    async with websockets.connect(f"ws://{URL}/play") as ws:
        for i in range(30):
            await ws.send(str(json.dumps({"id": id})))
            resp = await ws.recv()
            resp = json.loads(resp)
            while not resp.get("gameResult"):
                await ws.send(str(json.dumps({"direction": "w",
                    "closeGame": False,
                    "newGame": False})))
                resp = await ws.recv()
                resp = json.loads(resp)
            counters_fields.append((resp["counter"], [x + 1 for x in resp["permutation"]]))
            if i != 29:
                await ws.send(str(json.dumps({
                    "direction": "w",
                    "closeGame": False,
                    "newGame": True
                })))
                resp = await ws.recv()
                resp = json.loads(resp)
        await ws.send(str(json.dumps({
            "direction": "",
            "closeGame": False,
            "newGame": True
        })))
        resp = await ws.recv()
        resp = json.loads(resp)
        res = exploit(counters_fields, resp["counter"])
        res = [x - 1 for x in res]
        print("field for counter", resp["counter"], "is", res)
        snake = Snake([(4, 1), (3, 1), (2, 1), (1, 1)], Direction.RIGHT)
        directions = solve(res, snake)
        print(len(directions))
        for direction in directions:
            if direction == Direction.UP:
                move = "w"
            elif direction == Direction.DOWN:
                move = "s"
            elif direction == Direction.LEFT:
                move = "a"
            elif direction == Direction.RIGHT:
                move = "d"
            else:
                raise Exception("Bad direction!!")
            await ws.send(str(json.dumps({
                "direction": move,
                "closeGame": False,
                "newGame": False
            })))
            resp = await ws.recv()
            print(direction)
            try:
                print_map(json.loads(resp)["gameMap"])
                print("=============================")
            except Exception:
                print(json.loads(resp))

asyncio.get_event_loop().run_until_complete(sploit(sys.argv[1]))
