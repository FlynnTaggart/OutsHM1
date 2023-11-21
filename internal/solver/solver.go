package solver

import (
	"errors"
	"math"
)

const defaultEps = 0.0000001

var (
	errZeroCoefficient = errors.New("zero a coefficient")
	errNoRoots         = errors.New("under zero discriminant and the equation has no roots")
)

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
	// -epsilon > a < epsilon, нулевой коэффициент при x^2
	if a > -s.eps && a < s.eps {
		return []float64{}, errZeroCoefficient
	}

	D := b*b - 4*a*c

	// дискриминант меньше нуля
	if D < 0 {
		return []float64{}, errNoRoots
	}

	// 0 > D < epsilon, дискриминант равен нулю, один корень кратности 2
	if D < s.eps {
		return []float64{-b / 2 * a}, nil
	}

	// стандартные два корня
	return []float64{
		(-b + math.Sqrt(D)) / 2 * a,
		(-b - math.Sqrt(D)) / 2 * a,
	}, nil
}
