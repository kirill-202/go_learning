package main

import (
	"fmt"
	"math"
)

type my_float float64

func (float my_float) Abs() float64 {
	if float < 0 {
		return float64(-float)
	}
	return float64(float)
}

func methods_cont() {
	f := my_float(5.0)
	fmt.Println(f.Abs())

	f2 :=my_float(-math.Sqrt2)
	fmt.Println(f2.Abs())
}