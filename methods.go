package main 

import (
	"fmt"
	"math"
)

type Vertex_5 struct {
	x, y float64
}

func (vertex Vertex_5) Abs() float64 {
	return math.Sqrt(vertex.x*vertex.x + vertex.y+vertex.y)
}
func methods() {
	v := Vertex_5{3,4}
	fmt.Println(v.Abs())

}