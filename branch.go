package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		//err如果不为空就说明出错了
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//contents,err因为是在if里定义的，所以它的作用域只在if里，出了if就读取不到了
	//知识点：if的条件里可以赋值
	//if的条件里赋值的变量作用域就在这个if语句里
	fmt.Println(
		grade(0),
		grade(59),
		grade(60),
		grade(82),
		grade(99),
		grade(100),
		grade(101))
}
func grade(score int) string {
	g := ""
	//这里不需要break，go语言会自动添加break。你要不想使用break，除非使用fallthrough
	//switch后面可以没有表达式
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf(
			"Wrong score:%d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score <= 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

//for的条件里不需要括号
//for的条件里可以省略初始条件，结束条件，递增表达式
