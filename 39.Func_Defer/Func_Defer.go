package main

import "fmt"

/*
func main() {
	f("1")
	fmt.Println("2")
	// 会被延迟到最后执行
	defer f("3")
	fmt.Println("4")
	// 会被延迟到最后执行
	defer f("5")
	fmt.Println("6")
	// 会被延迟到最后执行
	defer f("7")
	fmt.Println("8")
	f("9")
	f("10")
}

func f(s string) {
	fmt.Println(s)
}
*/

// defer 语句可以用来做关闭操作：关闭文件，解锁资源，打印最终报告，关闭数据库连接等
func main() {
	a := 10
	fmt.Println(a)
	// 参数在这里就已经传递进去了，但是不会执行，会被延迟到最后执行
	defer f(a)
	a++
	fmt.Println("end a = ", a)
}

func f(s int) {
	fmt.Println("函数里面的 a = ", s)
}
