package main

import "fmt"

func main() {
	/* 	非go语言的交换变量值
	   	// a = 100
	   	// b = 200
	   	// temp = 0
	   	// 1. temp (a) = a
	   	// 2. a = b
	   	// 3. b = temp(a)
	*/
	// go语言的交换变量值
	var a int = 100
	var b int = 200

	b, a = a, b
	fmt.Println("a = ", a, "b = ", b)

}
