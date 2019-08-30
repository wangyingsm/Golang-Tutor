package main

import (
	"fmt"
	"time"
)

func main() {
	// finished := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("In sub routine.")
		// finished <- true
	}()

	fmt.Println("In main routine.")
	time.Sleep(3 * time.Second)
	// <-finished
}
