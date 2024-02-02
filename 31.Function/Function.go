package main

import "fmt"

/*
	函数是基本的代码块，用于执行一个任务。
	Go语言最少有个 main() 函数。
	可以通过函数来划分不同功能，逻辑上每个函数执行的是制定的任务。
	函数声明告诉了编译器函3件事：1.函数的名称；2.参数；3.返回类型。
*/

func main() {
	fmt.Println("Hello, Xaiat")

	// 调用函数 函数名(参数1, 参数2, ...)
	fmt.Println(add(1, 2))

}

// func 函数名(参数1 参数类型, 参数2 参数类型, ...) 函数调用后的返回值类型 {
//		函数体：执行一段代码
//		return 返回结果
//}

func add(a int, b int) int {
	c := a + b
	return c
}
