package utils

import (
	"errors"
	"unicode"
)

var ErrAlphabetPositionInput = errors.New("input error: only positions between 1 and 26 is accepted")

// Takes in a single alphabet of lower or uppercase and returns its position in the English alphabet
func AlphabetPosition(alphabet rune) (int, error) {

	return int(unicode.ToUpper(alphabet) - 'A' + 1), nil
}

//Takes in an int and returns the corresponding Uppercase letter in the English alphabet (1 indexed i.e. A is 1)
func GetAlphabetInPosition(i int32) (rune, error) {
	if i > 26 || i < 1 {
		return ' ', ErrAlphabetPositionInput
	}

	return 'A' - 1 + i, nil
}
