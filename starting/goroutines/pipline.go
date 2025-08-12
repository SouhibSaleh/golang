package goroutines

import (
	"fmt"
	"math"
)

func Generate(ch chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Generated ", i)

		ch <- i
	}
	close(ch)
}
func Compute(ch1 <-chan int, ch2 chan<- int) {
	for i := 0; i < 10; i++ {
		m := <-ch1
		fmt.Println("Computed ", m)
		ch2 <- m * m
	}
	close(ch2)
}
func Show(ch <-chan int) {
	for i := 0; i < 11; i++ {
		m, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Println("Showed ", math.Sqrt(float64(m)))
	}
}

func Pipline() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	fmt.Println("Starting...")
	go Generate(ch1)

	go Compute(ch1, ch2)
	Show(ch2)

}
