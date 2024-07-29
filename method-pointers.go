package main

import (
	"fmt"
	"math"
)

type Vertex_6 struct {
	X, Y float64

}

func (v Vertex_6) Abs() float64 {
	return math.Sqrt(v.X*v.X+v.Y*v.Y)
}

func (v *Vertex_6) Scale(nunber float64)  {
 v.X = v.X * nunber
 v.Y = v.Y * nunber
}

func methods_pointers() {
	ver := Vertex_6{3,4}

	ver.Scale(10)
	fmt.Println(ver.Abs())
}