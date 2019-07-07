package _interface

import "testing"

/**
接口
*/
type Programmer interface {
	WriteHelloWorld() string
}

/**
实现
*/
type GoProgrammer struct {
}

//方法实现
func (p *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"Hello World!\")"
}

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())

}
