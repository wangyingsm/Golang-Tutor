package main

import "fmt"

var mem = make(map[int64]int64)

func fibonacci(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	if m, ok := mem[n]; ok {
		return m
	}
	result := fibonacci(n-1) + fibonacci(n-2)
	mem[n] = result
	return result
}

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println(fibonacci(int64(i)))
	}
}
