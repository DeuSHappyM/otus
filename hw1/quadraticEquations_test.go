package hw1

import (
	"testing"
)

func TestSolveD0(t *testing.T) {
	if len(Solve(1, 0, 1)) != 0 {
		t.Fatal("D = 0, no solution")
	}
}
