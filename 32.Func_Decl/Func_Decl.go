package main

import "fmt"

func main() {
	// 调用无参数，无返回值，的函数
	printinfo()

	// 调用有参数的函数，我们就需要按接收函数的参数类型来传递
	myprint("Xaiat.com")

	// 调用有返回值的函数，我们就需要按接收函数的返回值类型来接收
	c := add(1, 2)
	myprintnum(c)

	// 调用有多个返回值的函数
	x, y := swap("Xaiat", ".com")
	fmt.Println(x, y)

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
// 有1个返回值，的函数
func add(a int, b int) int {
	c := a + b
	return c
}

// 有多个返回值，的函数
func swap(x string, y string) (string, string) {
	return y, x
}
