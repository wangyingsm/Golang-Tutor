package main

import "fmt"

// 一个简单的闭包函数，主要为了说明语法
func closure(a, b int) func(int) int {
	return func(c int) int {
		return a + b + c
	}
}

func main() {
	addThree := closure(10, 20)
	// 输出60
	fmt.Println(addThree(30))
	// 输出600
	fmt.Println(closure(100, 200)(300))
}
