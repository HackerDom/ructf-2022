package abstract

import (
	"errors"
)

type Field[T any] struct {
	cells  []T
	width  int64
	height int64
}

func NewField[T any](width, height int64) *Field[T] {
	cells := make([]T, width*height)

	return &Field[T]{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (field *Field[T]) Width() int64 {
	return field.width
}

func (field *Field[T]) Height() int64 {
	return field.height
}

func (field *Field[T]) Has(x, y int64) bool {
	return 0 <= x && x < field.width &&
		0 <= y && y < field.height
}

func (field *Field[T]) Get(x, y int64) (element T, err error) {
	if !field.Has(x, y) {
		err = errors.New("invalid coordinates")
		return
	}

	element = field.cells[field.position(x, y)]
	return
}

func (field *Field[T]) Set(x, y int64, element T) error {
	if !field.Has(x, y) {
		return errors.New("invalid coordinates")
	}

	field.cells[field.position(x, y)] = element

	return nil
}

func (field Field[T]) Fill(value T) {
	for i := range field.cells {
		field.cells[i] = value
	}
}

func (field *Field[T]) position(x, y int64) int64 {
	return y*field.height + x
}
