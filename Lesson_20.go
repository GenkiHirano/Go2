// 宣言した後、値を代入パターン
var msg string
msg = "hello world"

// 宣言と代入を一緒にするパターン

var msg string = "Hello World"

// 宣言と代入を一緒にするパターン (型省略可能)
var msg = "Hello Hello"

// 宣言と代入を一緒にするパターン (var省略)
msg := "Super Hello"

// 複数の同じ型の変数を同時に定義
var a, b int
a, b = 10, 15

// 複数の同じ型の変数を同時に定義 (var省略)
a, b := 10, 14

// 複数の型違いの変数を同時に定義
var (
  c int
  d string
)

c = 20
d = "hoge"

// 複数の型違いの変数を同時に定義 (型省略パターン)

var (
  c = 20
  d = "hoge"
)
