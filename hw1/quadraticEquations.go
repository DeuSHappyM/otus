package hw1

import (
	"math"
	"sort"
)

const float64EqualityThreshold = 1e-7

func Compare(x, y float64) bool {
	if math.Abs(x-y) > float64EqualityThreshold {
		return false
	} else {
		return true
	}
}

func Solve(a, b, c float64) []float64 {

	D := b*b - 4*a*c

	switch {
	case D < 0:
		return []float64{}
	case D > 0:
		result := []float64{(-b + math.Sqrt(D)) / 2 * a, (-b - math.Sqrt(D)) / 2 * a}
		sort.Float64s(result)
		return result
	case Compare(D, 0):
		return []float64{(-b + math.Sqrt(D)) / 2 * a}
	default:
		return []float64{}
	}
}
