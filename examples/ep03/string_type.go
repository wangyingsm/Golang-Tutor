package main

import "fmt"

func main() {
	welcome := "你好，世界"
	fmt.Println(welcome[:3])
	// welcome[7] = ':' 不能改变string的值
	fmt.Println("Bytes of welcome:", len(welcome))
	for _, r := range []rune(welcome) {
		fmt.Printf("%s\n", string(r))
	}
	fmt.Println("Length of welcome:", len([]rune(welcome)))
}
