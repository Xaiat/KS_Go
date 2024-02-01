package main

import "fmt"

func main() {
	var x int
	var y float64

	// 定义了两个变量
	// 需求: 从键盘输入一个整数,并将这个整数赋值给变量x，然后再从键盘输入一个浮点数,并将这个浮点数赋值给变量y
	// fmt.Println()	// 打印并换行
	// fmt.Printf()		// 格式化输出
	// fmt.Print()		// 打印不换行输出

	fmt.Println("请输入两个数 1、整数， 2、浮点数，用空格隔开，例如: 100 3.14")
	// 变量取地址：&变量名
	// 指针、地址来修改和操作变量，指针式一个指向内存地址的变量
	fmt.Scanln(&x, &y) // 从键盘输入

	// fmt.Sacanln()	// 接收键盘输入 Scan line
	// fmt.Sacanf()		// 从键盘输入
	// fmt.Sacan()		// 从键盘输入

}
