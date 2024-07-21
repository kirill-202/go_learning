package main

import (
	"fmt"
	"math"
)

func type_conversion() {
	var x,y int = 3,4
	var f float64 = math.Sqrt(float64(x*x+y*y))
	var z uint = uint(f)
	var d string = string('0'+x) + "a"

	fmt.Println(x,y,z,f,d)
}
