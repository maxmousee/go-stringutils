package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenType(t *testing.T) {
	assertions := assert.New(t)

	result := TokenType{
		Type:  "keyword",
		Words: []string{"for", "if"},
	}
	assertions.Equal([]string{"for", "if"}, result.Words)
	assertions.Equal("keyword", result.Type)
}

func TestToken(t *testing.T) {
	assertions := assert.New(t)

	result := Token{
		Type:     "keyword",
		Position: 1,
		Text:     "if",
	}
	assertions.Equal("keyword", result.Type)
	assertions.Equal("if", result.Text)
	assertions.Equal(1, result.Position)
}
