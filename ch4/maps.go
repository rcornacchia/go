package main

import (
	"fmt"
)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}

	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}

	fmt.Println(ages)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
