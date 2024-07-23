package main

import "fmt"

func while() {
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Print(sum, ", ")
	}
	fmt.Println("total sum", sum)
}