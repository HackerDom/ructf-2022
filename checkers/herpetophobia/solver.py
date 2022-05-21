import enum
import random
import collections
from typing import List, Tuple, Set, Iterator, Sequence, Optional

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