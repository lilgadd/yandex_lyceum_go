package main

import (
	"errors"
	"unicode/utf8"
)
func GetUTFLength(input []byte) (int, error) {
	var ErrInvalidUTF8 = errors.New("invalid utf8")
	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}