package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	ss = "kkk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}
func main() {
	k, l, m, s1, s2 := true, false, 3, "hello", "world"
	fmt.Println(k, l, m, s1, s2)
	fmt.Print("让编译器自动决定类型")
	var stockcode = 123
	var name = "he"
	var enddate = "2020-12-31"
	var url = "Code=%d&endDate=%s"
	var target_url = fmt.Sprintf(url, stockcode, enddate)
	fmt.Println(target_url)
	variableZeroValue()
	/*
		这是我的第一个简单程序
	*/
	fmt.Printf("hello, world\n")
	fmt.Printf("Google" + "Runoob" + name)
	//内建变量  bool string
	fmt.Println()
	euler()
	fmt.Println("======================================")
	triangle()
}
func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Printf("%.3f", cmplx.Exp(1i*math.Pi)+1)
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}
func triangle() {
	var a, b int = 3, 4
	var c int
	//勾股定理
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
	consts()
}
func consts() {
	const (
		filename = "abc.txt"
		a, b     = 3, 4
	)
	//常量就是一个文本，这里既可以是int，也可以是float,这里的常量不用全部大写，const数值可以作为各种类型使用，就不需要强制类型转换
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
	fmt.Println("========================================================")
	eunms()
}
func eunms() {
	//iota就是一个自增值
	const (
		cpp = iota
		java
		python
		golang
	)
	//b,kb,mb,gb,tb,pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
