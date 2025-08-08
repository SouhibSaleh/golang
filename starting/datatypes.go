package main

import "fmt"

type WrapperA float64
type WrapperB float64

func (a WrapperA) String() string {
	return fmt.Sprintf("%f $$", a)
}

func dataTest() {
	first := WrapperA(1)
	second := WrapperB(first)
	fmt.Println(first*2 + WrapperA(second)*0.2)

}
