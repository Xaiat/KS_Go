package main

import "fmt"

func main() {
	// 定义一个整型
	// byte 好比 uint8
	// rune 好比 int32
	// int  好比 int32 或 int64
	var age int = 18
	fmt.Printf("%T, %d\n", age, age)

	// 定义一个浮点型
	// 默认是6位小数打印 3.140000
	var money float64 = 3.14
	fmt.Printf("%T, %.2f\n", money, money)

	// Go语言中的浮点型默认是float64类型
	// float64 尽量使用 float64 来定义浮点类型的小数
	// float32 精度只有小数点后6位，可能会因精度不够，而有数据损失
	// float64 精度只有小数点后15位，一般情况下，精度够用
	var num1 float32 = -123.0000901
	var num2 float64 = -123.0000901
	fmt.Println("num1 = ", num1, "num2 = ", num2)

	// 学习初期主要记住3种：int、float64、byte

}
