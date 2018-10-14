package main

import (
	"fmt"
	"time"
)

func sum(s []int, c chan int, sleepTime time.Duration) {
	time.Sleep(sleepTime * time.Millisecond)
	fmt.Println("SUMMING:")
	fmt.Println(s)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c, time.Duration(200))
	go sum(s[len(s)/2:], c, time.Duration(4100))
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
