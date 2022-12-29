package dictionary

const (
	// ErrNotFound means the definition could not be found for the given word.
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists means you are trying to add a word that is already known.
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	// ErrWordExists occurs when trying to update a word not in the dictionary.
	ErrWordDoesFound = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr are errors that can happen when interacting with the dictionary.
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary store definitions to words.
type Dictionary map[string]string

// DictionaryErr are errors that can happen when interacting with the dictionary.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound

	}
	return definition, nil
}

// Add inserts a word and definition into the dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, ok := d[word]
	if ok {
		return ErrWordExists
	}
	d[word] = definition
	return nil
}

// Update changes the definition of a given word.
func (d Dictionary) Update(word, definition string) error {
	_, ok := d[word]
	if !ok {
		return ErrWordDoesFound
	}

	d[word] = definition
	return nil
}

// Delete removes a word from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
