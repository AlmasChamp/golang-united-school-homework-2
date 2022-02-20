package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type Num int

const (
	Triangle Num = 3
	Square   Num = 4
	Circle   Num = 0
)

func CalcSquare(sideLen float64, sidesNum Num) float64 {
	if sidesNum == Triangle {
		return (math.Sqrt(3) / 4) * (sideLen * sideLen)
	} else if sidesNum == Square {
		return sideLen * sideLen
	} else if sidesNum == Circle {
		return math.Pi * (sideLen * sideLen)
	}
	return 0
}
