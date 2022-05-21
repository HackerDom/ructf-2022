package game

type Direction byte

const (
	DIRECTION_LEFT  Direction = iota
	DIRECTION_RIGHT Direction = iota
	DIRECTION_UP    Direction = iota
	DIRECTION_DOWN  Direction = iota
)

func (direction Direction) Opposite() Direction {
	switch direction {
	case DIRECTION_LEFT:
		return DIRECTION_RIGHT
	case DIRECTION_RIGHT:
		return DIRECTION_LEFT
	case DIRECTION_UP:
		return DIRECTION_DOWN
	case DIRECTION_DOWN:
		return DIRECTION_UP

	default:
		panic("not implemented")
	}
}
