package solver

import (
	"encoding/binary"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

const (
	bigEps             float64 = 0.001
	biggerEps          float64 = 0.01
	nearZeroA          float64 = 0.001
	bWithSmallFraction float64 = 2.0001
)

// Отсутствие корней
func TestSolver_NoRoots(t *testing.T) {
	solver := NewSolver(defaultEps)
	_, err := solver.Solve(1, 0, 1)
	assert.ErrorIs(t, err, errNoRoots)
}

// Два корня кратности один
func TestSolver_2Roots(t *testing.T) {
	solver := NewSolver(defaultEps)
	res, err := solver.Solve(1, 0, -1)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []float64{-1, 1}, res)
}

// Один корень кратности 2
// Учитываем, что дискриминант может быть не точно равен нулю
func TestSolver_1Root(t *testing.T) {
	solver := NewSolver(bigEps)
	// 0 < D < epsilon
	res, err := solver.Solve(1, bWithSmallFraction, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{-bWithSmallFraction / 2}, res)
}

// Нулевой коэффициент при x^2
// Также учитываем, что коэффициент может быть близко к нулю (в интервале epsilon), но не равен
func TestSolver_ZeroCoefficient(t *testing.T) {
	solver := NewSolver(biggerEps)
	_, err := solver.Solve(0, 1, 1)
	assert.ErrorIs(t, err, errZeroCoefficient)

	_, err = solver.Solve(nearZeroA, 1, 1)
	assert.ErrorIs(t, err, errZeroCoefficient)
}

// Not very necessary test case because we doesn't matter what float64 values Solve() will get
func TestSolver_Corrupted(t *testing.T) {
	solver := NewSolver(defaultEps)
	res, err := solver.Solve(math.Float64frombits(binary.LittleEndian.Uint64([]byte{0, 0, 0, 0, 0, 0, 240, 63})), -2, 1)
	assert.NoError(t, err)
	assert.Equal(t, []float64{1}, res)
}
