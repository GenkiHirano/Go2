package main

import "fmt"

// 構造体を定義
type user struct {
	// フィールド
	name  string
	score int
}

// メイン関数
func main() {
	// newでuser構造体分の領域を確保して、初期化して、そのポインタを返す
	u1 := new(user)
	// ポインタが返ってきているので下記の書き方でもOK
	// (*u).name = "fujimoto"
	u1.name = "fujimoto"
	u1.score = 200

	fmt.Println(u1)

	// 初期化する時に直接値を与える事も可能
	// こちらの場合はポインタではない構造体のデータが入ってくる。
	u2 := user{name: "arita", score: 300}
	fmt.Println(u2)
}
