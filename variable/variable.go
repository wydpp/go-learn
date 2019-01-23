package main

import (
	"fmt"
	"strconv"
)

//var 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后
//var 语句可以出现在包或函数级别
var c,python,java bool

//变量声明可以包含初始值，每个变量对应一个。
//如果初始化值已存在，则可以省略类型；便量会从初始值中获取类型。
var i,j = 1,2

var (
	cust string = "lili"
	age int = 23
)

//产量声明
const name = "dpp"

func main() {
	var a int
	fmt.Println(a,c,python,java)

	fmt.Println(i,j)

	fmt.Println("常量name="+name)

	//在函数中，简洁赋值语句 := 可以在类型明确的地方代替var声明。
	//函数外的每个语句都必须以关键字开始（var func等），因此 := 结构不能再函数外使用
	name := "jerry"
	fmt.Println("函数内部变量name="+name)

	fmt.Println("your name is "+name+" and age is "+ strconv.Itoa(age))
}
