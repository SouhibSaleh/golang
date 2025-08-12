package test

import (
	"fmt"
)

type Special int

func (s *Special) Write(p []byte) (n int, err error) {
	*s += Special(len(p))
	fmt.Print("this method has been called")
	return len(p), nil
}
func (s *Special) test() {
	fmt.Println("testing is working")
}

type Meower interface {
	Meow()
}

func (s *Special) Meow() {
	fmt.Println("this method has been Meowing")
}

func Mmeow(i Meower) {
	i.Meow()
}

func InfacePlayground() {
	var s Special
	var m Meower
	m = &s
	m.Meow()

}
