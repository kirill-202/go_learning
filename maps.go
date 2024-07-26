package main

import "fmt"

type Vertex_4 struct {
	Lat, Long float64
}


func maps(){
	m := make(map[string]Vertex_4)
	m["Bell Labs"] = Vertex_4{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}