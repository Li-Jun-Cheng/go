package main

import (
	"fmt"
	"unicode/utf8"
)

//寻找最长不重复子串
func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[byte]int) //创建一个Map集合
	start := 0
	maxLength := 0

	for index, ch := range []byte(s) { //这里转为byte只支持英文字母了unicode编码的原因
		//Map的特性就是键不会重复，这里的Map键是字符串中拆解的字符，值是每个字符的下标
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if lastOccurred[ch] >= start {
			start = lastOccurred[ch] + 1 //更新其实下
		}
		if index-start+1 > maxLength {
			maxLength = index - start + 1 //更新最长不重复子串的结果
		}
		lastOccurred[ch] = index //把字符的下标更新到Map中
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	/**
	每个中文是三个字节，go语言原生就用UTF-8编码；UTF-8编码使用的是可变长的存储方案；英文字母就是一个字节，中文就是三个字节
	rune就是一个int32类型的变量(四个字节的类型)使用,utf8这个包有很多方法可以处理中文字符
	rune相当于go的char
	使用rang遍历pos,rune对
	使用utf8.RuneCountInString获得字符数量
	使用len获得字节长度
	使用[]byte获得字节
	strings包里提供了很多操作字符串的函数
	*/
	count := utf8.RuneCountInString("我是中国人") //就可以正确统计中文字符的个数了
	fmt.Println(count)
	s := "HelloWorld你好世界"
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch) //我们就能获得第几个字符是谁
	}
}
