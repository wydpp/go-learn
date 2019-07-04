package main

import (
	"fmt"
	"math"
)

/*
import测试
 */
func main() {
	fmt.Printf("Now you hava %g problems.\n", math.Sqrt(7))
	//再go语言中，如果一个名字以大写字母开头，表示它可以再包外被访问的，类似java中的public static
	fmt.Printf("圆周率是%g", math.Pi)
}
