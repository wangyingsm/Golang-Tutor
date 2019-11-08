package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	a := [4]int{0, 1, 2, 3}
	fmt.Printf("Before calling noChange: %v\n", a)
	noChange(a)
	fmt.Printf("After called noChange: %v\n", a)
	// 此处可以很安全的忽略返回的错误，因为代码已经确定不会发生错误
	m, _ := maximum(10, -10, 100, 0)
	fmt.Printf("Max of serial numbers: %d\n", m)
	s := []int{20, 4, 5, -9, 3}
	var err error
	// 将slice打散进行调用，获得最大值，如果发生错误，将错误输出到标准错误中
	if m, err = maximum(s...); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Printf("Max of slice nubmers: %d\n", m)
}

// 参数永远传值，noChange函数不会改变数组的内容
func noChange(arr [4]int) {
	arr[3] = 100
}

// 定义一个寻找任意多个整数最大值的函数，返回最大值或错误
func maximum(x ...int) (int, error) {
	if len(x) == 0 {
		return 0, errors.New("empty arguments")
	}
	result := x[0]
	for i, v := range x {
		if i == 0 {
			continue
		}
		if v > result {
			result = v
		}
	}
	return result, nil
}
