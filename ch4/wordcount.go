package main

import (
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	counts := make(map[string]int)
	var curr string
	space := ' '

	for i, letter := range s {
		if letter == space {
			counts[curr]++
			curr = ""
		} else if i+1 == len(s) {
			curr += string(letter)
			counts[curr]++
		} else {
			curr += string(letter)
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}
