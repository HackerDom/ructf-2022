package generators

import (
	"encoding/binary"
	"snake/game"
	"snake/game/abstract"
	"snake/math"

	"github.com/pkg/errors"
)

const SizeofInt64 = 8

const MaxSteps = 256
const FoodCount = 8
const LevelWidth, LevelHeight = 16, 16
const FoodTTL = 2
const FoodNum = 1

var InitialSnakeBody = []game.Coordinate{
	{X: 0, Y: 1},
	{X: 1, Y: 1},
	{X: 2, Y: 1},
	{X: 3, Y: 1},
}

var InitialSnakeDirection = game.DIRECTION_RIGHT

var MainGroup = math.NewSymmetricGroup[byte](LevelWidth * LevelHeight)

func GenerateSeed(initial []byte, secret []byte, counter uint64) ([]byte, error) {
	permutation, err := math.NewPermutation(initial)
	if err != nil {
		return nil, err
	}

	if !MainGroup.Contains(permutation) {
		return nil, errors.New("main group does not contain initial permutation")
	}

	if len(secret) == 0 {
		return nil, errors.New("secret is empty")
	}

	secret = append(secret, make([]byte, SizeofInt64-len(secret)%SizeofInt64)...)

	var power = counter

	for i := 0; i < len(secret); i += SizeofInt64 {
		power ^= binary.BigEndian.Uint64(secret[i : i+SizeofInt64])
	}

	result, err := math.Exponentiation(MainGroup, permutation, int64(power))
	if err != nil {
		return nil, err
	}

	seed, ok := result.(*math.Permutation[byte])
	if !ok {
		return nil, errors.New("got unsupported element after exponentiation")
	}

	return seed.Elements(), nil
}

func GenerateLevel(seed []byte) (level *game.Level, err error) {
	if len(seed) != LevelWidth*LevelHeight {
		err = errors.New("invalid seed length")
		return
	}

	snake, err := game.NewSnake(InitialSnakeBody, InitialSnakeDirection)
	if err != nil {
		err = errors.Wrap(err, "failed to construct new snake")
		return
	}

	level = &game.Level{
		Food:     make([]game.Coordinate, FoodCount),
		Field:    abstract.NewField[game.Cell](LevelWidth, LevelHeight),
		Snake:    snake,
		MaxSteps: MaxSteps,
		FoodTTl:  FoodTTL,
		FoodNum:  FoodNum,
	}

	level.FoodSteps = level.MaxSteps / FoodCount

	for i, element := range seed {
		x, y := i%LevelWidth, i/LevelHeight

		if int(element)%int(level.FoodSteps) == 0 {
			var position = int(element) / int(level.FoodSteps)

			level.Food[position] = game.Coordinate{
				X: int64(x), Y: int64(y),
			}
		}
	}

	return
}
