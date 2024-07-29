package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func interfaces_implicit() {
	my_type := T{"lol"}
	my_type.M()
}