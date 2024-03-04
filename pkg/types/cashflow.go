package types

import "time"

type CashFlow struct {
	Value float64
	Time  time.Time
}

type CashFlowSlice []CashFlow

func (c CashFlowSlice) Values() (s FloatSlice) {
	for _, item := range c {
		s = append(s, item.Value)
	}
	return s
}

func (c CashFlowSlice) Durations() (s FloatSlice) {
	startTime := c[0].Time
	for _, item := range c {
		s = append(s, item.Time.Sub(startTime).Hours()/24/365)
	}
	return s
}
