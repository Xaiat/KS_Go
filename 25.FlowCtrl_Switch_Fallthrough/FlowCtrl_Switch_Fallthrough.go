package main

import "fmt"

func main() {

	a := false
	switch a {
	case false:
		fmt.Println("1. case 条件为 false")
		fallthrough // fallthrough 用来 case 穿透，不管下一个条件是否满足，都强制执行后面的 case。实际开发中很少使用。
	case true:
		if a == false {
			break // break 终止，不管后面的条件是否满足，都不执行
		}
		fmt.Println("2. case 条件为 true")
	}

}
