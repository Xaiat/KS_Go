package main

import "fmt"

func main() {
	// 常量：程序运行期间，不可以改变的量
	// 显式类型定义
	const URL string = "http://www.xaiat.com"
	// 隐式类型定义
	const URL2 = "http://www.baidu.com"

	// 一次定义多个常量
	const a, b, c = 3.14, "Xaiat", false

	fmt.Println("URL = ", URL)
	fmt.Println("URL2 = ", URL2)
	fmt.Println("a = ", a, "b = ", b, "c = ", c)
}
