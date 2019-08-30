package main

import "fmt"

func closure(a, b int) func(int) int {
	return func(c int) int {
		return a + b + c
	}
}

func main() {
	addThree := closure(10, 20)
	fmt.Println(addThree(30))
	fmt.Println(closure(100, 200)(300))
}
