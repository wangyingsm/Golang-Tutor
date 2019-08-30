package main

import "fmt"

// Celcius 摄氏度
type Celcius float32

func main() {
	var c Celcius = 36.985
	fmt.Println(c)
}

func (c Celcius) String() string {
	return fmt.Sprintf("%.2f ℃", float64(c))
}
