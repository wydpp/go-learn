package encapsutation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee)
	e2.Age = 22
	e2.Id = "2"
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

func (e Employee) String() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("e.id=%s,e.name=%s,e.age=%d", e.Id, e.Name, e.Age)
}

func (e *Employee) Stringp() string {
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("e.id=%s,e.name=%s,e.age=%d", e.Id, e.Name, e.Age)
}

func TestFuncMethod(t *testing.T) {
	e := Employee{"0", "Bob", 20}
	fmt.Printf("Address is %x\n", unsafe.Pointer(&e.Name))
	//t.Log(e.String())
	//stringp打印的name地址和e对象的相同
	t.Log(e.Stringp())
}
