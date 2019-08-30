package main

import (
	"fmt"
	"strconv"
)

func myRoutine(aChan <-chan int, bChan <-chan string) {
	for {
		select {
		case i := <-aChan:
			fmt.Printf("An integer received: %d\n", i)
		case s := <-bChan:
			fmt.Printf("An message received: '%s'\n", s)
		}
	}
}

func main() {
	a := make(chan int, 10)
	defer close(a)
	b := make(chan string, 10)
	defer close(b)
	for i := 0; i < 10; i++ {
		go myRoutine(a, b)
	}
	for i := 100; i > 0; i-- {
		a <- i
		b <- strconv.Itoa(i)
	}
}
