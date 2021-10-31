package main

import (
	"fmt"
)

func main()  {
	//指定变量类型，声明后若不赋值，使用默认值
	var num bool
	fmt.Println("num=",num)
	//根据值自行判断变量类型
	var i = 10.22
	fmt.Println("i=",i)
	//省略var
	name := "jiangjun"
	fmt.Println("name=",name)
	var char byte
	char='a'
	fmt.Println("char=",char)
	var a= "hello"
	var b = "world"
	var c bool
	println(a,b,c)
	//多变量声明
	/*	var name1,name2,name3 =v1,v2,v3
		println(name1,name2,name3)*/
	/*注意事项
	1.变量必须先定义才能使用
	2.go语言是静态变量，要求变量得类型和赋值得类型必须一致
	3.变量名不能冲突
	4.简短定义方式，左边得变量名至少有一个是新的
	5.简短定义方式，不能定义全局变量
	6.变量的零值也叫默认值
	7.定义变量后就要使用*/
	x := 20;
	println("x=",&x)
	x,y :=50, "jiangjun"
	println(&x,x)
	println(y)
	//基本语法 常量constant
	/*const indentifier [type] =value*/
    //显示类型 const b string = "abc"
//	隐式类型定义：const b = "abc"
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const j , k , l = 1,false,"str"
	area = LENGTH*WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(j,k,l)
//	常量可以作为枚举，常量组
	const (
		Unknown uint16 = 23
		Female
		Male = "20"
		Kavent
	)
	//常量组中如不指定类型和初始化值，则与上一行非空变量右值相同
	fmt.Printf("%T,%v\n",Female, Female)
	fmt.Printf("%T,%v\n",Kavent,Kavent)
}
