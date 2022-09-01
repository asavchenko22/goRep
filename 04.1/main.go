package main

import (
	"math"
)

const pi = 3.14

type AreAre interface {
}

type Rectangle struct {
	bodyType string
	width    float64
	height   float64
}

func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

func (r *Rectangle) Type() string {
	return r.bodyType
}

type Circle struct {
	bodyType string
	radius   float64
}

func (c *Circle) Type() string {
	return c.bodyType
}

func (c *Circle) Area() float64 {
	return pi * (math.Pow(c.radius, 2))
}

func figureArea(f AreAre) {
	switch f.(type) {
	case *Rectangle:
		rect := f.(*Rectangle)
		rect.Area()
	case *Circle:
		circle := f.(*Circle)
		circle.Area()
	}

}

//func figureType(a AreAre) {
//	fmt.Println(a.Type())
//}

func main() {
	rect := &Rectangle{"rectangle", 5, 5}

	figureArea(rect)

	circle := &Circle{"circle", 5}

	figureArea(circle)
}
