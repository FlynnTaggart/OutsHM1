package solver

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolver_NoRoots(t *testing.T) {
	solver := NewSolver()
	res, err := solver.Solve(1, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{}, res)
}

func TestSolver_2Roots(t *testing.T) {
	solver := NewSolver()
	res, err := solver.Solve(1, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{-1, 1}, res)
}

func TestSolver_1Root(t *testing.T) {
	solver := NewSolver()
	res, err := solver.Solve(1, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{-1}, res)
}

func TestSolver_ZeroCoefficient(t *testing.T) {
	solver := NewSolver()
	_, err := solver.Solve(0, 1, 1)
	assert.ErrorIs(t, err, errZeroCoefficient)
}
