package learning

import (
	"math"
)

// we can use both method in same interface but for laearning we are using this
type Shape interface {
	Area() float64
}
type Shape2 interface {
	Perim() float64
}
type Rectangle struct {
	Width, Height float64
}

// we we use pointer we simply has to add & operator in the main in slice
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func (r Rectangle) Perim() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

func Calculate(s Shape) float64 {
	return s.Area()

}
func Calculate1(s Shape2) float64 {

	return s.Perim()
}
