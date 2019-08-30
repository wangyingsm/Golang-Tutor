package main

import (
	"fmt"
	"strconv"
	"time"
)

func compare(x, y interface{}) (int, bool) {
	switch x.(type) {
	case int:
		switch y.(type) {
		case int:
			return x.(int) - y.(int), true
		case string:
			yn, _ := strconv.Atoi(y.(string))
			return x.(int) - yn, true
		}
	case string:
		xn, _ := strconv.Atoi(x.(string))
		switch y.(type) {
		case int:
			return xn - y.(int), true
		case string:
			yn, _ := strconv.Atoi(y.(string))
			return xn - yn, true
		}
	}
	return 0, false
}

func main() {
	if result, ok := compare("10", 100); ok {
		fmt.Println(result)
	}
	if result, ok := compare(1000, "100"); ok {
		fmt.Println(result)
	}
	if result, ok := compare(time.Now(), "100"); ok {
		fmt.Println(result)
	}
}
