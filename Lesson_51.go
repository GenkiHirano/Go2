package main

import "fmt"

// 複数の返り値を返す事ができる
func swap(a int, b int) (int, int) {
    return b, a
}

func main() {
    fmt.Println(swap(5, 2))

    // 関数もデータ型なので、変数に代入可能
    // その際は関数名はいらない
    f := func(a int, b int) (int, int) {
        return b, a
    }

    fmt.Println(f(3, 8))

    // 即時関数的な事も可能
    func(msg string) {
        fmt.Println(msg)
    }("Fujimoto")
}
