package main

import (
	"fmt"
)

func main() {
	var runes []rune
	for _, r := range "Hello World" {
		runes = append(runes, r)
	}
	d := make([]rune, 9)
	numElementsCopied := copy(d, runes)

	fmt.Printf("%q\n", d)
	fmt.Println(numElementsCopied)
}
