package main

import "fmt"


func defer_order() {
	fmt.Println("Started counting...")
	for i:=1; i<10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done counting.")
}