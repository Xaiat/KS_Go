package main

import "fmt"

/*
打印一个方阵
* * * * * 换行
* * * * *
* * * * *
* * * * *
* * * * *
思路：
1.先用for循环打印 1 行 5 个*
2.再在上面的for循环外面再套一个for循环，打印 5 行
3.最后在第二个for循环外面加一个换行
*/
func main() {

	for j := 1; j <= 5; j++ {
		for i := 1; i <= 5; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}

}
