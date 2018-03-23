package main

import "fmt"

// 数组
func main() {
	var array1 [5]int // print [0 0 0 0 0]

	array2 := [3]int{1, 3, 5}

	array3 := [...]int{1,2,3,4,5}

	var grid [2][3] bool // print [[false false false] [false false false]]

	fmt.Println(array1, array2,array3,grid)
}
