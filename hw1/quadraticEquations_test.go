package hw1

import (
	"fmt"
	"math"
	"testing"
)

func TestSolveDBelowZero(t *testing.T) {
	result, err := Solve(1, 0, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 0 {
		t.Fatal("D = 0, no solution")
	}
}

func TestSolveDAboveZero(t *testing.T) {

	result, err := Solve(1, 0, -1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 2 {
		t.Fatal("D > 0, must be 2 solutions")
	} else {
		if !(Compare(result[0], -1) && Compare(result[1], 1)) {
			t.Fatalf("x1 must be -1, result[x1] = %f ,x2 must be 1, result[x2] = %f \n", result[0], result[1])
		}
	}
}

func TestSolveDEqualZero(t *testing.T) {

	result, err := Solve(1, 2, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatal("D = 0, must be 1 solutions")
	} else {
		if !(Compare(result[0], -1)) {
			t.Fatalf("x1 must be -1, result[x1] = %f \n", result[0])
		}
	}
}

func TestSolveAEqualZero(t *testing.T) {
	var a float64 = 0
	_, err := Solve(a, 2, 1)
	if err != nil {
		if err.Error() != ZeroAError {
			t.Fatalf("Unknown error: %s \n", err.Error())
		}
	} else {
		t.Fatalf("A argument isn`t equal to 0, a = %f", a)
	}
}

func TestSolveDLessFloat64EqualityThreshold(t *testing.T) {
	result, err := Solve(1, 2.000000000025, 1)
	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatalf("D = 0, must be 1 solutions, and there are %d solutions \n", len(result))
	}
}

func TestSolveAMaxFloat64(t *testing.T) {
	a := math.MaxFloat64
	_, err := Solve(a, 2, 1)
	if err != nil {
		if err.Error() != fmt.Sprint(MaxFloatError, "a", a) {
			t.Fatalf("Unknown error: %s \n", err.Error())
		}
	} else {
		t.Fatalf("A argument isn`t equal to 0, a = %f", a)
	}
}

func TestSolveBMaxFloat64(t *testing.T) {
	b := math.MaxFloat64
	_, err := Solve(1, b, 1)
	if err != nil {
		if err.Error() != fmt.Sprint(MaxFloatError, "b", b) {
			t.Fatalf("Unknown error: %s \n", err.Error())
		}
	} else {
		t.Fatalf("B argument isn`t equal to 0, a = %f", b)
	}
}

func TestSolveCMaxFloat64(t *testing.T) {
	c := math.MaxFloat64
	_, err := Solve(1, 2, c)
	if err != nil {
		if err.Error() != fmt.Sprint(MaxFloatError, "c", c) {
			t.Fatalf("Unknown error: %s \n", err.Error())
		}
	} else {
		t.Fatalf("C argument isn`t equal to 0, a = %f", c)
	}
}
