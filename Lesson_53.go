package main

import "fmt"

// メイン関数
func main() {
    a := [5]int{2, 10, 8, 15, 4}

    b := a[2:4] // [8, 15]
    c := a[2:]  // [8, 15, 4]
    d := a[:4]  // [2, 10, 8, 15]
    e := a[:]   // [2, 10, 8, 15, 4]
    fmt.Println(b)
    fmt.Println(c)
    fmt.Println(d)
    fmt.Println(e)

    // スライスの長さを取得
    fmt.Println(len(c)) // 3

    // スライスの最大容量を取得
    fmt.Println(cap(c))
}
