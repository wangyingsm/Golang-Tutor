package main

import "fmt"

func main() {
	aString := "hello world"
	aInteger := 0
	aBool := false
	fmt.Println(aString)
	fmt.Println(aInteger + 123)
	if aBool {
		fmt.Println("不会打印这里")
	}
}
