package dictionary

const (
	// ErrNotFound means the definition could not be found for a given word
	ErrNotFound = DictionaryErr("word not found")

	// ErrWordsExists means you are trying to add a word that already exists
	ErrWordsExists = DictionaryErr("cannot add word alreay exists")

	// ErrWordDoesNotExist means that you are trying to update a word not found in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word, it does not exist")
)

// DictionaryErr are errors that can happen when interactinf with the dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary stores words and definitions
type Dictionary map[string]string

// Search find a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add inserts a word and defintion into the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordsExists
	default:
		return err
	}

	return nil
}

// Update changes the definition of a given word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete removes a given word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
