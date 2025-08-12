package besting

import (
	"fmt"
	"os"
)

func starting() {
	var s string
	//	var s = ""
	//	s:=" "

	for i := 1; i < len(os.Args); i++ {
		s += os.Args[i]
	}
	fmt.Println(s)
	s = ""

	for _, arg := range os.Args[1:] {
		s += arg
	}
	fmt.Println(s)
}
