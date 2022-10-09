package hw1

import (
	"testing"
)

func TestSolveDBelowZero(t *testing.T) {
	if len(Solve(1, 0, 1)) != 0 {
		t.Fatal("D = 0, no solution")
	}
}

func TestSolveDAboveZero(t *testing.T) {

	result := Solve(1, 0, -1)

	if len(result) != 2 {
		t.Fatal("D > 0, must be 2 solutions")
	} else {
		if !(Compare(result[0], -1) && Compare(result[1], 1)) {
			t.Fatalf("x1 must be -1, result[x1] = %f ,x2 must be 1, result[x2] = %f \n", result[0], result[1])
		}
	}
}

func TestSolveDEqualZero(t *testing.T) {

	result := Solve(1, 2, 1)

	if len(result) != 1 {
		t.Fatal("D = 0, must be 1 solutions")
	} else {
		if !(Compare(result[0], -1)) {
			t.Fatalf("x1 must be -1, result[x1] = %f \n", result[0])
		}
	}
}
