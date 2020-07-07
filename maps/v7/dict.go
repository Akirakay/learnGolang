package maps

//Dict is an original map
type Dict map[string]string

//DictErr is an original string
type DictErr string

const (
	//ErrNotFound means 404
	ErrNotFound = DictErr("could not find the word you were looking for")
	//ErrWordExists means word already exists
	ErrWordExists = DictErr("cannot add word because it already exists")
)

// Error return stringfy of e
// implement of error
func (e DictErr) Error() string {
	return string(e)
}

// Search return the value by the received word
func (d Dict) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add add element to this dict
func (d Dict) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

//Update means change the value by word
func (d Dict) Update(word, newDefinition string) {
	d[word] = newDefinition
}
