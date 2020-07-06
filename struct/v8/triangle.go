package structdemo

// Triangle is a struct
type Triangle struct {
	base   float64
	height float64
}

// Area return the area of this shape
func (t Triangle) Area() float64 {
	return (t.base * t.height) / 2
}
