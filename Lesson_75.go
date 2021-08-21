package main

import "fmt"

func main() {
	var intface interface{} = "hello"

	variable := intface.(string)
	fmt.Println(variable) //=> hello

	variable, ok := intface.(string)
	fmt.Println(variable, ok) //=> hello true

	float, ok := intface.(float64)
	fmt.Println(float, ok) //=> 0 false
	//格納失敗が、成功したかの有無を確かめるokが存在するのでエラーにはならない。

	float = intface.(float64)
	fmt.Println(float) //=> panic: interface conversion: interface {} is string, not float64
	//成功したかの有無を確かめるokが存在しないのでエラーが発生する。
}
