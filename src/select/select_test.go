package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 800)
	return "Done"
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

/**
超过一定时间返回错误，类似java timeout
*/
func TestAsyncService(t *testing.T) {
	select {
	case ret := <-asyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 500):
		t.Error("time out")
	}
}
