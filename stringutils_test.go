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

func TestEqualsAny(t *testing.T) {
	assertions := assert.New(t)

	result := EqualsAny("India", []string{"India", "Canada"})
	assertions.True(result)

	result = EqualsAny("Brazil", []string{"India", "Canada"})
	assertions.False(result)

	result = EqualsAny("In", []string{"India", "Canada"})
	assertions.False(result)

	result = EqualsAny("Brazil", []string{""})
	assertions.False(result)
}

func TestEndsWithAny(t *testing.T) {
	assertions := assert.New(t)

	result := EndsWithAny("abc.exe", []string{"exe", "txt"})
	assertions.True(result)

	result = EndsWithAny("abc.exe", []string{"cpp", "txt"})
	assertions.False(result)

	result = EndsWithAny("", []string{"txt"})
	assertions.False(result)
}

func TestStartsWithAny(t *testing.T) {
	assertions := assert.New(t)

	result := StartsWithAny("abc.exe", []string{"abc", "txt"})
	assertions.True(result)

	result = StartsWithAny("abc.exe", []string{"b", "c"})
	assertions.False(result)

	result = StartsWithAny("", []string{"txt"})
	assertions.False(result)
}

func TestEqualsIgnoreCase(t *testing.T) {
	assertions := assert.New(t)

	result := EqualsIgnoreCase("abc.exe", "ABC.exe")
	assertions.True(result)

	result = EqualsIgnoreCase("abc.exe", "abc.exe")
	assertions.True(result)

	result = EqualsIgnoreCase("ABC.EXE", "abc.exe")
	assertions.True(result)

	result = EqualsIgnoreCase("XYZ", "abc")
	assertions.False(result)
}

func TestEqualsAnyIgnoreCase(t *testing.T) {
	assertions := assert.New(t)

	result := EqualsAnyIgnoreCase("abc.exe", []string{"ABC.exe"})
	assertions.True(result)

	result = EqualsAnyIgnoreCase("abc.exe", []string{"abc.exe"})
	assertions.True(result)

	result = EqualsAnyIgnoreCase("ABC.EXE", []string{"abc.exe"})
	assertions.True(result)

	result = EqualsAnyIgnoreCase("XYZ", []string{"abc"})
	assertions.False(result)
}
