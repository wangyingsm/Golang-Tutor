package main

import "fmt"

var aString = "hello world"
var aInteger int
var aBool bool

func main() {
	fmt.Println(aString)
	fmt.Println(aInteger + 123)
	if aBool {
		fmt.Println("不会打印这里")
	}
}
