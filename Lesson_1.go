package main

import "fmt"

func cal(x, y int) int {
		return (x / y)
	}

	func main() {
		result := cal(6, 3)
		fmt.Println(result)
	}