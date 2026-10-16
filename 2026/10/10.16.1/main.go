package main

import "fmt"

func main() {
	var num1 int // 先定义
	num1 = 10    // 后赋值
	fmt.Println("num1 = ", num1)

	var num2 int = 20 // 定义的同时赋值
	fmt.Println("num2 = ", num2)

	var num3 = 30 // 定义的同时赋值, 并省略数据类型
	fmt.Println("num3 = ", num3)

	num4 := 40 // 定义的同时赋值, 并省略关键字和数据类型
	/*
		num4  := 40 等价于
		var num4 int
		num4 = 40
	*/
	fmt.Println("num4 = ", num4)
}
