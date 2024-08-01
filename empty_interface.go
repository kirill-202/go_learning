package main

import "fmt"


func empty_interface() {
	var i interface{}
	describe_empty(i)

	i = 42
	describe_empty(i)

	i = "hello"
	describe_empty(i)
}

func describe_empty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
