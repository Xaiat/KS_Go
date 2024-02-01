package main

import "fmt"

// if 如果...否则... else
func main() {

	var a int = 15

	// if 如果 条件被满足，那么执行 if { 语句块中的代码 }
	// if 如果 条件不被满足，那么执行 else { 语句块中的代码 }
	if a > 20 {
		fmt.Println("a > 20")
	}
	if a > 10 {
		fmt.Println("a > 10")
	}
	/*
		// 分数
		var score int = 80

		// A 90-100 B 80-89 C 70-79 D 60-69
		if score >= 90 && score <= 100 {
			fmt.Println("A")
		}
		if score >= 80 && score < 90 {
			fmt.Println("B")
		}
		if score >= 70 && score < 80 {
			fmt.Println("C")
		}
		if score >= 60 && score < 70 {
			fmt.Println("D")
		}
	*/
	// 分数
	var score int = 92
	// A 90-100 B 80-89 C 70-79 D 60-69
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score < 90 {
		fmt.Println("B")
	} else if score >= 70 && score < 80 {
		fmt.Println("C")
	} else if score >= 60 && score < 70 {
		fmt.Println("D")
	} else {
		fmt.Println("不及格")
	}

}
