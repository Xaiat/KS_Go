package main

import "fmt"

func main() {
	var a int = 10
	var b int = 3
	// var c int

	// + - * / % ++ --
	fmt.Println("a = ", a, "\nb = ", b) // a = 10 b = 3
	fmt.Println("a + b = ", a+b)        // 10 + 3 = 13
	fmt.Println("a - b = ", a-b)        // 10 - 3 = 7
	fmt.Println("a * b = ", a*b)        // 10 * 3 = 30
	fmt.Println("a % b = ", a%b)        // 10 % 3 = 1
	a++                                 // a = a + 1
	fmt.Println("a ++ = ", a)           // 11
	a = 100                             // a = 100
	a--                                 // a = a - 1
	fmt.Println("a -- = ", a)           // 99
}
