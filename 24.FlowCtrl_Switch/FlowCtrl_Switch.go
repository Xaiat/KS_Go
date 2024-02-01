package main

import "fmt"

// Switch 语句
// 如果需要判断一个变量需要用到范围判断，可以使用if语句，但是如果需要判断的值是固定的，可以使用switch语句
func main() {

	var score int = 66

	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 50, 60, 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// switch 后面如果什么都不跟话，就是默认条件，是一个布尔值，bool = true
	switch {
	case false:
		fmt.Println("false")
	case true:
		fmt.Println("true")
	default:
		fmt.Println("其他")
	}
}
