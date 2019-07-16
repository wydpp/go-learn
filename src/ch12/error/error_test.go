package error

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, errors.New("n should be not less than 2")
	} else if n > 100 {
		return nil, errors.New("n should be not larger than 100")
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

var LessThanTwoError = errors.New("n should be not less than 2")
var LargeThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci2(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	} else if n > 100 {
		return nil, LargeThanHundredError
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestError(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		t.Error(err)
	} else {
		t.Log(v)
	}

	if v, err := GetFibonacci2(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Print("less than 2")
		} else if err == LargeThanHundredError {
			fmt.Print("large than 100")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
