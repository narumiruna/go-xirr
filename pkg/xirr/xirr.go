package xirr

import (
	"fmt"
	"go-xirr/pkg/types"
	"math"
)

type XIRR struct {
	Tolerance     float64
	Epsilon       float64
	MaxIterations int
	InitialValue  float64
}

func New() *XIRR {
	return &XIRR{
		Tolerance:     1e-10,
		Epsilon:       1e-10,
		MaxIterations: 100,
		InitialValue:  0.1,
	}

}

func (r *XIRR) Compute(cashflows types.CashFlowSlice) (float64, error) {
	if len(cashflows) == 0 {
		return 0, fmt.Errorf("empty slice")
	}

	durations := cashflows.Durations()
	values := cashflows.Values()

	if values.NonNegative() {
		return 0, fmt.Errorf("cash flows should have at least one negative value")
	}

	if values.NonPositive() {
		return 0, fmt.Errorf("cash flows should have at least one positive value")
	}

	// Newton's method
	x := r.InitialValue
	for range r.MaxIterations {
		y := f(x, durations, values)
		dy := df(x, durations, values)

		if math.Abs(dy) < r.Epsilon {
			break
		}

		dx := -y / dy

		if math.Abs(dx) < r.Epsilon {
			break
		}

		x += dx
	}

	return x, nil
}

func f(x float64, durations types.FloatSlice, amounts types.FloatSlice) float64 {
	sum := 0.0

	for i := range durations {
		sum += amounts[i] / math.Pow(1+x, durations[i])
	}

	return sum
}

func df(x float64, durations types.FloatSlice, amounts types.FloatSlice) float64 {
	sum := 0.0

	for i := range durations {
		sum += -amounts[i] * durations[i] / math.Pow(1+x, durations[i]+1)
	}

	return sum
}
