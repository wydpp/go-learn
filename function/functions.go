package main

import "fmt"

/*
 函数可以没有参数或接受多个参数，返回参数也可以没有，有的话写在 ) 后面
 和java不同的是类型在变量名之后，函数返回类型也在后面
 */

func hello(s string){
	fmt.Print(s)
}

func add(x int, y int) int {
	return x + y
}
/*
当连续两个或多个已命名形参类型相同时，除最后一个类型外，其他的都可以省略
 */
func plus(x,y int) int {
	return x * y
}
/*
函数可以返回任意数量的返回值
swap返回了两个字符串
 */
func swap(x,y string) (string,string) {
	return y,x
}
/*
go的返回值可被命名，他们会被视作定义在函数顶部的变量
返回值的名称应当具有一定的意义，它可以做为文档来使用。
没有参数的return语句返回已命名的返回值。也就是直接返回。
 */
func split(sum int) (x,y int){
	x = sum * 4 /9
	y = sum - x
	return
}

/*
 */
func main() {
	fmt.Println(add(1,1))
	hello("你好\n")

	fmt.Println(plus(1,1))

	x :="你好"
	y := "世界"
	fmt.Println(x,y)
	x,y = swap(x,y)
	fmt.Println(x,y)

	a,b := split(10)
	fmt.Println(a,b)
}
