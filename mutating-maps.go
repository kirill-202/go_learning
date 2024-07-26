package main

import "fmt"

func mutating_maps(){
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value", m["answer"])

	delete(m, "answer")
    fmt.Println("The value:", m["answer"])

	element, ok := m["answer"]
	fmt.Println("The value", element, "Present?", ok)
}