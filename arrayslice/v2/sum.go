package array

// Sum return sum of array members
func Sum(numbers [4]int) int {
	var sum int
	for _, item := range numbers {
		sum += item
	}
	return sum
}
