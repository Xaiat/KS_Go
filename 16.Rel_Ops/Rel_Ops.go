package main

import "fmt"

func main() {
	var a int = 11
	var b int = 10

	// 关系运算符 结果都是布尔值 bool 类型 true 或 false
	// = 赋值
	// == 等于
	// != 不等于
	// > 大于
	// < 小于
	// >= 大于等于
	// <= 小于等于
	fmt.Printf("a == b = %t\n", a == b) // 判断是否相等
	fmt.Printf("a != b = %t\n", a != b) // 判断是否相等
	fmt.Printf("a > b = %t\n", a > b)   // 判断是否相等
	fmt.Printf("a < b = %t\n", a < b)   // 判断是否相等
	fmt.Printf("a >= b = %t\n", a >= b) // 判断是否相等
	fmt.Printf("a <= b = %t\n", a <= b) // 判断是否相等

	// 判断 if 如果。。。那么。。。否则。。。
}
