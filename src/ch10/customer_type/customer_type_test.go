package customer_type

import (
	"fmt"
	"testing"
	"time"
)

//自定义类型
type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 2)
	return op
}

func TestCustomerType(t *testing.T) {
	f := timeSpent(slowFun)
	i := f(10)
	t.Log(i)
}
