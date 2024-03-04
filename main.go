package main

import (
	"fmt"
	"go-xirr/pkg/types"
	"go-xirr/pkg/xirr"
	"time"
)

func parse(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}

func main() {
	cashflows := types.CashFlowSlice{
		{
			Value: 100,
			Time:  parse("2023-01-01"),
		},
		{
			Value: 101,
			Time:  parse("2023-02-01"),
		},
		{
			Value: 102,
			Time:  parse("2023-04-01"),
		},

		{
			Value: 110,
			Time:  parse("2023-05-01"),
		},
		{
			Value: -425,
			Time:  parse("2023-06-01"),
		},
	}

	res, err := xirr.New().Compute(cashflows)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
