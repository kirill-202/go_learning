package main

import "fmt"

var pows = []int{1, 2, 4, 8, 16, 32, 64, 128}
func ranges() {
	for index, value := range pows {
		fmt.Printf("2**%d = %d\n", index, value)
	}
}