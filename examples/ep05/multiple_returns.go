package main

import (
	"errors"
	"fmt"
	"os"
)

// 仍然是整除求商和余数的例子
// 多返回值，命名返回值
func divmod(a, b int) (quotient, remainder int, err error) {
	if b == 0 {
		err = errors.New("divide by zero")
		return
	}
	quotient = a / b
	remainder = a % b
	return
}

func main() {
	q, r, err := divmod(100, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("100 divmod 3 = %d ... %d", q, r)
}
