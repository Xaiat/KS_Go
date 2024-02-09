package main

import "fmt"

// func() 函数本身也是一个数据类型，可以赋值给变量，也可以作为函数的参数和返回值。
func main() {

	// f1 如果不加()括号，函数就是一个变量
	// f1() 如果加了()括号则是调用函数，不加括号则是函数类型。
	fmt.Printf("%T\n", f1) // func() | func(int, int) | func(int, int) int
	// fmt.Printf("%T", 10)       // int
	// fmt.Printf("%T\n", "Xiat") // string

	// 定义函数类型的变量
	var f5 func(int, int)
	// 函数本身是引用类型的传递，所以可以赋值给变量
	f5 = f1
	fmt.Println(f5) // 0x10b0b60
	fmt.Println(f1) // 0x10b0b60

	f5(1, 2)
}

/*
func f1(a, b int) int {
	return 0

}
*/

func f1(a, b int) {
	fmt.Println(a, b)
}
