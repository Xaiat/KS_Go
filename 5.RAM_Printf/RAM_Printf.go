package main

import "fmt"

func main() {
	var num int
	num = 100
	fmt.Printf("num:%d, 内存地址：%p", num, &num)
}
