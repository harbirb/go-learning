package shapes

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

// interface resolution is implicit. If the type passed in matches what the interface asks for, it will compile
// passing a Circle works because Shape.Area() will be call Circle.Area()
type Shape interface {
	Area() float64
}

// method for rectangle types only. Similar to using this.method
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Width * t.Height / 2
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
