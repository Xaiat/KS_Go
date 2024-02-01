package main

import "fmt"

func main() {

	// for循环中常用的两个退出循环方法：break和continue
	// break：彻底的退出循环，不再执行循环体中的代码
	for i := 0; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("=============================")

	// for循环中常用的两个退出循环方法：break和continue
	// continue：退出本次循环，继续执行下一次循环
	for i := 0; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}

}
