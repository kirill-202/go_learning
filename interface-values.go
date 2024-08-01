package main

import (
	"fmt"
	"math"
)

type I_1 interface {
	M_1()
}


type T_1 struct {
	S string
}

func (t *T_1) M_1() {
	fmt.Println(t.S)
}

type F float64

func (f F) M_1() {
	fmt.Println(f)
}

func describe(i I_1) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func interface_values() {
	var i I_1

	i = &T_1{"Hello"}
	describe(i)
	i.M_1()

	i = F(math.Pi)
	describe(i)
	i.M_1()
}
