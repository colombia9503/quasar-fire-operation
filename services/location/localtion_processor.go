package location

import (
	"errors"
	"github.com/colombia9503/quasar-fire-operation/trilateration"
)

func GetLocation(distances ...float32) (x, y float32, err error) {
	if dist := len(distances); dist < 3 && dist > 3 {
		return 0, 0, errors.New("distances array out of bounds, expected exactly 3 distances")
	}

	trilaterationSpace := trilateration.Space{
		A: trilateration.Point2{X: -500, Y: -200, R: float64(distances[0])},
		B: trilateration.Point2{X: 100, Y: -100, R: float64(distances[1])},
		C: trilateration.Point2{X: 500, Y: 100, R: float64(distances[2])},
	}

	x1, y1 := trilaterationSpace.Solve()
	return float32(x1), float32(y1), nil
}