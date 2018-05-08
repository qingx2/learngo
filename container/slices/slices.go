package main

import (
	"fmt"
)

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	// features: slice view
	fmt.Println("Slice...")
	view(arr[:])

	// features: slice extend
	fmt.Println("Extending Slice...")
	extending(arr[:])

}

func extending(arr []int) {
	s1 := arr[2:6]
	fmt.Println(s1)
	// 2,3,4,5
	s2 := s1[3:5]
	fmt.Println(s2)
	// 5,6 tips: Slice view 时可以向后view 底层 array 扩展,但不能向前

	// s3 := s1[4] // Will fail :index out of range
	// fmt.Println(s3)

	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d \n",
		s1, len(s1), cap(s1))

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d \n",
		s2, len(s2), cap(s2))

	// append feature
	/*
		添加元素时如果超越 cap , 系统会重新分配更大的底层数组
		由于值传递的关系, 必须接收 append 的返回值
		ex: s = append(s, val)
	*/
	s10 := append(s2, 10)
	s11 := append(s10, 11)
	s12 := append(s11, 12)
	fmt.Println("s10, s11, s12 =", s10, s11, s12)
	// s11 and s12 no longer view arr
	fmt.Println("arr =", arr)
}

func view(arr []int) {
	fmt.Println("arr[2:6] = ", arr[2:6])
	fmt.Println("arr[:6] = ", arr[:6])
	fmt.Println("arr[2:] = ", arr[2:])
	fmt.Println("arr[:] = ", arr[:])
	fmt.Println("Slice view...")
	s := arr[2:6]
	fmt.Println(s)
	// 2,3,4,5
	s = s[:3]
	fmt.Println(s)
	// 2,3,4
	s = s[1:]
	fmt.Println(s)
	// 3,4
	s = arr[:]
	fmt.Println(s)
	// 0, 1, 2, 3, 4, 5, 6, 7
}

