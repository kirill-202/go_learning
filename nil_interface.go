package main

import "fmt"

type I_4 interface {
	M()
	
}

func nil_interfaces() {
	var inter_face I_4
	describe_2(inter_face)
	inter_face.M()
}

func describe_2(i I_4) {
	fmt.Printf("(%v, %T)\n", i, i)
}
