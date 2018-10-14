package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	time.Sleep(11100 * time.Millisecond)

	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
