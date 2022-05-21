package game

import (
	"fmt"
	"snake/game/abstract"
	"strings"

	"github.com/pkg/errors"
)

type Level struct {
	steps     int64
	foodCount int64
	status    Status

	Food      []Coordinate
	Field     *abstract.Field[Cell]
	Snake     *Snake
	MaxSteps  int64
	FoodSteps int64
	FoodTTl   int64
}

func (level *Level) Status() Status {
	return level.status
}

func (level *Level) Step(direction Direction) error {
	if level.status != STATUS_UNFINISHED {
		return nil
	}

	if len(level.Food) > 0 && level.steps%level.FoodSteps == level.FoodSteps-1 {
		var food Coordinate
		food, level.Food = level.Food[0], level.Food[1:]

		err := level.Field.Set(food.X, food.Y, CELL_FOOD)
		if err != nil {
			return err
		}

		level.foodCount += 1
	}

	err := level.Snake.Move(direction)
	if err != nil {
		return errors.Wrap(err, "failed to move the snake")
	}

	if level.Snake.HasIntersection() {
		level.status = STATUS_LOSE
		return nil
	}

	head, err := level.Snake.Head()
	if err != nil {
		return errors.Wrap(err, "failed to get snake's head")
	}

	if !level.Field.Has(head.X, head.Y) {
		level.status = STATUS_LOSE
		return nil
	}

	cell, err := level.Field.Get(head.X, head.Y)
	if err != nil {
		return err
	}

	if cell == CELL_FOOD {
		level.foodCount -= 1

		err = level.Field.Set(head.X, head.Y, CELL_EMPTY)
		if err != nil {
			return err
		}

		err = level.Snake.Grow()
		if err != nil {
			return errors.Wrap(err, "failed to grow the snake")
		}
	}

	if len(level.Food) == 0 && level.foodCount <= 0 {
		level.status = STATUS_WIN
		return nil
	}

	level.steps += 1

	if level.steps > level.MaxSteps || level.FoodSteps+level.FoodTTl <= level.steps {
		level.status = STATUS_LOSE
		return nil
	}

	return nil
}

func (level *Level) Str() string {
	var lines [][]string

	for y := int64(0); y < level.Field.Height(); y++ {
		var line []string

		for x := int64(0); x < level.Field.Width(); x++ {
			cell, _ := level.Field.Get(x, y)

			switch cell {
			case CELL_EMPTY:
				line = append(line, ".")
				break
			case CELL_FOOD:
				line = append(line, "*")
				break
			}
		}

		lines = append(lines, line)
	}

	for _, coordinates := range level.Snake.Body() {
		if level.Field.Has(coordinates.X, coordinates.Y) {
			lines[coordinates.Y][coordinates.X] = "#"
		}
	}

	head, _ := level.Snake.Head()

	if level.Field.Has(head.X, head.Y) {
		lines[head.Y][head.X] = "@"
	}
	levelStr := fmt.Sprintf("steps: %d\n", level.steps)
	for _, line := range lines {
		levelStr += strings.Join(line, " ") + "\n"
	}
	return levelStr
}

func (level *Level) Map() [][]string {
	var lines [][]string

	for y := int64(0); y < level.Field.Height(); y++ {
		var line []string

		for x := int64(0); x < level.Field.Width(); x++ {
			cell, _ := level.Field.Get(x, y)

			switch cell {
			case CELL_EMPTY:
				line = append(line, ".")
				break
			case CELL_FOOD:
				line = append(line, "*")
				break
			}
		}

		lines = append(lines, line)
	}

	for _, coordinates := range level.Snake.Body() {
		if level.Field.Has(coordinates.X, coordinates.Y) {
			lines[coordinates.Y][coordinates.X] = "#"
		}
	}

	head, _ := level.Snake.Head()

	if level.Field.Has(head.X, head.Y) {
		lines[head.Y][head.X] = "@"
	}
	return lines
}

func (level Level) Steps() int64 {
	return level.steps
}
