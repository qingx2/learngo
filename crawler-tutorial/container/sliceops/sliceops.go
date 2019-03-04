package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d,cap=%d \n", s, len(s), cap(s))
}

func main() {
	fmt.Println("\n Creating Slice")
	var s []int // Zero value for slice is nil

	for i := 0; i < 100; i++ {
		printSlice(s)

		s = append(s, 2*i+1)
	}

	fmt.Println(s)

	// [2 4 6 8], len=4,cap=4
	s1 := []int{2, 4, 6, 8}

	// [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0], len=16,cap=16
	s2 := make([]int, 16)

	// [0 0 0 0 0 0 0 0 0 0], len=10,cap=32
	s3 := make([]int, 10, 32)

	printSlice(s1)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("\n Copying Slice")
	// [2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0], len=16,cap=16
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("\n Deleting elements from Slice")
	// [2 4 6 0 0 0 0 0 0 0 0 0 0 0 0], len=15,cap=16
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)

	fmt.Println("\n Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	printSlice(s2)

	fmt.Println("\n Popping from back")
	back := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(back)
	printSlice(s2)

}
