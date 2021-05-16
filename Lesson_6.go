package main

import "fmt"

func main() {
	hello := func(greeting string) {
		fmt.Println(greeting)
	}
	hello("Good morning")
}
