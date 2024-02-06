package main

import "fmt"

// 递归函数 Func_Recur (Function Recursive)
// 定义：一个函数自己调用自己，就叫做递归函数
// 注意：递归函数需要有一个出口，逐渐向出口靠近，没有出口就会形成死循环

// 想要做一个求和：1，2，3，4，5
// getsum(5)
//   getsum(4)+5
//	   getsum(3)+4
//       getsum(2)+3
//		   getsum(1)+1
//			 getsum(1)就是1本身

func main() {

	sum := getSum(5)
	fmt.Println(sum)
}

// 5
// getSum(5)=getSum(4)+5
// gerSum(4)+5=gerSum(3)+4+5
// gerSum(3)+4+5=gerSum(2)+3+4+5
// gerSum(2)+3+4+5=gerSum(1)+2+3+4+5
// gerSum(1)+2+3+4+5=1+2+3+4+5=15

func getSum(n int) int {

	if n == 1 {
		return 1
	}

	return getSum(n-1) + n
}

/*
这段代码定义了一个名为 getSum 的递归函数，它的目的是计算从 1 到 n 的所有整数的和。递归是一种编程技术，函数在执行过程中调用自身。

在 getSum 函数中，我们首先检查参数 n 是否为 1。如果是，函数就返回 1，这是递归的基本情况。如果 n 不是 1，函数就返回 getSum(n-1) + n。
这是递归的递归情况，函数调用自身，但是每次调用时 n 的值都减少 1，这样就可以逐渐接近基本情况。

例如，如果我们调用 getSum(5)，函数将执行以下操作：

检查 5 是否为 1。因为 5 不是 1，所以函数返回 getSum(4) + 5。
为了计算 getSum(4)，函数再次调用自身，这次 n 的值为 4。函数返回 getSum(3) + 4。
这个过程会一直持续下去，直到 n 的值为 1。此时，函数返回 1。
然后，所有的函数调用开始返回，每次返回时都会加上 n 的当前值。所以，最终的结果是 1 + 2 + 3 + 4 + 5 = 15。
这就是递归函数的工作原理。它们可以非常强大，但也需要小心使用，因为如果没有正确的基本情况或者递归情况，它们可能会导致无限循环。
*/
