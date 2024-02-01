package main

import "fmt"

func main() {
	// for 循环的第一种写法：后面要跟3个条件参数：初始化语句、条件表达式、后置语句 ====1====
	// for 条件的起始值；循环条件；控制变量自增或自减
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// 计算1到10的数字之和
	// sum := 0
	// for i := 1; i <= 1000; i++ {
	// 	sum = sum + i
	// 	fmt.Println(sum)
	// }

	// for 循环的第二种写法：只有第2个条件参数：条件表达式；相当于 while 循环 ====2====
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// for 循环的第三种写法：没有条件参数，{} 无限循环 ====3====
	i := 1
	for {
		fmt.Println(i)
		i++
	}

}
