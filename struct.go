package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func structs() {
	vertex := Vertex{1,2}
	fmt.Println(vertex)
}