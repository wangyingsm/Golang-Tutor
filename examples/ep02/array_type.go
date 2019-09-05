package main

func main() {
	aArray := [5]int{0, 1, 2, 3, 4}
	// fmt.Println(sum(aArray)) 编译错误
	_ = aArray
}

func sum(a [6]int) int {
	result := 0
	for _, v := range a {
		result += v
	}
	return result
}
