package main

import "fmt"

type greeter interface {
	greet()
}

// 空のインタフェース型で引数を受け取る
func show(t interface{}) {
	// 型アサーション
	// 2つの値が返る
	_, ok := t.(japanese)
	// okを使って条件分岐
	if ok {
		fmt.Println("i am japanese")
	} else {
		fmt.Println("i am not japan")
	}

	// 型Switch
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
	fmt.Println("Hello")
}

// メイン関数
func main() {
	greeters := []greeter{japanese{}, american{}}

	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
