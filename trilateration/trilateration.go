package trilateration

import "math"

type Space struct {
	A Point2
	B Point2
	C Point2
}

type Point2 struct {
	X float64
	Y float64
	R float64
}

func (ts Space) Solve() (x, y float64) {
	p2pDistance := getDistanceBetweenPoints(ts.A, ts.B)
	exx := (ts.B.X - ts.A.X) / p2pDistance
	exy := (ts.B.Y - ts.A.Y) / p2pDistance

	ix := exx * (ts.C.X - ts.A.X)
	iy := exy * (ts.C.Y - ts.A.Y)

	eyx := (ts.C.X - ts.A.X - ix*exx) / getVectorMagnitude(ts.A, ts.C, ix, exx, iy, exy)
	eyy := (ts.C.Y - ts.A.Y - iy*exy) / getVectorMagnitude(ts.A, ts.C, ix, exx, iy, exy)

	jx := eyx * (ts.C.X - ts.A.X)
	jy := eyy * (ts.C.Y - ts.A.Y)

	x = (math.Pow(ts.A.R, 2) - math.Pow(ts.B.R, 2) + math.Pow(p2pDistance, 2)) / (2 * p2pDistance)
	y = (math.Pow(ts.A.R, 2)-math.Pow(ts.C.R, 2)+math.Pow(iy, 2)+math.Pow(jy, 2))/2*jy - ix*x/jx
	return x, y
}

func getDistanceBetweenPoints(a Point2, b Point2) float64 {
	return math.Sqrt(math.Pow(b.X-a.X, 2) + math.Pow(b.Y-a.Y, 2))
}

func getVectorMagnitude(a Point2, b Point2, ix float64, exx float64, iy float64, exy float64) float64 {
	return math.Sqrt(math.Pow(b.X-a.X-ix*exx, 2) + math.Pow(b.Y-a.Y-iy*exy, 2))
}
