package trilateration

import (
	"log"
	"testing"
)

func TestTrilateration(t *testing.T) {
	space := Space{
		A: Point2{R: 100, X: -500, Y: -200},
		B: Point2{R: 115.5, X: 100, Y: -100},
		C: Point2{R: 142.7, X: 500, Y: 100},
	}

	x, y := space.Solve()
	log.Println(x, y)
	if int(x) != 301 && int(y) != 1 {
		t.Error("Test failed, arguments dont match")
	}
}
