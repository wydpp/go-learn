package error

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

/**
错误恢复

Let it Crash! 往往时我们恢复不确定性错误的最好方法。
*/
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err)
		}
	}()
	fmt.Println("Start!")
	panic(errors.New("NullPointException!"))
	fmt.Println("End!")
}

func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("defer!")
	}()
	fmt.Println("Start!")
	panic(errors.New("NullPointException!"))
	fmt.Println("End!")
}
