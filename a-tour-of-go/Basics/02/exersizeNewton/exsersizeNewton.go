package exersizeNewton

import "math"

func Sqrt(x float64) float64 {
	z := 1.0
	eps:= 1e-15

	for math.Abs(x-z*z) > eps {
		z -= (z*z - x) / (2 * z)
	}

	return z
}
