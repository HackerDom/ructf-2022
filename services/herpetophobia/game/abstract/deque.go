package abstract

import (
	"errors"
)

type Deque[T any] struct {
	elements []T
	left     int64
	right    int64
	size     int64
	count    int64
}

func NewDeque[T any](size int64) *Deque[T] {
	return &Deque[T]{
		elements: make([]T, size),
		left:     0,
		right:    0,
		size:     size,
		count:    0,
	}
}

func (deque *Deque[T]) Count() int64 {
	return deque.count
}

func (deque *Deque[T]) AppendRight(element T) error {
	if deque.count == deque.size {
		return errors.New("deque is full")
	}

	deque.elements[deque.right] = element
	deque.right = (deque.right + 1) % deque.size
	deque.count += 1

	return nil
}

func (deque *Deque[T]) AppendLeft(element T) error {
	if deque.count == deque.size {
		return errors.New("deque is full")
	}

	deque.left = (deque.size + deque.left - 1) % deque.size
	deque.elements[deque.left] = element
	deque.count += 1

	return nil
}

func (deque *Deque[T]) PopRight() (element T, err error) {
	if deque.count == 0 {
		err = errors.New("deque is empty")
		return
	}

	deque.right = (deque.size + deque.right - 1) % deque.size
	element = deque.elements[deque.right]
	deque.count -= 1

	return
}

func (deque *Deque[T]) PopLeft() (element T, err error) {
	if deque.count == 0 {
		err = errors.New("deque is empty")
		return
	}

	element = deque.elements[deque.left]
	deque.left = (deque.left + 1) % deque.size
	deque.count -= 1

	return
}

func (deque *Deque[T]) PeekRight() (element T, err error) {
	if deque.count == 0 {
		err = errors.New("deque is empty")
		return
	}

	position := (deque.size + deque.right - 1) % deque.size
	return deque.elements[position], nil
}

func (deque *Deque[T]) PeekLeft() (element T, err error) {
	if deque.count == 0 {
		err = errors.New("deque is empty")
		return
	}

	return deque.elements[deque.left], nil
}

func (deque *Deque[T]) Elements() []T {
	var elements []T

	position := deque.left
	count := int64(0)

	for count < deque.count {
		elements = append(elements, deque.elements[position])
		position = (position + 1) % deque.size
		count += 1
	}

	return elements
}
