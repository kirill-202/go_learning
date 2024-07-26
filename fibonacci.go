package main

import "fmt"

func fibonacci() func() int {

	start := 0
	incrementer := 1
	return func() int {
		start, incrementer = incrementer, start+incrementer

		return start
	}
}

func fibo() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}