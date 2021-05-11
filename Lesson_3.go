package main

import "fmt"

func cal(x, y int) (int, int) {
	a := x / y
	b := x * y
	return a, b
	}

	func main() {
		result01, result02 := cal(10, 5)
		fmt.Println(result01, result02)
	}