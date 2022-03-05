package main

import "fmt"

func deferSum(n1 int, n2 int) int {

	//当执行到defer时，暂时不执行，会将defer后面的语句压入到独立的栈(defer栈)
	//当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 n1=", n1) //defer 3. ok1 n1 = 10
	defer fmt.Println("ok2 n2=", n2) //defer 2. ok2 n2= 20
	//增加一句话
	n1++                         // n1 = 11
	n2++                         // n2 = 21
	res := n1 + n2               // res = 32
	fmt.Println("ok3 res=", res) // 1. ok3 res= 32
	return res

}

func main() {
	res := deferSum(10, 20)
	fmt.Println("res=", res) // 4. res= 32
}

/*
注意事项和细节
当 go 执行到一个 defer 时，不会立即执行 defer 后的语句，而是将 defer 后的语句压入到一个栈中, 然后继续执行函数下一个语句。
当函数执行完毕后，在从 defer 栈中，依次从栈顶取出语句执行(注：遵守栈 先入后出的机制)，
在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。
说明：
在 golang 编程中的通常做法是，创建资源后，比如(打开了文件，获取了数据库的链接，或者是锁资源)， 可以执行 defer file.Close() defer connect.Close()
在 defer 后，可以继续使用创建资源.
当函数完毕后，系统会依次从 defer 栈中，取出语句，关闭资源.
这种机制，非常简洁，程序员不用再为在什么时机关闭资源而烦心。
*/
