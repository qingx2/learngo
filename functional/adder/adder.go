package main

import "fmt"

// 非正统闭包写法
func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v

		return sum
	}
}

type iAdder func(int) (int, iAdder)

// 正统闭包写法
func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

// 闭包
func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + ... + %d = %d\n", i, a(i))
	}

	fmt.Println(" 正统闭包")

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0 + ... + %d = %d\n", i, s)
	}
}
