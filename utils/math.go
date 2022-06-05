package utils

import "math"

var (
	NegativeInfinity = math.Inf(-1)
	PositiveInfinity = math.Inf(1)
)

const epsilon = 0.000001

func LessThanOrEqualFloat64(x, y float64) bool {
	diff := math.Abs(x - y)
	return diff < epsilon || x < y
}

func GreaterThanFloat64(x, y float64) bool {
	diff := math.Abs(x - y)
	return diff > epsilon && x > y
}
