package location

import (
	"errors"
	"github.com/colombia9503/quasar-fire-operation/model"
	"github.com/colombia9503/quasar-fire-operation/trilateration"
)

func GetLocation(distances ...float32) (x, y float32, err error) {
	if dist := len(distances); dist > 3 || dist < 3 {
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

func GetShipLocation(tempSatellites []model.TempSatellite, data [3]model.SatelliteData) (x, y float32, err error) {
	if len(data) < len(tempSatellites) && len(data) > len(tempSatellites) {
		return 0, 0, errors.New("distances array out of bounds, expected exactly 3 distances")
	}

	trilaterationSpace := trilateration.Space{}
	for k, v := range tempSatellites {
		switch k {
		case 0:
			trilaterationSpace.A = trilateration.Point2{X: float64(v.X), Y: float64(v.Y), R: float64(data[k].Distance)}
		case 1:
			trilaterationSpace.B = trilateration.Point2{X: float64(v.X), Y: float64(v.Y), R: float64(data[k].Distance)}
		case 2:
			trilaterationSpace.C = trilateration.Point2{X: float64(v.X), Y: float64(v.Y), R: float64(data[k].Distance)}
		}
	}

	x1, y1 := trilaterationSpace.Solve()
	return float32(x1), float32(y1), nil
}
