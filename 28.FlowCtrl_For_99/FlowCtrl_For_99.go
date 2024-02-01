package main

import "fmt"

// 目标：使用for循环打印 99 乘法表
// 1x1=1
// 1x2=2 2x2=4
// 1x3=3 2x3=6 3x3=9
// 1x4=4 2x4=8 3x4=12 4x4=16
// 1x5=5 2x5=10 3x5=15 4x5=20 5x5=25
// 1x6=6 2x6=12 3x6=18 4x6=24 5x6=30 6x6=36
// 1x7=7 2x7=14 3x7=21 4x7=28 5x7=35 6x7=42 7x7=49
// 1x8=8 2x8=16 3x8=24 4x8=32 5x8=40 6x8=48 7x8=56 8x8=64
// 1x9=9 2x9=18 3x9=27 4x9=36 5x9=45 6x9=54 7x9=63 8x9=72 9x9=81

// 方法：使用for循环嵌套打印
// 1.先打印 1 行
// 2.再变成 9 行
// 3.再让 9 行变化起来

// 1. 首先打印最底下的一行，9 是固定的，所以可以使用for循环	==== 第 1 步 ====
// for i := 1; i <= 9; i++ {
// 	fmt.Printf("%d*9=%d \t", i, i*9)
// }

// 2. 再用for循环嵌套打印上面的行,再加上换行 fmt.Println() ==== 第 2 步 ====
// for j := 1; j <= 9; j++ {
// 	for i := 1; i <= 9; i++ {
// 		fmt.Printf("%d*9=%d \t", i, i*9)
// 	}
// 	fmt.Println()
// }

// 3. 再把`fmt.Printf("%d*9=%d \t", i, i*9)`当中的第一个`9`改成`%d`并在`i,`后面添加一个变量`j`。第二个`9`改成`j` ==== 第 3 步 ====
// fmt.Printf("%d*%d=%d \t", i, j, i*j)

// 4. 再把上一行的`for i := 1; i <= 9; i++ {`当中的`9`改成变量`j` ==== 第 4 步 ====

func main() {
	/*
		for j := 1; j <= 9; j++ {
			for i := 1; i <= j; i++ {
				fmt.Printf("%d*%d=%d \t", i, j, i*j)
			}
			fmt.Println()
		}

	*/

	/* ==== 第 1 步 ====
	for i := 1; i <= 9; i++ {
		fmt.Printf("%d*9=%d \t", i, i*9)
	}
	*/
	/* ==== 第 2 步 ====
	for j := 1; j <= 9; j++ {
		for i := 1; i <= 9; i++ {
			fmt.Printf("%d*9=%d \t", i, i*9)
		}
		fmt.Println()
	}
	*/
	/* ==== 第 3 步 ====
	for j := 1; j <= 9; j++ {
		for i := 1; i <= 9; i++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println()
	}
	*/
	/* ==== 第 4 步 ====
	for j := 1; j <= 9; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d*%d=%d \t", i, j, i*j)
		}
		fmt.Println()
	}
	*/

	for j := 1; j <= 19; j++ {
		for i := 1; i <= j; i++ {
			fmt.Printf("%d*%d=%d ", i, j, i*j)
		}
		fmt.Println()
	}
}
