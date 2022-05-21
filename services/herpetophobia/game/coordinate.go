package game

type Coordinate struct {
	X int64
	Y int64
}

func (coordinate Coordinate) Equals(other Coordinate) bool {
	return coordinate.X == other.X && coordinate.Y == other.Y
}

func (coordinate Coordinate) Move(direction Direction) Coordinate {
	switch direction {
	case DIRECTION_LEFT:
		return Coordinate{
			X: coordinate.X - 1,
			Y: coordinate.Y,
		}
	case DIRECTION_RIGHT:
		return Coordinate{
			X: coordinate.X + 1,
			Y: coordinate.Y,
		}
	case DIRECTION_UP:
		return Coordinate{
			X: coordinate.X,
			Y: coordinate.Y - 1,
		}
	case DIRECTION_DOWN:
		return Coordinate{
			X: coordinate.X,
			Y: coordinate.Y + 1,
		}

	default:
		panic("not implemented")
	}
}
