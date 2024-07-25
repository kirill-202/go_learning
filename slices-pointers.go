package main

import "fmt"

func slices_pointers() {
	names := [4]string{"Joe", "Ell", "Nick", "Tony"}
	fmt.Println(names)

	slice_names := names[:]

	slice_names[1] = "Goga"

	fmt.Println(names)
}