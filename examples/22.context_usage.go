package main

import (
	"context"
	"fmt"
	"os"
	"time"
)

func myRoutine(ctx context.Context) {
	fmt.Println("In myRoutine.")
	func(context.Context) {
		for i := 1; ; i++ {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("In innerRoutine.", i)
				time.Sleep(1 * time.Second)
			}
		}
	}(ctx)
	select {
	case <-ctx.Done():
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	myRoutine(ctx)
	select {
	case <-ctx.Done():
		os.Exit(0)
	}
}
