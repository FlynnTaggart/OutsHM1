package solver

import (
	"errors"
	"math"
)

const defaultEps = 0.0000001

var errZeroCoefficient = errors.New("zero a coefficient")

type Solver struct {
	eps float64
}

func NewSolver(eps float64) *Solver {
	if eps <= 0 {
		return &Solver{eps: defaultEps}
	}

	return &Solver{eps: eps}
}

func (s *Solver) Solve(a, b, c float64) ([]float64, error) {
	// -epsilon > a < epsilon
	if a > -s.eps && a < s.eps {
		return []float64{}, errZeroCoefficient
	}

	D := b*b - 4*a*c

	if D < 0 {
		return []float64{}, nil
	}

	// 0 > D < epsilon
	if D < s.eps {
		return []float64{-b / 2 * a}, nil
	}

	return []float64{
		(-b + math.Sqrt(D)) / 2 * a,
		(-b - math.Sqrt(D)) / 2 * a,
	}, nil
}
