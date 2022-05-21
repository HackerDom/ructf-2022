package math

import "errors"

type SymmetricGroup[T PermutationElement] struct {
	length int
}

func NewSymmetricGroup[T PermutationElement](length int) *SymmetricGroup[T] {
	return &SymmetricGroup[T]{
		length: length,
	}
}

func (group *SymmetricGroup[T]) Length() int {
	return group.length
}

func (group *SymmetricGroup[T]) Neutral() Element {
	elements := make([]T, group.length)

	for i := 0; i < group.length; i++ {
		elements[i] = T(i)
	}

	return &Permutation[T]{
		elements: elements,
	}
}

func (group *SymmetricGroup[T]) Contains(element Element) bool {
	permutation, ok := element.(*Permutation[T])
	if !ok {
		return false
	}

	return len(permutation.elements) == group.length
}

func (group *SymmetricGroup[T]) Invert(element Element) (Element, error) {
	permutation, ok := element.(*Permutation[T])
	if !ok {
		return nil, errors.New("element should be permutation")
	}

	if !group.Contains(permutation) {
		return nil, errors.New("group does not contain permutation")
	}

	elements := make([]T, group.length)

	for i := 0; i < group.length; i++ {
		elements[permutation.elements[i]] = T(i)
	}

	return &Permutation[T]{
		elements: elements,
	}, nil
}

func (group *SymmetricGroup[T]) Operation(left Element, right Element) (Element, error) {
	if !group.Contains(left) {
		return nil, errors.New("group does not contain left")
	}

	if !group.Contains(right) {
		return nil, errors.New("group does not contain right")
	}

	leftPermutation := left.(*Permutation[T])
	rightPermutation := right.(*Permutation[T])

	elements := make([]T, len(rightPermutation.elements))

	for i, element := range rightPermutation.elements {
		elements[i] = leftPermutation.elements[element]
	}

	return &Permutation[T]{
		elements: elements,
	}, nil
}
