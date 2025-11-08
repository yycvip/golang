package main
import"fmt"

func a()(int,bool){
	return 0, false
}
func main(){

	//匿名变量
	var _ int
	_,ok :=a()
	fmt.Println(ok)
	//iota,特殊常量，可以认为是一个可以被编译器修改的常量
	const(
		err1 = iota+1
		err2 //= iota
		err3 //= iota
		err4 //= iota
	)
	fmt.Println(err1,err2,err3,err4)
	const(
		err00 = iota
	)
	fmt.Println(err00)
	/*
	如果中断了iota那么必须显示的恢复
	*/
}