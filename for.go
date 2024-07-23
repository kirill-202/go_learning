package main

import "fmt"

func for_go() {
	sum := 0
	for i :=0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println("This is the sum", sum)

}