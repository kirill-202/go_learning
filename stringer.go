package main

import "fmt"

type Person struct {
	Age int
	Name string
}

func (p Person) String() string {
	return fmt.Sprintf("%v, %v years", p.Name, p.Age)
}

func stringer() {
	person := Person{33, "Kirill"}

	var person_2 Person
	person_2 = Person{11, "Sasha"}

	fmt.Println(person, person_2)
}