package main

import "fmt"

func main(){
	var age int = 18
	//&符号 + 变量  就可以获取这个变量内存的地址
	fmt.Println(&age) //0xc000010110

	//定义一个指针变量：
	// var代表要声明一个变量
	// ptr 指针变量的名字
	// ptr对应的类型是：*int是一个指针类型（可以理解为指向int类型的指针）
	// &age就是一个地址，是ptr变量的具体值
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr本身的存储空间的地址是：",&ptr)
	fmt.Printf("ptr指向的数值为：%v",*ptr)
	//总结：指针就是内存地址
		// 最重要的就是两个符号：
		// 1.&取内存地址；
		// 2.*根据地址取值。
}