package goroutines

import (
	"fmt"
	"sync"
	"time"
)

func solver(x int, ch <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Println(x)
	time.Sleep(time.Second * 3)
	<-ch
	wg.Done()

}
func selector(wg *sync.WaitGroup) {
	wg.Done()

}

func GoControl() {
	ch := make(chan struct{}, 5)

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {

		ch <- struct{}{}
		wg.Add(1)
		go solver(i, ch, &wg)
	}
	go selector(&wg)

}
