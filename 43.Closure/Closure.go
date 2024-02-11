package main

import "fmt"

/*
一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量;
并且该外层函数的返回值就是这个内层函数；
这个内层函数和外层函数的局部变量，统称为闭包结构。
局部变量的生命周期就会发生改变，正常的局部变量会随着函数的调用而创建，随着函数的结束而销毁；
但是比变结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还在继续使用。
JavdScript中的闭包结构，和Go语言中的闭包结构是一样的，因为Go语言和JavaScript都是函数式编程语言。
本质原因是都支持回调函数，所以都有闭包结构。
*/

func main() {

	r1 := increment()
	// 打印r1的地址，并没有执行r1的自增函数代码
	fmt.Println(r1)
	// 执行r1的自增函数代码
	v1 := r1()
	fmt.Println(v1)
	// 再次把r1赋值给v2
	v2 := r1()
	fmt.Println(v2)
	fmt.Println(r1())
	fmt.Println(r1())
	fmt.Println(r1())

	// increment()函数中的r1已经销毁了，所以r2和r1是两个不同的闭包
	r2 := increment()
	v3 := r2()
	fmt.Println(v3)
	fmt.Println(r1())
	fmt.Println(r2())
}

// 自增函数
func increment() func() int {
	// 定义了一个局部变量i
	i := 0
	// 定义了一个匿名函数，给变量自增并返回
	// 内层函数，只是定义了，还没有执行
	fun := func() int {
		i++
		return i
	}
	return fun
}
