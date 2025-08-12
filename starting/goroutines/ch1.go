package goroutines

import "fmt"

func solve(i int) {
	for x := 0; x < 300; x++ {
		fmt.Println(i)
	}
}
func Ch1() {
	for i := 0; i < 3; i++ {
		go solve(i)
		fmt.Println("started:", i)
	}
	for {

	}

}
