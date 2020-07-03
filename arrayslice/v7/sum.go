package slice

// SumAllTails return sum of slice tail member in each slice
func SumAllTails(arrs ...[]int) []int {
	var sums []int
	for _, item := range arrs {
		if len(item) == 0 {
			sums = append(sums, 0)
		} else {
			tail := item[1:]
			sums = append(sums, Sum(tail))
		}
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
