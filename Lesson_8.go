package main

import "fmt"

func cal(x, y int) (r int) {
	r = x + y
	return
}

func main() {
	result := cal(10, 5)
	fmt.Println(result)
}