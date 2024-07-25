package main

import "fmt"

func making() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 5, 10)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[3:7]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s, capacity=%d, length=%d for slice %v\n", s, cap(x), len(x), x)
}