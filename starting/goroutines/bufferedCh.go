package goroutines

import (
	"fmt"
)

func generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("generating: ", i)
		ch <- i
	}
	close(ch)
}
func reslove(ch <-chan int) {
	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("ch is not ok")
			break
		}
		fmt.Println("resolve", i)
	}
}

func BufferedCh() {

	ch1 := make(chan int, 2)

	go generate(ch1)
	reslove(ch1)

}
