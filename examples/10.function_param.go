package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3}
	fmt.Printf("Pointer of a: %p\n", &a)
	fmt.Println("Contents referenced by a:", a)
	transSlice(a)
}

func transSlice(b []int) {
	b[3] = 10
	fmt.Printf("Pointer of b: %p\n", &b)
	fmt.Println("Contents referenced by b:", b)
}
