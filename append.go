package main


func appends() {
	var s []int
	printSlice("s", s)

	s = append(s, 0)
	printSlice("s", s)
	
	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2,3,5,5,5,5,5)
	printSlice("s", s)
}