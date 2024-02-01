/* ====1、从键盘输入一个整数，然后打印出来====
package main

import "fmt"

func main() {
	var x int
	var y float64

	// 定义了两个变量
	// 需求: 从键盘输入一个整数,并将这个整数赋值给变量x，然后再从键盘输入一个浮点数,并将这个浮点数赋值给变量y
	// fmt.Println()	// 打印并换行
	// fmt.Printf()		// 格式化输出
	// fmt.Print()		// 打印不换行输出

	fmt.Println("请输入两个数 1、整数， 2、浮点数，用空格隔开，例如: 100 3.14")
	// 变量取地址：&变量名
	// 指针、地址来修改和操作变量，指针式一个指向内存地址的变量
	// Scanln()	的作用：阻塞等待用户的键盘输入，按回车后才会继续向下执行
	fmt.Scanln(&x, &y) // 从键盘输入
	fmt.Println("x=", x)
	fmt.Println("y=", y)

	// fmt.Sacanln()	// 接收键盘输入 Scan line
	// fmt.Sacanf()		// 从键盘输入
	// fmt.Sacan()		// 从键盘输入

}
*/
/* ====2、从键盘输入一个字符串，然后打印出来====
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// 输入和输出, fmt包

	// 输出：
	// fmt.Println()	// 打印并换行
	// fmt.Print()		// 打印不换行
	// fmt.Printf()	// 格式化打印

	// 输入：
	// fmt.Scanf()		// 从键盘输入
	// fmt.Scanln()	// 从键盘输入
	// fmt.Scan()		// 从键盘输入


	var x int
	var y float64

	fmt.Println("1、请输入一个整数，一个浮点数，")
	//读取键盘的输入，通过操作地址，赋值给x和y
	fmt.Scanln(&x, &y)
	fmt.Printf("x的数值：%d, y的数值：%f\n", x, y)

	// 未来在 IO 中，可以使用 bufio 包来读取用户输入的内容
	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到的数据：", s1)
}
*/

package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("请输入两个数字和运算符，用空格隔开，例如: 10 + 5")
	fmt.Scanln(&num1, &operator, &num2)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	default:
		fmt.Println("无效的运算符")
		return
	}

	fmt.Printf("结果: %.2f\n", result)
}
