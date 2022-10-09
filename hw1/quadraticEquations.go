package hw1

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

const (
	float64EqualityThreshold = 1e-7
	ZeroAError               = "A argument is equal to 0"
	MaxFloatError            = "Max float %s argument %f"
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

	if Compare(math.Abs(a), math.MaxFloat64) {
		return nil, errors.New(fmt.Sprint(MaxFloatError, "a", a))
	}

	if Compare(math.Abs(b), math.MaxFloat64) {
		return nil, errors.New(fmt.Sprint(MaxFloatError, "b", b))
	}

	if Compare(math.Abs(c), math.MaxFloat64) {
		return nil, errors.New(fmt.Sprint(MaxFloatError, "c", c))
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
