package main

import (
	"fmt"
)

type greeter interface {
	greet()
}

func show(t interface{}) {
	_, ok := t.(japanese)

	if ok {
		fmt.Println("i am japanese")
	} else {
		fmt.Println("i am not japan")
	}

	switch t.(type) {
	case japanese:
		fmt.Println("僕は日本人だよ")
	default:
		fmt.Println("僕は日本人じゃないよ")
	}
}

type japanese struct{}
type american struct{}

func (ja japanese) greet() {
	fmt.Println("こんにちは！")
}

func (us american) greet() {
	fmt.Println("Hello!")
}

func main() {
	greeters := []greeter{japanese{}, american{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
