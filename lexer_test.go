package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenizeCaseSensitive(t *testing.T) {
	assertions := assert.New(t)

	var tokenTypes []TokenType
	tokenTypes = append(tokenTypes, TokenType{
		Type:  "keyword",
		Words: []string{"if", "for"},
	})
	result := Tokenize("if", tokenTypes)
	assertions.Equal("if", result[0].Text)
	assertions.Equal(0, result[0].Position)
	assertions.Equal("keyword", result[0].Type)
}

func TestTokenizeCaseInsensitive(t *testing.T) {
	assertions := assert.New(t)

	var tokenTypes []TokenType
	tokenTypes = append(tokenTypes, TokenType{
		Type:  "keyword",
		Words: []string{"if", "for"},
	})
	result := Tokenize("IF", tokenTypes)
	assertions.Equal("IF", result[0].Text)
	assertions.Equal(0, result[0].Position)
	assertions.Equal("keyword", result[0].Type)
}

func TestTokenizeWord(t *testing.T) {
	assertions := assert.New(t)

	var tokenTypes []TokenType
	tokenTypes = append(tokenTypes, TokenType{
		Type:  "keyword",
		Words: []string{"if", "for"},
	})
	result := TokenizeWord("if", 0, tokenTypes)
	assertions.Equal("if", result.Text)
	assertions.Equal(0, result.Position)
	assertions.Equal("keyword", result.Type)
}

func TestLookupType(t *testing.T) {
	assertions := assert.New(t)

	var tokenTypes []TokenType
	tokenTypes = append(tokenTypes, TokenType{
		Type:  "keyword",
		Words: []string{"if", "for"},
	})

	result := LookupType("for", tokenTypes)
	assertions.Equal("keyword", result.Type)
	assertions.Equal([]string{"if", "for"}, result.Words)
}

func TestLookupTypeNotFound(t *testing.T) {
	assertions := assert.New(t)

	var tokenTypes []TokenType
	tokenTypes = append(tokenTypes, TokenType{
		Type:  "keyword",
		Words: []string{"if", "for"},
	})

	result := LookupType("var", tokenTypes)
	assertions.Equal("", result.Type)
	assertions.Equal([]string{""}, result.Words)
}
