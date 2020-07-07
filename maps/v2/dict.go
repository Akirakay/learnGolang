package maps

// Search return the value by the received word
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
