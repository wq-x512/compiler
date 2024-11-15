package main

import (
	"fmt"
	"pkg/lexer"
)

func main() {
	stream := "int a = 10;x + y == 10.0? union"
	lexer.Start(stream)
	Tokens := lexer.GetToken()
	for _, i := range Tokens {
		fmt.Println(i)
	}
	//str := "你好"
	//for k, v := range str {
	//	fmt.Println(k, v >= '\u4e00' && v <= '\u9fff')
	//}
}
