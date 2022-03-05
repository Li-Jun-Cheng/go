package tree

import "fmt"

type Node struct {
	Value       int   //结构体就相当于Java中的类，这里定义的就是类中的属性
	Left, Right *Node //这里定义了这个结构体的两个指针对象，指向自身，分别为左子树和右子树
}

// Print go语言给结构体定义方法，这里的node就是结构体接收定义，说明了这个方法接收的是treeNode这个结构体。相当于Java中类的方法，
//只是go语言中把类属性的定义弄到了结构体中，而类中的方法独立出来定义/**
func (node Node) Print() { //在go语言中，我认为这种定义方式更好，因为go中没有class类，只有这种定义才能写成root.Print()  语义上更清晰
	//(node Node)这个接收者
	fmt.Print(node.Value)
}
func mmt(node Node) { //这种写法跟Java就一样了，但是意思和上面相同
	fmt.Println(node.Value)
}

// SetValue func (node Node) SetValue(Value int) {
//	node.Value = Value //很明显，这就是go语言版的set方法,但是这种方法是改不了属性值的，因为go语言中的值传递。所有要改为指针
//}
//修改如下
func (node *Node) SetValue(value int) { //加指针的意义就是方法里改变了传递进来属性的值，外面也能接收到
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	//这里也是我们不用区分这个node到底是指针还是实例，直接点就能引用到想要的属性了
	node.Value = value //很明显，这就是go语言版的set方法,但是这种方法是改不了属性值的，因为go语言中的值传递。所有要改为指针
}

// CreateNode 工厂函数
func CreateNode(value int) *Node {
	//这里结构体地址返回给别人使用，就会分配到堆上，外面对指针使用完以后才会被垃圾回收
	return &Node{Value: value} //go语言局部变量也可以返回给别人用，跟C++不一样，一般我们工厂函数返回结构体的地址，不用考虑这个局部变量在哪里分配
}

// Traverse 方法也可以单独写到一个go文件中
func (node *Node) Traverse() {
	if node == nil {
		return
	}
	//在go言中不用判断node.left是否是空，nil在go语言中只是一个普通的函数。也能调用方法,你只要判断了就行
	//中序遍历就是先左再中再右
	node.Left.Traverse()  //先遍历它的左子树,这里也是用到了递归
	node.Print()          //再遍历自己
	node.Right.Traverse() //再遍历它的右子树
}

//go语言有自动垃圾回收，不需要关心变量是分配在堆上还是栈上，由运行环境自动分配
