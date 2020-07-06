package structdemo

// Rectangle is a type
type Rectangle struct {
	width  float64
	height float64
}

// Area return the area of this shape
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter return the perimeter of this shape
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}
