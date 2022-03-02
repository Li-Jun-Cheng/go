package main

import "fmt"

func main() {
	var s []int //定义了一个slices  Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	printSlice(s)
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 32)
	printSlice(s3)

	fmt.Println("Copying slice")
	copy(s2, s1) //把s1的元素“填充”到s2中
	printSlice(s2)

	//删除操作
	fmt.Println("Deleting elements from slice")
	s2 = append(s2[:3], s2[4:]...) //删除下标为3的元素
	printSlice(s2)

	fmt.Println("Popping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println(front)
	fmt.Println("Popping from back")
	tail := s2[len(s2)-1]
	s2 = s2[:len(s2)-1]
	fmt.Println(tail)
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("%v,len=%d, cap=%d\n", s, len(s), cap(s))
}
