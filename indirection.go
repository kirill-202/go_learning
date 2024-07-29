package main

import "fmt"

type Vertex_7 struct {
	X,Y float64
}

func (ver *Vertex_7) Scale(scaler float64)  {
	ver.X = ver.X * scaler
	ver.Y = ver.Y * scaler
}

func ScalerFunc( ver *Vertex_7, scaler float64)  {
	ver.X *= scaler
	ver.Y *= scaler
}

func compare_pointers() {
	v:= Vertex_7{3,4}
	v.Scale(2)
	ScalerFunc(&v, 10)

	p := &Vertex_7{4, 3}
	p.Scale(3)
	ScalerFunc(p, 8)

	fmt.Println("this is v", v, "this is p", p)
}