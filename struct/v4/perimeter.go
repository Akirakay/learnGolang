package structdemo

// Perimeter return the perimeter of shape
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.width + rectangle.height)
}

// Area return the area of shape
func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}
