package slice

// SumAll return sum of slice members in each slice
func SumAll(arrs ...[]int) (sum []int) {
	lengthOfNumbers := len(arrs)
	sum = make([]int, lengthOfNumbers)

	for i, numbers := range arrs {
		sum[i] = Sum(numbers)
	}
	return
}

// Sum return sum of slice members
func Sum(numbers []int) int {
	var sum int
	for _, item := range numbers {
		sum += item
	}
	return sum
}
