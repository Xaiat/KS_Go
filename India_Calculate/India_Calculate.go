package main

import (
	"fmt"
)

// IndianMultiplication 计算并展示印度乘法心算法的步骤
func IndianMultiplication(a, b int) int {
	// 步骤1: 确保a是较大的数
	if a < b {
		a, b = b, a
	}

	// 步骤2: 大数增加小数个位数
	step2 := a + b%10

	// 步骤3: 结果乘以10
	step3 := step2 * 10

	// 步骤4: 乘个位数
	step4 := (a % 10) * (b % 10)

	// 步骤5: 求和
	result := step3 + step4

	// 展示计算步骤
	fmt.Printf("大数增加小数个位数: %d\n", step2)
	fmt.Printf("中间结果乘以10: %d\n", step3)
	fmt.Printf("两个数个位数相乘: %d\n", step4)
	fmt.Printf("最终结果: %d\n", result)

	return result
}

func main() {
	var a, b int
	fmt.Println("请输入两个数字进行印度乘法心算（例如：13 14）：")
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		fmt.Println("输入错误，请确保输入两个整数。")
		return
	}

	fmt.Printf("%d 乘以 %d 的印度乘法心算结果是：\n", a, b)
	IndianMultiplication(a, b)
}
