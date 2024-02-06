package main

import "fmt"

func main() {
	f("1")
	fmt.Println("2")
	// 会被延迟到最后执行
	defer f("3")
	fmt.Println("4")
	// 会被延迟到最后执行
	defer f("5")
	fmt.Println("6")
	// 会被延迟到最后执行
	defer f("7")
	fmt.Println("8")
	f("9")
	f("10")
}

func f(s string) {
	fmt.Println(s)
}
