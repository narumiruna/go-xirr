package types

import "math"

type FloatSlice []float64

func (f FloatSlice) Max() float64 {
	m := -math.MaxFloat64
	for _, v := range f {
		m = math.Max(m, v)
	}
	return m
}

func (f FloatSlice) Min() float64 {
	m := math.MaxFloat64
	for _, v := range f {
		m = math.Min(m, v)
	}
	return m
}

func (f FloatSlice) NonNegative() bool {
	return f.Min() >= 0
}

func (f FloatSlice) NonPositive() bool {
	return f.Max() <= 0
}
