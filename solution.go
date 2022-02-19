package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type SidesCount int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func triangleArea(len float64) float64 {
	return len * len * math.Sqrt(3.0) / 4.0
}

func squareArea(len float64) float64 {
	return len * len
}

func circleArea(len float64) float64 {
	return math.Pi * len * len
}

func CalcSquare(sideLen float64, sidesNum SidesCount) float64 {
	switch sidesNum {
	case SidesCircle:
		return circleArea(sideLen)
	case SidesTriangle:
		return triangleArea(sideLen)
	case SidesSquare:
		return squareArea(sideLen)
	}
	return 0
}
