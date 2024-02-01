package main

import "fmt"

// if 嵌套语句
func main() {
	/*
		var a int = 100
		var b int = 201

		// if 嵌套语句
		if a == 100 {
			fmt.Println("a满足条件")
			if b == 200 {
				fmt.Println("b满足条件")
			}
		}
	*/

	// 验证密码，再次输入密码，理解为输入两次密码, a是第一次输入，b是第二次输入
	var a, b int
	var pwd int = 20240201
	// 用户输入密码
	fmt.Println("请输入密码:")
	fmt.Scan(&a)

	// 业务逻辑：验证密码是否正确
	if a == pwd {
		fmt.Println("请再次输入密码:")
		fmt.Scan(&b)
		if b == pwd {
			fmt.Println("登录成功了")
		} else {
			fmt.Println("登录失败，第二次密码错误")
		}
	} else {
		fmt.Println("登录失败，密码错误")
	}

}
