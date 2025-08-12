package test

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct {
	X, Y int
}
type ColoredPoint struct {
	point Point
	color.RGBA
}
type Shape []Point

func (p *Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}
func (p *Point) Distance(o Point) float64 {
	fmt.Println("from Point")
	return math.Sqrt(math.Pow(float64(p.X-o.X), 2) +
		math.Pow(float64(p.Y-o.Y), 2))
}
func (p *ColoredPoint) Distance(o Point) float64 {
	fmt.Println("from ColoredPoint")
	return p.point.Distance(o)
}

//func (s Shape) String() string {
//	return "noice"
//}

func rec() {
	p := &Point{1, 2}
	p = nil
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println(p.X)

}

func Draw() {
	p := &Point{1, 2}
	p1 := Point{3, 4}
	fmt.Println(p)
	fmt.Println(&p1)

	s := Shape{*p, *p, *p}
	fmt.Println(s)
	s = append(s, Point{1, 2})
	fmt.Println(s)
	fmt.Println(p)

	cp := ColoredPoint{Point{1, 2}, color.RGBA{255, 0, 0, 255}}
	cp.point.X = 4
	fmt.Println(cp)
	cp.Distance(Point{1, 2})
	ss := Point{3, 4}
	fmt.Println(ss.String())
}
