package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func main() {
	fmt.Println(eval(3, 4, "x"))
	if result, err := eval(3, 4, "X"); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	q, r := div(13, 3)
	//Go语言中可以有多个返回值，这点是很不一样的！！！！
	//多个返回值的时候，将有一个返回值是不需要的时候，用下划线代替就代表忽略该位置的返回值
	//多返回值一般都是返回一个值加一个err，很少真的返回两个值，如果出错的话就把错误信息返回回去
	//panic就是中断程序的意思
	a, _ := div(17, 6)
	fmt.Println(a)
	fmt.Println(q, r)
	//pow要单独定义嫌麻烦的话，可以像java一样用匿名内部类的写法
	fmt.Println(apply(pow, 3, 4))
	fmt.Println("模仿Java匿名内部类的写法展示结果：=====================》》》》》》》")
	fmt.Println(apply(func(a int, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 3, 4))
	//匿名内部函数的打印信息里main.main.func1 表示main包下的main方法中的第一个匿名内部函数所以加func1
	fmt.Println(sum(1, 2, 3, 4, 5)) //我们传一个可变长参数列表，用起来呢就像是一个数组一样用
	//总结：1.返回值类型写在最后面  2.可返回值多个值  3.函数作为参数  4.没有默认参数，可选参数
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println("Plus版本的交换，用逻辑交换，不去使用指针交换")
	y, h := 5, 8
	y, h = swapPlus(y, h)
	//这里相当于给y和h两个变量里存储的值进行了交换，上面用指针的方案是对变量指向的内存地址做交换。值交换和内存地址交换一般来说，值交换更好d1sdl
	fmt.Println(y, h)
}

//&是取地址符号，*是指向地址的指针定义符
//Go语言中只有值传递，没有引用传递。如果想用引用传递就通过指针，因为指针指向了同一块内存地址。
func swap(a, b *int) { //这里定义了要传进来两个int类型数据的内存地址（指针）
	*b, *a = *a, *b
}

func swapPlus(a, b int) (int, int) {
	return b, a
}

//Go语言没有什么默认参数，函数重载这些东西，只有一个可变长度参数
//如果一个函数的形参列表中有可变参数，则可变参数需要放在形参列表最后
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func pow(a, b int) int {
	//int(需要强制类型转换的数)，这就是Go语言中的强制类型转换了
	return int(math.Pow(float64(a), float64(b)))
}

//函数返回多个值时可以起名字
//仅用于非常加简单的函数，我认为实际放回的时候，能封装就封装，返回多个参数反而不好处理，除非遇到特殊的场景非这么做不可
//起不起名字对调用者而言没有区别
func div(a, b int) (q, r int) {
	/*return a/b,a%b*/
	//第二种写法,不建议
	q = a / b
	r = a % b
	return
}

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operation:%s", op)
	}
}

//使用函数式编程来重写eval方法,op这个函数呢是数据两个int类型的参数，
//输出一个参数int,这两个参数来源与后面的a,b,参数的类型是int,最后定义了apply的返回值是int
//go语言名在前，类型在后的这种写法在定义复合函数的时候变的更容易
//语义分析：有一个函数apply传进去一个函数，参数是a,b.得到这个传进去参数的返回值，将返回值int作为apply的参数进入到apply这个函数里，最后返回值是Int
//在 Go 中，函数也是一种数据类型，可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用
//函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用
func apply(op func(int, int) int, a, b int) int {
	//reflect这个包，它呢可以进行反射，去获得这个op(函数变量)里真正的值，Pointer获得这个函数真正的指针，指向了传进来的这个函数，最后通过指针得到这个函数的名字
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	//Printf即格式化输出，opName,a,b依次会填充到%s %d %d这个三个位置，填充的数据类型%s代表字符串，%d代表整数
	//Printf有两个参数，前面的是格式定义，后面的是一个可变长参数代表填充在前面定义的格式里的数据
	fmt.Printf("Calling function %s with args"+
		"(%d,%d)\n", opName, a, b)
	return op(a, b)
}
