package iteration

// Repeat return 5X of seed
func Repeat(seed string) string {
	var char string
	for i := 0; i < 5; i++ {
		char += seed
	}
	return char
}
