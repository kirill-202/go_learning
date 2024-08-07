package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint((math.Sqrt(x)))
 }

func ifs() {
	fmt.Println(sqrt(4), sqrt(-4))
}