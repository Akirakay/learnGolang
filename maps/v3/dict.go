package maps

import "errors"

//Dict is an original map
type Dict map[string]string

// Search return the value by the received word
func (d Dict) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}

	return definition, nil
}
