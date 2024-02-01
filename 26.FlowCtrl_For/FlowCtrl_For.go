package main

import "fmt"

func main() {
	// for 后面要跟三个条件：初始化语句、条件表达式、后置语句
	// 循环10次，for 条件的起始值；循环条件；控制变量自增或自减
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// 计算1到10的数字之和
	// sum := 0
	// for i := 1; i <= 1000; i++ {
	// 	sum = sum + i
	// 	fmt.Println(sum)
	// }

	//
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

}
