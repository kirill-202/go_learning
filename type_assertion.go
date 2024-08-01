package main

import "fmt"

func type_assert() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)


	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(bool)
	fmt.Println(f, ok)

	f = i.(bool)
	fmt.Println(f)
}