package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:= 1.0
	descrepancy:= 0.001
	for i:=0; i<10; i++ {
		old_z := z
		z -= (z*z -x)/ (2 *z)
		real_root:= math.Sqrt(x)
		if i> 1 && old_z <= z+descrepancy {
			fmt.Printf(
			"Number of iterations needed for good guess = %d\nThe best guess is %v\nThe real root via math module is %v", i+1, z, real_root)
			return z
		}
		fmt.Printf("Iteration %d: z = %v\n", i+1, z)
	}
	
	return z
}

func loop_funcs(root_value int) {
	fmt.Println(Sqrt(float64(root_value)))
}