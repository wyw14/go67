package patch

import "errors"

func Apply(a []string, op string, index int, value string) ([]string, error) {
	limit := len(a) - 1
	if op == "add" {
		limit = len(a)
	}
	if index < 0 || index > limit {
		return nil, errors.New("index")
	}
	if op == "replace" {
		a[index] = value
		return a, nil
	}
	a = append(a, "")
	copy(a[index+1:], a[index:])
	a[index] = value
	return a, nil
}
