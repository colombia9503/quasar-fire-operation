package messages

import (
	"errors"
	"github.com/colombia9503/quasar-fire-operation/model"
	"strings"
)

func GetMessage(messages ...[]string) (msg string, err error) {
	var msgArr model.ReducedMessage = messages

	result, err := msgArr.Reduce(func(current []string, next []string) (arr []string, err error) {
		if len(current) != len(next) {
			arr = nil
			err = errors.New("cannot determine the message, arrays with different length")
		} else {
			arr = []string{}
			err = nil
			for k, v := range current {
				switch {
				case v == next[k]:
					fallthrough
				case v == "" && next[k] != "":
					arr = append(arr, next[k])
				case next[k] == "" && v != "":
					arr = append(arr, v)
				default:
					arr = nil
					err = errors.New("cannot determine the message, incompatible trace between messages")
					break
				}
			}
		}
		return arr, err
	})

	return strings.Join(result, " "), err
}
