package maps

import "errors"

//Dict is an original map
type Dict map[string]string

//ErrNotFound represent 404
var ErrNotFound = errors.New("could not find the word you were looking for")

// Search return the value by the received word
func (d Dict) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add add element to this dict
func (d Dict) Add(word, definition string) {
	d[word] = definition
}
