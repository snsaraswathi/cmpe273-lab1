package main
import "math"
type Shaper interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	length, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
 }

func (r Rectangle) Perimeter() float64 {
	return 2*(r.length + r.width)
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.radius
}
