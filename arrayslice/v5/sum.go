package slice

// SumAll return sum of slice members in each slice, append result to tail
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Sum return sum of slice members
func Sum(numbers []int) int {
	var sum int
	for _, item := range numbers {
		sum += item
	}
	return sum
}
