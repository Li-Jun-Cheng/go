package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//该算法将整数转换成二进制的表达式
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

//go语言中函数名在前，返回值的类型在后，参数里定义的变量也一样，参数名在前，类型在后
func printFile(filename string) {
	file, err := os.Open(filename)
	//err错误信息不等于空就说明有错误，有错误就打印错误信息并暂停程序
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
func forever() {
	i := 0
	//什么都省略了就是死循环
	//for,if后面的条件没有括号，if条件里也可以定义变量,没有while，switch不需要break,也可以直接switch多个条件
	for {
		fmt.Print(convertToBin(i))
		i++
	}
}
func main() {
	fmt.Println(
		convertToBin(5),  //101
		convertToBin(13), //
		convertToBin(5546547976532132),
		convertToBin(0))
	printFile("abc.txt")
	//forever()
}
