package xirr

import (
	"fmt"
	"go-xirr/pkg/types"
	"math"
)

const (
	defaultEpsilon            = 1e-10
	defaultNumberOfIterations = 10
	defaultInitialValue       = 0.1
)

type XIRR struct {
	Epsilon       float64
	NumIterations int
	InitialValue  float64
}

func New() *XIRR {
	return &XIRR{
		Epsilon:       defaultEpsilon,
		NumIterations: defaultNumberOfIterations,
		InitialValue:  defaultInitialValue,
	}

}

func (x *XIRR) Compute(cashflows types.CashFlowSlice) (float64, error) {
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
	xn := x.InitialValue
	for range x.NumIterations {
		step := -f(xn, durations, values) / df(xn, durations, values)

		xn += step

		if math.Abs(step) < x.Epsilon {
			break
		}
	}

	return xn, nil
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
