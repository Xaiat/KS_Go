package main

import "fmt"

func main() {
	var str string
	str = "Hello, Xaiat"
	fmt.Printf("%T, %s\n", str, str)

	// 单引号 字符，整型-ASCII字符码
	v1 := '中'
	v2 := "A"

	// 编码表 ASCII字符码
	// 扩展知识：
	// 所有汉字编码 GBK 编码表：兼容ASCII，中文2个字节
	// 所有文字编码 Unicode编码表：兼容ASCII，中文3个字节

	fmt.Printf("%T, %d\n", v1, v1) // int32, 65
	fmt.Printf("%T, %s\n", v2, v2)

	// 字符串连接在一起，称为字符串拼接,使用+号
	fmt.Println("Hello" + ",Xaiat")

	// 转义字符用 \ 反斜杠开头
	fmt.Println("Hello\"Xaiat")
	// \n 是换行符号
	fmt.Println("Hello\nXaiat")
	// \t 是制表符号
	fmt.Println("Hello\tXaiat")
}
