package main

import "fmt"

func main() {
	// var 变量名 数据类型
	// bool: tur or false
	// bool 默认值为 false
	var isFlag bool = true
	var isFlag2 bool = false
	fmt.Println("isFlag = ", isFlag)

	fmt.Printf("%T, %t\n", isFlag, isFlag)
	fmt.Printf("%T, %t\n", isFlag2, isFlag2)
}
