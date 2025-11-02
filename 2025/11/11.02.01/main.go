package main
import"fmt"
func main(){
	//iota,特殊常量，可以认为是一个可以被编译器修改的常量
	const(
		err1 = iota
		err2 = iota
		err3 = iota
		err4 = iota
	)
	fmt.Println(err1,err2,err3,err4)
}