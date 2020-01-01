package main

func main(){
	/* 定義式（どれも同じ意味） */

	// 定義式1
	var hoge int = 100
	var foo string = "Hello world"
	println(hoge, foo)

	// 定義式2
	var hogehoge = 100
	var foofoo string = "Hello world"
	println(hogehoge, foofoo)

	// 定義式3 (:=は定義する書き方なので同一変数に2回め以降は使えない)
	hogehogehoge := 100
	foofoofoo := "Hello world"
	println(hogehogehoge, foofoofoo)
}