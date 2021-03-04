package model

type ReducedMessage [][]string

func (rm ReducedMessage) Reduce(f func(current []string, next []string) ([]string, error)) (result []string, err error) {
	for _, msg := range [][]string(rm) {
		if len(result) == 0 {
			result = msg
		} else {
			result, err = f(result, msg)
		}
	}
	return result, err
}