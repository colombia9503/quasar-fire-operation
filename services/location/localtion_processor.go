package location

import (
	"errors"
	"github.com/colombia9503/quasar-fire-operation/model"
	"github.com/colombia9503/quasar-fire-operation/trilateration"
)

func GetLocation(distances ...float64) (x, y float64, err error) {
	if dist := len(distances); dist > 3 || dist < 3 {
		return 0, 0, errors.New("distances array out of bounds, expected exactly 3 distances")
	}

	trilaterationSpace := trilateration.Space{
		A: trilateration.Point2{X: -500, Y: -200, R: distances[0]},
		B: trilateration.Point2{X: 100, Y: -100, R: distances[1]},
		C: trilateration.Point2{X: 500, Y: 100, R: distances[2]},
	}

	x1, y1 := trilaterationSpace.Solve()
	return x1, y1, nil
}

func GetShipLocation(tempSatellites []model.TempSatellite, data [3]model.SatelliteData) (x, y float64, err error) {
	if len(data) < len(tempSatellites) && len(data) > len(tempSatellites) {
		return 0, 0, errors.New("distances array out of bounds, expected exactly 3 distances")
	}

	trilaterationSpace := trilateration.Space{}
	for k, v := range tempSatellites {
		switch k {
		case 0:
			trilaterationSpace.A = trilateration.Point2{X: v.X, Y: v.Y, R: data[k].Distance}
		case 1:
			trilaterationSpace.B = trilateration.Point2{X: v.X, Y: v.Y, R: data[k].Distance}
		case 2:
			trilaterationSpace.C = trilateration.Point2{X: v.X, Y: v.Y, R: data[k].Distance}
		}
	}

	x1, y1 := trilaterationSpace.Solve()
	return x1, y1, nil
}
