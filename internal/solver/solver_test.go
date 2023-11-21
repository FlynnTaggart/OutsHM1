package solver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolve_NoRoots(t *testing.T) {
	assert.Equal(t, []float64{}, Solve(1, 0, 1))
}
