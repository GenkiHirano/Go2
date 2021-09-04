package main

import "fmt"

type greeter interface {
	greet()
}

type japanese struct{}
type american struct{}

func (ja japanese) greet() {
	fmt.Println("こんにちは")
}

func (us american) greet() {
	fmt.Println("Hello")
}

func main() {
	greeeters := []greeter{japanese{}, american{}}

	for _, greeter := range greeeters {
		greeter.greet()
	}
}
