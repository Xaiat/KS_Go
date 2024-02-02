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

	// %c 以字符方式打印 "Hello, Xaiat" 的第三个字符
	fmt.Printf("%c\n", str[2])

	// for 循环遍历字符串, 从0开始循环到字符串的长度，步进为1
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	// for 循环遍历字符串, 以字符方式打印 "Hello, Xaiat" 的第全部字符
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	// 特殊的for循环：用range配合for使用循环，遍历数组，遍历切片，遍历字符串
	// 会返回下标和对应的值，使用这个值就可以了
	for i, v := range str {
		fmt.Print(i)
		fmt.Printf("%c", v)
	}

	// string字符串是不可变的，不能修改字符串的某个字符
	str[2] = "A"
	fmt.Println(str[2])

}
