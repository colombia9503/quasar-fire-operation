package messages

import (
	"log"
	"testing"
)

func TestGetMessageSuccessful(t *testing.T) {
	if result, err := GetMessage(
		[]string{"Hola", "", "estas", ""},
		[]string{"", "como", "estas", ""},
		[]string{"Hola", "", "", "?"},
		[]string{"", "", "estas", "?"},
	); err != nil && !(result == "Hola como estas?") {
		t.Error("That's not supposed to return an exception")
	} else {
		log.Println(result)
	}
}

func TestGetMessageError(t *testing.T) {
	if _, err := GetMessage(
		[]string{"Hola", "", "estas", ""},
		[]string{"EXPECTATIVA", "como", "estas", ""},
		[]string{"Expectativa", "", "", "?"},
		[]string{"", "", "estas", "?"},
	); err == nil {
		t.Error("Expected an error")
	} else {
		log.Println(err)
	}
}

func TestGetMessageErrorDiffLengthTraces(t *testing.T) {
	if _, err := GetMessage(
		[]string{"Hola", "", "estas", ""},
		[]string{"", "como"},
		[]string{"?"},
		[]string{"", "estas", "?"},
	); err == nil {
		t.Error("Expected an error")
	} else {
		log.Println(err)
	}
}
