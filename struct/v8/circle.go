package structdemo

import "math"

// Circle is a shape
type Circle struct {
	radius float64
}

// Area return the area of this shape
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Perimeter return the perimeter of this shape
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
