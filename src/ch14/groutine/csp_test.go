package groutine

import (
	"fmt"
	"testing"
	"time"
)

//CSP vs Actor

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Other task is done.")
}

func TestCsp(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func asyncService() chan string {
	//retCh := make(chan string)
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result。")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := asyncService()
	otherTask()
	fmt.Println(<-retCh)

}
