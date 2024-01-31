package main

import "fmt"

func main() {
	var a int = 21
	var b int

	// '='是赋值运算符,将右边的值赋给左边的量
	b = a // b = 21
	fmt.Println(b)

	// `+=`是相加后再赋值运算符 b = b + a
	// 写简化代码的时候,可以使用+=,这样可以减少代码量
	b += a // b = 42
	fmt.Println(b)

	// `-=`是相减后再赋值运算符 b = b - a
	// 写简化代码的时候,可以使用-=,这样可以减少代码量
	b -= a // b = 21
	fmt.Println(b)

	// `*=`是相乘后再赋值运算符 b = b * a
	// 写简化代码的时候,可以使用*=,这样可以减少代码量
	b *= a // b = 441
	fmt.Println(b)

	// `/=`是相除后再赋值运算符 b = b / a
	// 写简化代码的时候,可以使用/=,这样可以减少代码量
	b /= a // b = 21
	fmt.Println(b)

	// `%=`是求余后再赋值运算符 b = b % a
	// 写简化代码的时候,可以使用%=,这样可以减少代码量
	b %= a // b = 0
	fmt.Println(b)

	// `<<=`是左移后再赋值运算符 b = b << a
	// 写简化代码的时候,可以使用<<=,这样可以减少代码量
	b <<= 2 // b = 0
	fmt.Println(b)

	// `>>=`是右移后再赋值运算符 b = b >> a
	// 写简化代码的时候,可以使用>>=,这样可以减少代码量
	b >>= 2 // b = 0
	fmt.Println(b)

	// `&=`是按位与后再赋值运算符 b = b & a
	// 写简化代码的时候,可以使用&=,这样可以减少代码量
	b &= 2 // b = 0
	fmt.Println(b)

	// `|=`是按位或后再赋值运算符 b = b | a
	// 写简化代码的时候,可以使用|=,这样可以减少代码量
	b |= 2 // b = 2
	fmt.Println(b)

	// `^=`是按位异或后再赋值运算符 b = b ^ a
	// 写简化代码的时候,可以使用^=,这样可以减少代码量
	b ^= 2 // b = 0
	fmt.Println(b)

	// `&^=`是按位清零后再赋值运算符 b = b &^ a
	// 写简化代码的时候,可以使用&^=,这样可以减少代码量
	b &^= 2 // b = 0
	fmt.Println(b)

	// `:=`是声明并赋值运算符
	// 写简化代码的时候,可以使用:=,这样可以减少代码量
	c := 21
	fmt.Println(c)

	// `++`是自增运算符
	// 写简化代码的时候,可以使用++,这样可以减少代码量
	c++
	fmt.Println(c)

	// `--`是自减运算符
	// 写简化代码的时候,可以使用--,这样可以减少代码量
	c--
	fmt.Println(c)

}
