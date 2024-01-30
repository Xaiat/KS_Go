package main

import "fmt"

// name 全局变量
var age int = 50
var name string = "Xaiat"

func main() {
	// name 局部变量
	var age int = 18
	var name string = "人工智能艾特"
	// 使用局部变量，就近原则，就近原则，就近原则
	fmt.Println("name = ", name, "age = ", age)
}

/*
func test() {
	fmt.Println("name = ", name, "age = ", age)
}
*/
