package main

import (
	"fmt"
	"strings"
)

func tick_tack_toe() {

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	board[0][0] = "X"
	board[2][0] = "O"
	board[0][1] = "X"
	board[2][1] = "O"
	board[0][2] = "X"

	for i := 0; len(board) > i; i++ {
		fmt.Printf("%s\n", strings.Join(board[i], "  "))
	}
}