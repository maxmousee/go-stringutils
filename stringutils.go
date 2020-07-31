package go_stringutils

import (
	"strings"
	"unicode/utf8"
)

// Abbreviate an string to a specific number of chars...
// For example, Abbreviate("abcdef", 5) returns ab...
// Abbreviate("abcdef", 6) returns "abcdef"
// If the length is an invalid value, such as less than 4,
// a negative value or a value larger than the original string length
// this function returns the original string
func Abbreviate(str string, length int) string {
	if length >= len(str) {
		return str
	}
	if length <= 3 {
		return str
	}
	strLenght := len(str) - length + 1
	str = str[:strLenght] + "..."
	return str
}

// Reverse an UTF-8 encoded string
func Reverse(str string) string {
	var size int
	tail := len(str)
	buf := make([]byte, tail)
	s := buf

	for len(str) > 0 {
		_, size = utf8.DecodeRuneInString(str)
		tail -= size
		s = append(s[:tail], []byte(str[:size])...)
		str = str[size:]
	}
	return string(buf)
}

// Checks if the string has only a specific subset of chars
func ContainsOnly(str string, allowedChar string) bool {
	if len(allowedChar) == 0 {
		return false
	}
	if str == allowedChar {
		return true
	}
	for _, char := range str {
		if !strings.Contains(allowedChar, string(char)) {
			return false
		}
	}
	return true
}
