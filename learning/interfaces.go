package learning

import (
	"math"
)

type Shape interface {
	area() float64
}
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.Radius * c.Radius

}
func Calculate(s Shape) float64 {
	return s.area()

}
