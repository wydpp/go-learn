package fib

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	//var (
	//	a = 1
	//	b = 1
	//)
	a := 1
	b := 1
	fmt.Print(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
	fmt.Println()
}
