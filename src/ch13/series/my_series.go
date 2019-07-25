package series

import "fmt"

func init() {
	fmt.Println("i am init method!")
}

func init() {
	fmt.Println("i am init2 method!")
}

func GetFibonacciSerie(n int) []int {
	fmt.Println("i am GetFibonacciSerie method")
	ret := []int{1, 1}
	for i := 2; i < n; i++ {
		ret = append(ret, ret[i-2]+ret[i-1])
	}
	return ret
}
