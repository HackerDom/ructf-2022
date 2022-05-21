package math

import "errors"

func Exponentiation(group Group, element Element, power int64) (result Element, err error) {
	if !group.Contains(element) {
		err = errors.New("group does not contain element")
		return
	}

	result = group.Neutral()

	if power < 0 {
		element, err = group.Invert(element)
		if err != nil {
			return
		}

		power = -power
	}

	for power > 0 {
		if power&1 == 1 {
			result, err = group.Operation(result, element)
			if err != nil {
				return
			}
		}

		element, err = group.Operation(element, element)
		if err != nil {
			return
		}

		power >>= 1
	}

	return
}
