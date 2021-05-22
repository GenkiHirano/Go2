package main

import "fmt"

type Student struct {
	name string
}

func (s Student) avg(math, english float64) {
	fmt.Println((math + english) / 2)
}

func main() {
	a001 := Student{"sato"}
	a001.avg(80, 70)
}