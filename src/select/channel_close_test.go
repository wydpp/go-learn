package _select

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
}

func dataReveiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("chan closed")
				break
			}
		}
	}()
}

func TestData(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReveiver(ch, &wg)
	//wg.Add(1)
	//dataReveiver(ch,&wg)
	wg.Wait()
}
