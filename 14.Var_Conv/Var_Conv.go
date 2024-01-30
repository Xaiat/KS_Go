package main

import "fmt"

// 变量数据类型转换
// 转换后的变量 a := 要转换的变量（要转换的类型）
// 备注：转换的时候，数据类型必须是兼容的（同一种）（比如：整型是不能转换为布尔类型的）
func main() {

	a := 3   // int
	b := 5.0 // float64

	// 需求：将int类型的变量a，转换为float64类型 类型转换
	c := float64(a)
	d := int(b)

	// bool 整形数据是不能转换为布尔类型的
	// e := bool(a)

	fmt.Printf("%T, %d\n", a, a)
	fmt.Printf("%T, %f\n", b, b)
	fmt.Printf("%T, %f\n", c, c)
	fmt.Printf("%T, %d\n", d, d)
	// fmt.Printf("%T, %t\n", e, e)

}
