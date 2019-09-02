package main

import "fmt"

// aString, aInteger, aBool are package variables

// declare aString as a string type variable and initialize it
// with a string constant "hello world".
var aString = "hello world"

// declare aInteger as a int type variable and compiler will
// initialize it with a int constant 123
var aInteger int

// declare aBool as a bool type variable and compiler will
// initialize it with a bool constant false
var aBool bool

func main() {
	fmt.Println(aString)
	fmt.Println(aInteger + 123)
	if aBool {
		fmt.Println("不会打印这里")
	}
}
