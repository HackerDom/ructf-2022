package math

import (
	"errors"
)

type PermutationElement interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Permutation[T PermutationElement] struct {
	elements []T
}

var _ Element = &Permutation[uint64]{}

func NewPermutation[T PermutationElement](elements []T) (*Permutation[T], error) {
	new := &Permutation[T]{
		elements: make([]T, len(elements)),
	}

	copy(new.elements, elements)

	if !new.IsValid() {
		return nil, errors.New("permutation is not valid")
	}

	return new, nil
}

func (permutation *Permutation[T]) Length() int {
	return len(permutation.elements)
}

func (permutation *Permutation[T]) Elements() []T {
	elements := make([]T, len(permutation.elements))

	copy(elements, permutation.elements)

	return elements
}

func (permutation *Permutation[T]) IsValid() bool {
	var counts = map[T]int{}

	for _, element := range permutation.elements {
		if _, ok := counts[element]; !ok {
			counts[element] = 0
		}

		counts[element] += 1
	}

	if len(counts) != len(permutation.elements) {
		return false
	}

	for i := range permutation.elements {
		if _, ok := counts[T(i)]; !ok {
			return false
		}
	}

	return true
}

func (permutation *Permutation[T]) Clone() Element {
	new, _ := NewPermutation(permutation.elements)
	return new
}

func (permutation *Permutation[T]) Equals(element Element) bool {
	other, ok := element.(*Permutation[T])
	if !ok {
		return false
	}

	if len(permutation.elements) != len(other.elements) {
		return false
	}

	for i := range permutation.elements {
		if permutation.elements[i] != other.elements[i] {
			return false
		}
	}

	return true
}
