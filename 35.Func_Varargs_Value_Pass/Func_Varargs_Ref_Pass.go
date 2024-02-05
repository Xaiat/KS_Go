package main

import "fmt"

// 按照数据的存储特点来划分：参数传递可以分为“值专递”和“引用传递”
// 值类型的数据：操作的数据本身、int、string、bool、float64、array...
// 引用类型的数据：操作的是数据的地址 slice(切片)、map(集合)、chan(通道)...
// 变量在内存中是存放在一定的地址上的，修改变量实际是修改变量地址处的内存
func main() {
	// 1.值传递
	// arr2 的数据是 从 arr 复制来的，所以是不同的空间
	// 修改 arr2 并不会影响 arr
	// 值传递的核心：传递的是数据的副本，修改数据，对于原始的数据没有影响
	// 值类型的数据，默认都是值传递：基础类型、array数组、struct结构体
	// 定义一个数组 [个数]类型
	arr := [4]int{1, 2, 3, 4}
	// 用[]中括号打印出来的都是数组的值
	fmt.Println(arr)
	// 传递：执行了拷贝arr动作
	update(arr)
	fmt.Println("调用修改后的数据：", arr)
	// 2.引用传递
}

func update(arr2 [4]int) {
	fmt.Println("arr2接收的数据：", arr2)
	// 数组里的值都是从0开始的
	arr2[0] = 100
	fmt.Println("arr2修改后的数据：", arr2)
}
