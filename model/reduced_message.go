package model

import "errors"

type ReducedMessage [][]string

func (rm ReducedMessage) Reduce(f func(current, next []string) ([]string, error)) (result []string, err error) {
	if len([][]string(rm)) < 2 {
		return nil, errors.New("length too short")
	}
	for _, msg := range [][]string(rm) {
		if len(result) == 0 {
			result = msg
		} else {
			result, err = f(result, msg)
		}
	}
	return result, err
}
