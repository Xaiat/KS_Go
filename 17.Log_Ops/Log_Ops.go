package main

import "fmt"

func main() {

	var a bool = true
	var b bool = true
	// 逻辑运算符`&&`就是`与`, 我与你，就是我和你都要满足一个结果，才执行 (理解为：和、并且)
	// a 和 b 都是 true 才会执行, 只要有一个是 false 就不会执行
	// 总结：都为真，结果才为真(执行)，否则为假(不执行)
	if a == true && b == true {
		fmt.Println(a && b)
	}

	var c bool = true
	var d bool = false
	// 逻辑运算符`||`就是`或`, 我或你，就是我和你只要有一个满足结果，就执行 (理解为：或、或者)
	// c 或 d 只要有一个是 true 就会执行, 都是 false 才不会执行
	// 总结：只要有一个为真，结果就为真(执行)，否则为假(不执行)
	fmt.Println(c || d)

	var e bool = true
	var f bool = false
	// 逻辑运算符`!`就是`非`, 非我即你，就是我和你只要有一个不满足结果，就执行 (理解为：非、不是)
	// e 和 f 无论是什么，！e 和 !f 都是相反的结果
	// 总结：`！`就是取反，true 变 false，false 变 true
	fmt.Println(!e)
	fmt.Println(!f)
}
