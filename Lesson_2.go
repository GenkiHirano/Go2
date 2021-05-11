package main

import "fmt"

func cal(x, y int) (int, int) {
		return (x / y), (x * y)
	}

	func main() {
		result01, result02 := cal(6, 3)
		fmt.Println(result01, result02)
	}