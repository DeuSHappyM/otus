package hw1

import (
	"errors"
	"math"
	"sort"
)

const (
	float64EqualityThreshold = 1e-7
	ZeroAError               = "A argument is equal to 0"
)

func Compare(x, y float64) bool {
	if math.Abs(x-y) < float64EqualityThreshold {
		return true
	}
	return false
}

func Solve(a, b, c float64) ([]float64, error) {

	if Compare(a, 0) {
		return nil, errors.New(ZeroAError)
	}

	D := b*b - 4*a*c

	switch {
	case D < 0:
		return []float64{}, nil
	case Compare(D, 0):
		return []float64{(-b + math.Sqrt(D)) / 2 * a}, nil
	case D > 0:
		result := []float64{(-b + math.Sqrt(D)) / 2 * a, (-b - math.Sqrt(D)) / 2 * a}
		sort.Float64s(result)
		return result, nil
	default:
		return []float64{}, nil
	}
}
