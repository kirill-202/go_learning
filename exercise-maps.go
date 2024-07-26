package main

import (
	"golang.org/x/tour/wc"
	"fmt"
	"strings"
)

func word_count(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for i := 0; i< len(words); i++ {
		_, ok := counts[words[i]]
		if ok {
			counts[words[i]] += 1
			fmt.Println("Word was present, counter incremented")
		} else {
			counts[words[i]] = 1
			fmt.Println("Word was not present, counter set to 1")
		}
	}
	fmt.Println(counts)
	return counts
}

func testing_count(){
	wc.Test(word_count)
}