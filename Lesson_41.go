package main

import "fmt"

// 構造体を定義
type user struct {
	// フィールド
	name  string
	score int
}

// レシーバで構造体とメソッド(関数)を紐付け
func (u user) show() {
	fmt.Printf("name: %s, socre: %d\n", u.name, u.score)
}

// デフォルトではuserには値渡し(コピー)でメソッドに渡される。
// *をつける事で参照渡しになる
func (u *user) hit() {
	u.score++
}

// メイン関数
func main() {
	u := user{name: "fujimoto", score: 500}
	u.hit()
	u.show()
}
