package main

import "fmt"

type Student struct {
	name string
}

func (s Student) avg(math, english float64) (avgResult float64) {
	avgResult = (math + english) / 2
	return
}

func main() {
	a001 := Student{"sato"}
	fmt.Println(a001.avg(80, 70))
}
