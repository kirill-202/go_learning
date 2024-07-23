package main

import (
	"fmt"
	"time"
)

func check_day() {
	fmt.Println("Is it Sunday?")
	today:= time.Now().Weekday()
	fmt.Println(today)
	
	switch time.Sunday {
	case today + 4:
		fmt.Println("Tomorrow")
	case today - 1:
		fmt.Println("Yesterday")
	case today - 2:
		fmt.Println("Today")
	default:
		fmt.Println("Too far away")
	}
}