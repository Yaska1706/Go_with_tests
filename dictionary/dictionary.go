package dictionary

import "errors"

type Dictionary map[string]string

var (
	ErrDoesNotExist = errors.New("the word does not exist")
	ErrWordExists   = errors.New("word already exists")
	ErrWordNotFound = errors.New("word is not found")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDoesNotExist
	}
	return definition, nil

}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrDoesNotExist:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrDoesNotExist:
		return ErrWordNotFound
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
