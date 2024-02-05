package main

import "fmt"

// 全局变量
// 定义变量的时候，要思考这个变量是在哪里使用，如果是全局变量，就定义在函数外部；如果是局部变量，就定义在函数内部
var num int = 100

func main() {
	/*
		// 函数体内的局部变量
		temp := 100

		// if语句、for语句定义的一次性变量都是局部变量，只能在语句内使用
		if b := 1; b <= 10 {
			// 语句内的局部变量
			temp := 50
			fmt.Println(temp) // 50，局部变量，遵循就近原则
			fmt.Println(b)    // 1
		}
		fmt.Println(temp) // 100
	*/

	num := 20
	// 无论是全局变量还是局部变量，都是找最近的变量去打印
	fmt.Println(num) // 20

	f1()
	f2()
	f3()
}

func f1() {
	num := 30
	// 无论是全局变量还是局部变量，都是找最近的变量去打印
	fmt.Println(num) // 30
}

func f2() {
	num := 40
	// fmt.Println(a) // 不能使用其他函数中定义的变量
	// 无论是全局变量还是局部变量，都是找最近的变量去打印
	fmt.Println(num) // 40
}

func f3() {
	// 无论是全局变量还是局部变量，都是找最近的变量去打印
	fmt.Println(num) // 100
}
