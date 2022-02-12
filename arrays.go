package main

import "fmt"

//go语言一般不直接使用数组，而是切片
func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int         //定义长度为5的数组
	arr2 := [3]int{1, 3, 5} //如果用:=就需要把数据的内容直接给它才行
	arr3 := [...]int{2, 4, 6, 8, 10}
	var grid [4][5]int
	fmt.Println(arr1, arr2, arr3)
	fmt.Println(grid)
	fmt.Println("printArray(arr1)")
	printArray(&arr1)
	fmt.Println("printArray(arr3)")
	printArray(&arr3)
	fmt.Println("arr1 and arr3")
	fmt.Println(arr1, arr3)
	//虽然和Java都一样是值传递，但是当数据作为参数这里和Java的值传递不一样，就Java中是将对象的引用拷贝以后再传递，而go语言中是直接将数组拷贝了一份再传递
}
