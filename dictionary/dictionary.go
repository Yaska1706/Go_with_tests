package dictionary

import "errors"

type Dictionary map[string]string

var ErrDoesNotExist = errors.New("the word does not exist")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrDoesNotExist
	}
	return definition, nil

}

func (d Dictionary) Add(word string, definition string) {
	d[word] = definition
}
