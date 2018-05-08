package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内部变量只能使用 var 声明
var aa = false

var (
	bb = "hi"
	cc = 23333
)

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "hello a b c "
	fmt.Println(a, s, b)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "shorter"
	fmt.Println(a, b, c, s)
}

// 欧拉公式
func euler() {
	// c :=3+4i
	// abs := cmplx.Abs(c)
	// fmt.Println(abs)
	// fmt.Println(fmt.Sprintf("%T",abs))

	// print (0+1.2246467991473515e-16i)
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	// the same
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	// print (0.000+0.000i)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

// 强制类型转换
func triangle() {
	a, b := 3, 4
	c := calcTriangle(a, b)
	fmt.Println(c)
}

func calcTriangle(a int, b int) int {
	var c int
	// math.Sqrt 需要转 float64
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

// 定义常量
func consts() {
	const (
		filename = "adf.txt"
		a, b     = 3, 4
	)
	var c int
	// const常量数值可作为各种类型使用,所以此处的 math.Sqrt 不必像 triangle()一样强制转换为 flot64
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp        = iota
		javascrpit
		_
		python
		java
		golang
		php
	)
	const (
		aa = 1 << (10 * iota)
		bb
		cc
		dd
		ee
	)

	fmt.Println(cpp, javascrpit, python, java, golang, php)
	fmt.Println(aa, bb, cc, dd, ee)
	fmt.Println(1 << (10 * 1))
}

// 左移、右移运算
func shifting(a int) int {
	a = a << 1
	a = a >> 1
	return a
}

func main() {
	fmt.Println("Hello GoPATH")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, bb, cc)

	euler()
	triangle()

	consts()
	enums()

	fmt.Println(shifting(3))
}
