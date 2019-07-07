package _func

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
返回多个值
*/
func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spend:", time.Since(start).Seconds())
		return ret
	}
}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFunc(t *testing.T) {
	a, b := returnMultiValue()
	t.Log(a, b)

	tsff := timeSpent(slowFunc)
	v := tsff(10)
	t.Log(v)
}

//可变参数
func sum(op ...int) int {
	ret := 0
	for _, op := range op {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(sum(1, 2, 3, 4))
	t.Log(sum(1, 2, 3, 4, 5))
}

func clear() {
	fmt.Println("Clear resource.")
}

//延迟执行defer. java-finally?
func TestDefer(t *testing.T) {
	//defer永远会执行？
	defer clear()
	fmt.Println("Start")
	//java 异常？
	panic("err")

}
