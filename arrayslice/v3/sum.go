package slice

// Sum return sum of slice members
func Sum(numbers []int) int {
	var sum int
	for _, item := range numbers {
		sum += item
	}
	return sum
}
