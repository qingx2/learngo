package main

import (
	"io/ioutil"
	"fmt"
)

func ifReadFile() {
	const filename = "aaa.txt"
	// if 的条件里赋值的变量(contents,err)作用域生命周期只在这个 if 语句里
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong socre: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	default:
		panic(fmt.Sprintf("Wrong socre: %d", score))
	}

	return g
}

func main() {
	ifReadFile()
	fmt.Println(grade(101))
}
