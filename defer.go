package main

import "fmt"

func defer_func() {
	defer fmt.Println("Fuck you")
	defer fmt.Println("world")
	fmt.Println("Hello")
}