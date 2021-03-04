package main

import (
	"fmt"
	"github.com/colombia9503/quasar-fire-operation/services/location"
	"github.com/colombia9503/quasar-fire-operation/services/messages"
)

type Satellite struct {
	Name string
	Coordinate
}

type Coordinate struct {
	X int
	Y int
	R float32
}

//se halla el punto media trilateración

func main() {
	fmt.Println(location.GetLocation(100, 115.5, 142.7))
	fmt.Println(messages.GetMessage(
		[]string{"este", "", "", "mensaje", ""},
		[]string{"", "es", "", "", "secreto"},
		[]string{"este", "", "un", "", ""},
	))
}
