package main

import "fmt"

func main() {
	func(greeting string) {
		fmt.Println(greeting)
	}("Good evening")
}
