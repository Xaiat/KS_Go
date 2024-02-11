package main

import "fmt"

// Calls Function f1 调用函数 f1
func main() {
	f1()
	// Functon f1 is also a variable f1 函数本身也是一个变量
	// Function f1 is of type func() f1 的类型是 func()
	f2 := f1
	f2()

	// Anonymous Function 匿名函数
	f3 := func() {
		fmt.Println("我是f3匿名函数")
	}
	f3()

	// Anonymous Function calling itself,
	// 匿名函数自己调用自己,可以有参数
	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是f4匿名函数")
	}(1, 2)

	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是f5匿名函数")
	}(1, 2)

	func(a, b int) {
		fmt.Println(a, b)
		fmt.Println("我是f5匿名函数")
	}(1, 2)
}

// Definitions Function f1 定义一个函数 f1
func f1() {
	fmt.Println("我是f1函数")
}

/*
Go语言是支持函数式编程：
1.将匿名函数作为另一个函数的参数，回调函数
2.将匿名函数作为另一个函数的返回值，可以形成闭包结构
*/
