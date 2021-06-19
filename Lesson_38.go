package main

import "fmt"

// メイン関数
func main() {
    // 普通のforの繰り返し
    for i := 0; i < 10; i++ {
        if i == 3 {
            continue
        } else if i == 8 {
            break
        }

        fmt.Println(i)

    }

    // while文的にも作成可能
    n := 0
    for n < 10 {
        fmt.Println(n)
        n++
    }