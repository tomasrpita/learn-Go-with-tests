package dictionary

const (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrWordExists    = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesFound = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound

	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesFound
	}

	d[word] = definition
	return nil
}
