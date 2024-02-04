package main

import "fmt"

// arg ...int 告诉 Go 这个函数接收不定数量的参数，类型全部是 int
func main() {
	// 调用函数时，可以传递任意数量的参数
	getSum("Xaiat", 1, 2, 3, 4, 5, 6, 7, 8, 100)
}

func getSum(msg string, nums ...int) {
	sum := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
		// sum = sum + nums[i], sum += nums[i] 是一样的
		sum += nums[i]
	}
	fmt.Println("msg:", msg)
	fmt.Println("sum:", sum)

}
