package main

import "fmt"

//构造函数

type person struct{
	name string
	age int
}

type dog struct{
	name string
}

func newPerson(name string,age int)person{
return person{
	name:name,
	age:age,
}
}

func newDog(name string) dog{
return dog {
	name:name,
}
}
func main(){
	p1:=newPerson("媛媛",18)
	p2:=newPerson("小小",900)
	fmt.Println(p1)
	fmt.Println(p2)
	d1:=newDog("小狗")
	fmt.Println(d1)
}
