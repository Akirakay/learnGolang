package iteration

const repeation = 5

// Repeat return 5X of seed
func Repeat(seed string) string {
	var char string
	for i := 0; i < repeation; i++ {
		char += seed
	}
	return char
}
