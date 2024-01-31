package main

import "fmt"

func main() {

	var a bool = true
	var b bool = false

	// 逻辑运算符`&&`就是`与`, 我与你，就是我和你都要满足一个结果，才执行
	if a == true && b == true {
		fmt.Println(a && b)
	}
	fmt.Println(a && b)
}
