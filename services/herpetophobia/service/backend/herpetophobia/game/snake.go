package game

import (
	"snake/game/abstract"

	"github.com/pkg/errors"
)

const MaxSnakeLength = 1024

type Snake struct {
	body      *abstract.Deque[Coordinate]
	direction Direction
	tail      Coordinate
}

func NewSnake(body []Coordinate, direction Direction) (snake *Snake, err error) {
	if len(body) == 0 {
		err = errors.New("length should be positive")
		return
	}

	snake = &Snake{
		body:      abstract.NewDeque[Coordinate](MaxSnakeLength),
		direction: direction,
	}

	for _, coordinate := range body {
		_ = snake.body.AppendRight(coordinate)
	}

	snake.tail, _ = snake.body.PeekLeft()

	return
}

func (snake Snake) Direction() Direction {
	return snake.direction
}

func (snake *Snake) Length() int64 {
	return snake.body.Count()
}

func (snake *Snake) Head() (Coordinate, error) {
	return snake.body.PeekRight()
}

func (snake *Snake) Tail() (Coordinate, error) {
	return snake.body.PeekLeft()
}

func (snake *Snake) Body() []Coordinate {
	return snake.body.Elements()
}

func (snake *Snake) Grow() error {
	return snake.body.AppendLeft(snake.tail)
}

func (snake *Snake) Move(direction Direction) error {
	if direction != snake.direction.Opposite() {
		snake.direction = direction
	}

	head, err := snake.Head()
	if err != nil {
		return errors.Wrap(err, "failed to get head")
	}

	head = head.Move(snake.direction)
	_ = snake.body.AppendRight(head)

	snake.tail, err = snake.body.PopLeft()
	if err != nil {
		return errors.Wrap(err, "failed to pop tail")
	}

	return nil
}

func (snake *Snake) HasIntersection() bool {
	if snake.Length() == 0 {
		return false
	}

	head, _ := snake.Head()
	count := 0

	for _, coordinate := range snake.Body() {
		if coordinate.Equals(head) {
			count += 1
		}
	}

	return count > 1
}
