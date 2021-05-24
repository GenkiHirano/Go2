package main

import "fmt"

type Student struct {
	name string
}

func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

func (s Student) junde(avg float64) (jundeResult float64) {
	if avg >= 60 {
		jundeResult = "passed"
	} else {
		jundeResult = "failed"
	}
	return
}

func main() {
	a001 := Student{"sato"}
	data := []float64{70, 65, 50, 90, 30}
	var avg float64 = a001.calAvg(data)
	result := a001.junde(avg)
	fmt.Println(avg)
	fmt.Println(a001.name + "" + result)
}
