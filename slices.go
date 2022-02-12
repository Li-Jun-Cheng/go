package main

import "fmt"

//切片
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	s1 := arr[2:]
	fmt.Println("arr[2:] =", s1)
	s2 := arr[:]
	fmt.Println("arr[:] =", s2)
	//slices是对arr的一个视图,可以理解为要看arr这个数组的什么区间的数据,所以就是在已经存在的数组区间上操作
	fmt.Println("After updateSlice(s1)")
	updateSlice(s1)
	fmt.Println(s1)
	fmt.Println(arr)

	fmt.Println("After updateSlice(s2)")
	updateSlice(s2)
	fmt.Println(s2)
	fmt.Println(arr)

	s2 = s2[:5]
	fmt.Println("s2[:5]", s2)
	s2 = s2[2:]
	fmt.Println("s2[2:]", s2)

	fmt.Println("Extending slice")
	arr[0], arr[2] = 0, 2
	s3 := arr[2:6]
	s4 := s3[3:5]
	fmt.Println(arr)
	fmt.Printf("s3=%v, len(s3)=%d, cap(s3)=%d\n",
		s3, len(s3), cap(s3))
	fmt.Printf("s4=%v, len(s4)=%d, cap(s4)=%d\n",
		s4, len(s4), cap(s4))
	fmt.Println()
	//没超过arr长度时会覆盖视图后面的值，超过cap长度了以后就会新开一个数组，将arr数组Copy过去在一个新的数组上操作，即系统会重新分配更大的底层数组
	//由于值传递的因素，必须要接受append的返回值
	s5 := append(s4, 10)
	s6 := append(s5, 11)
	s7 := append(s6, 12)
	fmt.Println("s5,s6,s7=", s5, s6, s7)
	fmt.Printf("cap(s5)=%d\n", cap(s5))
	fmt.Printf("cap(s6)=%d\n", cap(s6))
	fmt.Printf("cap(s7)=%d\n", cap(s7))
	fmt.Println("arr =", arr)
}
func updateSlice(s []int) {
	s[0] = 100
}
