package go_stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assertions := assert.New(t)

	result := Reverse("")
	assertions.Equal("", result)

	result = Reverse("gnirts esrever")
	assertions.Equal("reverse string", result)

	result = Reverse("中文如何？")
	assertions.Equal("？何如文中", result)

	result = Reverse("中en文混~排怎样？a")
	assertions.Equal("a？样怎排~混文ne中", result)
}

func TestAbbreviate(t *testing.T) {
	assertions := assert.New(t)

	result := Abbreviate("", 1)
	assertions.Equal("", result)

	result = Abbreviate("abcdef", 5)
	assertions.Equal("ab...", result)

	result = Abbreviate("abcdef", 6)
	assertions.Equal("abcdef", result)

	result = Abbreviate("abcdef", -1)
	assertions.Equal("abcdef", result)

	result = Abbreviate("abcdef", 0)
	assertions.Equal("abcdef", result)

	result = Abbreviate("abcdef", 1)
	assertions.Equal("abcdef", result)

	result = Abbreviate("abcd", 3)
	assertions.Equal("abcd", result)
}

func TestContainsOnly(t *testing.T) {
	assertions := assert.New(t)

	result := ContainsOnly("a", "a")
	assertions.True(result)

	result = ContainsOnly("a", "ab")
	assertions.True(result)

	result = ContainsOnly("a", "")
	assertions.False(result)

	result = ContainsOnly("c", "a")
	assertions.False(result)

	result = ContainsOnly("ac", "a")
	assertions.False(result)
}

func TestContainsNone(t *testing.T) {
	assertions := assert.New(t)

	result := ContainsNone("a", "a")
	assertions.False(result)

	result = ContainsNone("a", "ab")
	assertions.False(result)

	result = ContainsNone("a", "")
	assertions.True(result)

	result = ContainsNone("c", "a")
	assertions.True(result)

	result = ContainsNone("ac", "a")
	assertions.False(result)
}
