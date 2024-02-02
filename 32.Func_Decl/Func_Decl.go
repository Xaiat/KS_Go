package main

import "fmt"

func main() {
	// 函数的调用
	printinfo()

	myprint("Xaiat.com")

	myprintnum(add(1, 2))

}

// 无参数，无返回值，的函数
func printinfo() {
	fmt.Println("printinfo")
}

// 有1个参数，的函数
func myprint(msg string) {
	fmt.Println(msg)
}

// 有1个参数，的函数
func myprintnum(num int) {
	fmt.Println(num)
}

// 有2个参数，的函数
func add(a int, b int) int {
	c := a + b
	return c
}

// 有1个返回值，的函数
// 有多个返回值，的函数
