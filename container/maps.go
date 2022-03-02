package main

import "fmt"

func main() {
	//map的定于  map[k]v  复合map的定义是 map[k]map[k]v
	m := map[string]string{
		"name":    "ccmouse",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}
	fmt.Println(m)

	m2 := make(map[string]int)
	fmt.Println(m2)
	//几种方式都可以定义map
	var m3 map[string]int //m3==nil  go语言中的nil是可以参与运算的
	fmt.Println(m, m2, m3)

	fmt.Println("Traversing map") //map的遍历

	for k, v := range m {
		fmt.Println(k, v)
	} //k或者v都是可以省略的
	for k := range m {
		fmt.Println("只写一个参数接收的时候，接收的是key：" + k)
	}
	for _, v := range m {
		fmt.Println(v)
	}
	//这个map是一个hashMap，key是无序的，所以每次遍历得到的结果都不太一样
	fmt.Println("Getting values")
	courseName, ok := m["course"] //当key不存在时，得到得值是一个空串，go语言不怕变量不初始化
	fmt.Println(courseName, ok)   //输出的是map里的键“course”对应的值和是否存在的boolean值
	//那么怎么判断key是否存在呢
	causeName, ok := m["cause"]
	fmt.Println(causeName, ok)
	if coursName, ok := m["cau"]; ok {
		fmt.Println(coursName)
	} else {
		fmt.Println("key does not exist")
	}
	//删除元素
	fmt.Println("Deleting values")
	delete(m, "name")
	fmt.Println(m)
	site, ok := m["site"] //site和ok如果都定义过了，就不能用：=了，直接用=符号就可以
	fmt.Println(site, ok)
	delete(m, "site")
	fmt.Println(m)
	fmt.Println(len(m)) //获取map集合中元素的个数
	//总结：
	//创建：make(map[string] int)
	//获取元素：m[key]
	//key不存在时，获得value类型的初始值
	//用value,ok:=m[key]来判断是否存在key
	//用delete删除一个key

	//Map的遍历
	//1.使用range遍历key,或者遍历key,value对

	//2.不保证排序，如需要排序，需要手动对key排序

	//3.使用len获的元素的个数

	//map中什么样的类型能作为key呢,java中要实现hashcode和equals方法才能作为key
	//1.而map使用哈希表（又称为散列表），必须可以比较相等
	//2.除了slice(切片) ,map,function的内建类型都可以作为key
	//3.Struct类型（结构体）不包含上述字段，也可作为key(系统自动帮我做了很多)

}
