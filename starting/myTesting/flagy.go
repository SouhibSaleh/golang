package test

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func (p *Point) Set(s string) error {
	args := strings.Split(s, ",")
	var tp Point
	tp.X, _ = strconv.Atoi(args[0])
	tp.Y, _ = strconv.Atoi(args[1])
	p.X = tp.X
	p.Y = tp.Y
	return nil
}
func FlagPlayground() {

	num := flag.Int64("number", 0, "any number")

	var p Point
	flag.Var(&p, "point", "A point in the format X,Y")

	flag.Parse()

	fmt.Println(*num)
	fmt.Println(p)

	fmt.Println("Other args:", flag.Args())

}
