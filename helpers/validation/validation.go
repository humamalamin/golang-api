package validationHelpers

import (
	"strconv"
	"strings"
	"unicode"
)

// Validate That String Data Is A Digit
func StrIsDigit(data string) error {
	for _, v := range data {
		isDigit := unicode.IsDigit(v)
		if !isDigit {
			return ErrorStrIsNotDigit
		}
	}

	return nil
}

// Validation check digit
func IsDigit(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	if _, err := strconv.Atoi(s); err != nil {
		return false
	}

	return true
}
