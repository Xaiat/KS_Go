package main

import "fmt"

/*
根据Go语言的数据类型的特点，可以将一个函数作为另一个函数的参数。
fun2(), fun2()
将fun2函数作为fun2这个函数的参数
fun3函数：就叫做高阶函数，接受了一个函数作为参数
fun2函数：就叫做回调函数，作为另一个函数（fun2）的参数
*/

func main() {
	r1 := add(1, 2)
	fmt.Println("r1 =", r1)

	r2 := oper(3, 4, add)
	fmt.Println("r2 =", r2)
}

// 高阶函数,可以接受一个函数作为参数
func oper(a, b int, fun func(int, int) int) int {
	r := fun(a, b)
	return r
}

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}
