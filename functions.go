package main

import "fmt"

func add(x int, y int) int {
	return x+y
}

func functions() {
	fmt.Println("This is the result of add function =", add(55, 66)) 
}