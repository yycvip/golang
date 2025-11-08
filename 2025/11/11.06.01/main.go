package main

import "fmt"

func main(){
	// localName := "local"
	// fmt.Println(localName)

	var mname int
	 	mname = 19
	if  mname <= 3 {
		fmt.Println("我在上幼儿园")
	}else if mname <=12 {
		fmt.Println("我在上小学")
	}else if mname <= 15{
		fmt.Println("我在上初中")
	}else if mname <= 18{
		fmt.Println("我在上高中")
	}else if mname <= 22{
		fmt.Println("我在上大学")
	}else{
		fmt.Println("我毕业了，工作了")
	}

}
