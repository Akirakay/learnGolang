package array

// Sum return sum of array members
func Sum(numbers [4]int) int {
	var sum int
	for i := 0; i < 4; i++ {
		sum += numbers[i]
	}
	return sum
}
