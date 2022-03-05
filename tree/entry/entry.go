package main

import "fmt"

/**
总结：1.要改变内容必须使用指针接收者
     2.结构过大也考虑使用指针接收者
     一致性：如有指针接收者，最好都是指针接收者
*/
func main() {
	var root tree.Node
	root = Node{Value: 3}
	root.Left = &Node{}
	root.Right = &Node{5, nil, nil} //这里相当于Java中的构造方法，初始化时给对象中的属性赋值
	root.Right.Left = new(Node)     //new()返回的是Node的地址,这句话的意思是root这个结构体中右子树中左子树的值
	// 是一个空结构体new(Node),这里和Java中的new对象很像啊,这里new了一个新对象（空结构体）
	//指针也能够点,在go语言中，不管是指针还是实例，我们都是一路点下去
	nodes := []Node{
		{Value: 3},
		{},
		{6, nil, &root},
	} //现在的操作就相当于Java中针对list集合的操作，slice(切片)跟Java中的list集合很像，底层数据结构都是数组，都是对数组的封装
	fmt.Println(nodes)
	//不论是地址还是结构本身，一律使用.来访问成员，可以理解为所有访问的对象全部都是Java中的静态变量
	root.print()
	fmt.Println("========================")
	println(root)
	fmt.Println("|||||||||||||||||||||||||||||||||||||||||||||||")
	root.Right.Left.setValue(4) //这里root.Right.Left也不需要我们手动转为指针了，这里很人性化,指针调用把地址直接传给接受者，
	// 而实例调用还有点麻烦，系统会把地址解析出来，然后访问指针所指向的值，再拷贝一份传递给接收者
	root.Right.Left.print()

	//这个例子说明了不管我定义值接收者还是指针接收者，我都可以拿着一个结构体具体的值去调用它们
	root.print()       //值接收者会拷贝一份root这个结构体的值给我们调用的接收者
	root.setValue(100) //这里就不会拷贝，它会把root的地址拿出来直接给接收者
	fmt.Println("<><><><><><><><><><><><><><><><><><><<>")
	pRoot := &root
	pRoot.print()
	pRoot.setValue(200)
	pRoot.print()
	var mRoot *tree.Node
	mRoot.setValue(800)
	mRoot = &root
	mRoot.setValue(300)
	mRoot.print()
	fmt.Println()

	hroot := myThreeNode{&root}
	hroot.postOrder()
	fmt.Println()
}

type myThreeNode struct { //这里先定义了一个结构体
	node *tree.Node //结构体中的属性是一个tree包下的Node结构体指针
}

func (myNode *myThreeNode) postOrder() { //这里定义了一个myThreeNode的指针接收者
	if myNode == nil || myNode.node == nil {
		return
	}
	//literal:字面常量或者称为字面量
	//myThreeNode{myNode.node.Left}就是一个 literal ,literal不可以直接取到地址(cannot take the address of myThreeNode literal)
	//先弄一个变量去得到字面量的地址才能  .postOrder()
	left := myThreeNode{myNode.node.Left}
	right := myThreeNode{myNode.node.Right}
	left.postOrder() //给myThreeNode中的node指针赋值，该指针指向了tree包下的Node结构体，
	// 然后myThreeNode这个结构体调用了自己的指针接收者postOrder,这里又是递归，自身调用自身，实现了后序遍历，先是左子树然后右子树最后到自己
	//中序和后序，这里的中是指自己是在中间被遍历到还是最后被遍历到
	right.postOrder()
	myNode.node.Print()
}
