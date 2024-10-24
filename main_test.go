package main

import (
	"testing"
	"errors"
)
var ErrInvalidUTF8 = errors.New("invalid utf8")
func TestDeleteVowels(t * testing.T){
	invalidString := "\xff\xfe\xfd"
	invalidString1 := "\xC3"
	type Test struct{
		got string
		expected int
		err error
	}
	var tests = []Test{
		{"abcd", 4, nil},
		{invalidString, 0, ErrInvalidUTF8},
		{"привяо", 6, nil},
		{invalidString1, 0, ErrInvalidUTF8},
	}
	for _, test := range tests{
		result, err := GetUTFLength([]byte(test.got))
		if result != test.expected{
			t.Errorf("GetUTFLength([]byte(%q) =  %d, want %d)", test.got, result, test.expected)
		}
		if !errors.Is(err, test.err){
			t.Errorf("GetUTFLength([]byte(%q) =  %v, want %v)", test.got, err, test.err)
		}
	}

}