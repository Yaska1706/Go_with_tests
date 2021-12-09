package sminter

import "math"

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}
type Square struct {
	Length float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	perimeter := 2 * (r.Length + r.Width)
	return perimeter
}

func (r Rectangle) Area() float64 {
	area := r.Length * r.Width
	return area
}

func (c Circle) Area() float64 {
	area := math.Pi * c.Radius * c.Radius
	return area
}

func (t Triangle) Area() float64 {
	area := 0.5 * t.Base * t.Height
	return area
}

func (s Square) Area() float64 {
	area := s.Length * s.Length
	return area
}
