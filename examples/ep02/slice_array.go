package main

import "fmt"

func main() {
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(a[1:3])
	// type is slice []int
	fmt.Printf("Type of a[:]: %T\n", a[:])
	// len = 3, cap = 4
	fmt.Printf("Length of a[1:4]: %d, capacity of a[1:4]: %d\n",
		len(a[1:4]), cap(a[1:4]))
	// need expand capacity
	b := append(a[1:4], 5, 6, 7)
	fmt.Println(b)
	// type is slice []int
	fmt.Printf("Type of b: %T\n", b)
	// len = 6, cap = 8
	fmt.Printf("Length of b: %d, capacity of b: %d\n", len(b), cap(b))
}
