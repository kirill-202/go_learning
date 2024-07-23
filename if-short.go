package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v:= math.Pow(x, n); v < lim {
		fmt.Println(v)
		return v
	}
	v:= 20
	fmt.Println(v)
	return lim
}
func if_short() {
	fmt.Println(
		pow(3,2, 10),
		pow(3,3, 20),
	)
}