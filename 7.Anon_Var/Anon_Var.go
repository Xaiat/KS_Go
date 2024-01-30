package main

import "fmt"

func test() (int, int) {
	return 100, 200
}

func main() {
	a, _ := test()
	_, b := test()
	fmt.Println("a = ", a, "b = ", b)
}
