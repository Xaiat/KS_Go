package main

import "fmt"

func main() {

	// ASCII编码 72 101 108 108 111 44 32 88 97 105 97 116
	str := "Hello, Xaiat"
	fmt.Println(str)

	// 获取字符串的长度
	fmt.Println("字符串的长度为：", len(str))

	// 获取制定的字节
	fmt.Println("获取指定的字节：", str[0])

}
