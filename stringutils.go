package go_stringutils

import (
	"unicode/utf8"
)

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
