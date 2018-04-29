package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "MoLi茉莉花"

	// 4D 6F 4C 69 E8 8C 89 E8 8E 89 E8 8A B1
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}

	fmt.Println()

	// ch is a rune
	// (0 4D) (1 6F) (2 4C) (3 69) (4 8309) (7 8389) (10 82B1)
	for i, ch := range s {
		fmt.Printf("(%d %X) ", i, ch)
	}

	fmt.Println()

	// UTF-8 count string
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	// [77 111 76 105 232 140 137 232 142 137 232 138 177]
	fmt.Println(bytes)
	for len(bytes) > 0 {
		// 取得每个字符的size
		ch, size := utf8.DecodeRune(bytes)
		//fmt.Println(ch, size)
		// slice删除size,继续遍历
		bytes = bytes[size:]
		//fmt.Println(ch)
		// %c 输出单个字符
		fmt.Printf("%c ", ch)
	}

	fmt.Println()

	// (0 4D) (1 6F) (2 4C) (3 69) (4 8309) (5 8389) (6 82B1)
	for i, ch := range []rune (s) {
		// %d 输出十进制整数; %X 输出十六进制数
		fmt.Printf("(%d %X) ", i, ch)
	}


}
