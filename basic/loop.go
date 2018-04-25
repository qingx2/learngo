package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)
// for, if 后面的条件没有括号
// if 条件里可以定义变量
// 没有 while 关键字
// switch 不需要 break, 并且可以直接 case 多个条件


func converToBin(n int) string {
	result := ""
	// 可省略初始条件,相当于while,可以设置递增/结束条件
	for ; n > 0; n /= 2 {
		lsb := n % 2
		fmt.Println(lsb, result)
		result = strconv.Itoa(lsb) + result
	}

	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// 相当于 while
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever(){
	// 死循环 while
	for   {
		fmt.Println("aaacc")
	}
}

func main() {
	fmt.Println(
		//converToBin(5), // 101
		converToBin(1), // 1101
		//converToBin(1112312312321321),
	)

	printFile("aaa.txt")

	forever()
}
