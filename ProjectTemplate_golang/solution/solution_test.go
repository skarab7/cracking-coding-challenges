package solution

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	fmt.Println("Test!!")
	r := solve(1100)

	if r != -1 {
		t.Fatalf("Not as expected %v!", r)
	}
}
