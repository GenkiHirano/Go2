package main

import "fmt"

func main() {
	s := []int{2, 3, 8}

	for i, v := range s {
		fmt.Println(i, v)
	}

	for _, v := range s {
		fmt.Println(v)
	}

	m := map[string]int{"fujimoto": 200, "arita": 300}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
