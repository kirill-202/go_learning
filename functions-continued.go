package main

import (
	"fmt"
	"strconv"
	"reflect"
)

func add_str(x,y int) string {
	string_num := strconv.Itoa(x+y)
	return string_num
}

func functions_continued() {
	fmt.Println("This will return string instead of integer", add_str(55,56), "The type of the integer is", reflect.TypeOf(add_str(55,56)))
}