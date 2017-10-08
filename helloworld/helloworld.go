package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

/*
Go言語、セミコロンが無い代わりに改行に神経質です。
[code]func main() {
    fmt.Println("Hello, world!")
}[/code]
ここのfunc main() {の"{"は改行するとsyntax errorになるようです。
(https://golang.org/ でテスト可能)
コード書式がかっちり決まってるし、言語としては好きになれそう(inlove)
*/