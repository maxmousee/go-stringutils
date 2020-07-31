package go_stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	assertions := assert.New(t)

	result := Reverse("")
	assertions.Equal(result, "")

	result = Reverse("gnirts esrever")
	assertions.Equal(result, "reverse string")

	result = Reverse("中文如何？")
	assertions.Equal(result, "？何如文中")

	result = Reverse("中en文混~排怎样？a")
	assertions.Equal(result, "a？样怎排~混文ne中")
}
