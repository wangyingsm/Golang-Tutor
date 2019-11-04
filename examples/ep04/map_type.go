package main

// 使用动态规划算法，求前50个斐波那契数列值的程序

import "fmt"

// 定义一个包内全局的变量mem，用来存储斐波那契数列求得的结果
var mem = make(map[int64]int64)

func fibonacci(n int64) int64 {
	if n == 1 || n == 2 {
		return 1	// 斐波那契数列头两项
	}
	// 如果map中已经有算好的数列项值，则直接使用
	// 你可以试着将下面的程序块注释掉，再运行一遍程序，观察运行时间
	if m, ok := mem[n]; ok {
		return m
	}
	// 根据斐波那契数列定义递归求值
	result := fibonacci(n-1) + fibonacci(n-2)
	// 求得的值使用当前数列项序号作为key，保存在map当中
	mem[n] = result
	return result
}

func main() {
	for i := 1; i <= 50; i++ {
		fmt.Println(fibonacci(int64(i)))
	}
}
