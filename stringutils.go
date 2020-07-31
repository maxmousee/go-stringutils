package stringutils

import (
	"strings"
	"unicode"
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

// ContainsOnly checks if the string has only a specific subset of chars
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

// ContainsNone checks if the string has none of a specific subset of chars
func ContainsNone(str string, disallowedChars string) bool {
	if len(disallowedChars) == 0 {
		return true
	}
	if str == disallowedChars {
		return false
	}
	for _, char := range str {
		if strings.Contains(disallowedChars, string(char)) {
			return false
		}
	}
	return true
}

// EqualsAny checks if a string is equal to any of the strings in an array
func EqualsAny(str1 string, str2 []string) bool {
	for _, aString := range str2 {
		if str1 == aString {
			return true
		}
	}
	return false
}

// EndsWithAny checks if a string ends with any of the strings in an array
func EndsWithAny(str1 string, endings []string) bool {
	for _, aSuffix := range endings {
		if strings.HasSuffix(str1, aSuffix) {
			return true
		}
	}
	return false
}

// StartsWithAny checks if a string starts with any of the strings in an array
func StartsWithAny(str1 string, endings []string) bool {
	for _, aSuffix := range endings {
		if strings.HasPrefix(str1, aSuffix) {
			return true
		}
	}
	return false
}

// EqualsIgnoreCase checks if strings are equal ignoring case
func EqualsIgnoreCase(str1 string, str2 string) bool {
	return strings.EqualFold(str1, str2)
}

// EqualsAnyIgnoreCase checks if strings are equal to any string in the array, ignoring case
func EqualsAnyIgnoreCase(str1 string, str2 []string) bool {
	for _, aString := range str2 {
		if EqualsIgnoreCase(str1, aString) {
			return true
		}
	}
	return false
}

// ContainsIgnoreCase checks if str1 contains str2, ignoring case. (Unicode)
func ContainsIgnoreCase(str1 string, str2 string) bool {
	return strings.Contains(strings.ToLower(str1), strings.ToLower(str2))
}

// ContainsAnyIgnoreCase checks if str1 contains any of the strings in str2, ignoring case. (Unicode)
func ContainsAnyIgnoreCase(str1 string, str2 []string) bool {
	for _, aString := range str2 {
		if ContainsIgnoreCase(str1, aString) {
			return true
		}
	}
	return false
}

// DeleteWhitespace deletes all whitespaces from a given string
func DeleteWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}

// AppendIfMissing appends a given suffix to a string if the string does not ends with it
func AppendIfMissing(str string, suffix string) string {
	if strings.HasSuffix(str, suffix) {
		return str
	}
	return str + suffix
}

const minCJKCharacter = '\u3400'

// IsAlphabet checks r is a letter but not CJK character.
func IsAlphabet(r rune) bool {
	if !unicode.IsLetter(r) {
		return false
	}

	switch {
	// Quick check for non-CJK character.
	case r < minCJKCharacter:
		return true

	// Common CJK characters.
	case r >= '\u4E00' && r <= '\u9FCC':
		return false

	// Rare CJK characters.
	case r >= '\u3400' && r <= '\u4D85':
		return false

	// Rare and historic CJK characters.
	case r >= '\U00020000' && r <= '\U0002B81D':
		return false
	}
	return true
}

// WordSplit splits a string into words. Returns a slice of words.
// If there is no word in a string, return nil.
//
// Word is defined as a locale dependent string containing alphabetic characters,
// which may also contain but not start with `'` and `-` characters.
func WordSplit(str string) []string {
	if DeleteWhitespace(str) == "" {
		return []string{""}
	}

	var word string
	var words []string
	var r rune
	var size, pos int

	inWord := false

	for len(str) > 0 {
		r, size = utf8.DecodeRuneInString(str)

		switch {
		case IsAlphabet(r):
			if !inWord {
				inWord = true
				word = str
				pos = 0
			}

		case inWord && (r == '\'' || r == '-'):
			// Still in word.

		default:
			if inWord {
				inWord = false
				words = append(words, word[:pos])
			}
		}

		pos += size
		str = str[size:]
	}

	if inWord {
		words = append(words, word[:pos])
	}

	return words
}
