package math

type Element interface {
	Clone() Element
	Equals(element Element) bool
}

type Group interface {
	Neutral() Element
	Contains(element Element) bool

	Invert(element Element) (Element, error)
	Operation(left Element, right Element) (Element, error)
}
