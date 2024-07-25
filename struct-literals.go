package main

import "fmt"

type Vertex_3 struct {
	x,y int
}

var(
	v1  = Vertex_3{x:1}
	v2 = Vertex_3{}
	p = &Vertex_3{3,4}
)
	func str_literals() {
	fmt.Println(v1, p.x ,v2)
	}

 