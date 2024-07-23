package main

import "fmt"

func go_next() {
	sum := 999
	for ; sum < 1000; {
		sum += sum
		fmt.Println("new sum", sum)
	}
}