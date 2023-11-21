package solver

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestSolver_NoRoots(t *testing.T) {
	solver := NewSolver(defaultEps)
	res, err := solver.Solve(1, 0, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{}, res)
}

func TestSolver_2Roots(t *testing.T) {
	solver := NewSolver(defaultEps)
	res, err := solver.Solve(1, 0, -1)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []float64{-1, 1}, res)
}

func TestSolver_1Root(t *testing.T) {
	solver := NewSolver(0.001)
	// 0 < D < epsilon
	res, err := solver.Solve(1, 2.0001, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{-2.0001 / 2}, res)
}

func TestSolver_ZeroCoefficient(t *testing.T) {
	solver := NewSolver(0.01)
	_, err := solver.Solve(0, 1, 1)
	assert.ErrorIs(t, err, errZeroCoefficient)

	_, err = solver.Solve(0.001, 1, 1)
	assert.ErrorIs(t, err, errZeroCoefficient)
}

// Not very necessary test case because we doesn't matter what float64 values Solve() will get
func TestSolver_Corrupted(t *testing.T) {
	solver := NewSolver(defaultEps)
	res, err := solver.Solve(math.Float64frombits(binary.LittleEndian.Uint64([]byte{0, 0, 0, 0, 0, 0, 240, 63})), -2, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{1}, res)
}
