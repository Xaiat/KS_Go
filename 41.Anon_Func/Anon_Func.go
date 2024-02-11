package main

import "fmt"

// Calls Function f1
// 调用函数 f1
func main() {
	f1()
	// Functon
	// 函数本身也是一个变量
	f2 := f1
	f2()
}

// Definitions Function f1
// 定义一个函数 f1
func f1() {
	fmt.Println("我是f1函数")
}
