package main

import "fmt"

func nil_slices() {
	var s []int
	fmt.Println(cap(s), len(s), s)

	if s == nil {
		fmt.Println("s is nil!!!!")
	}
}