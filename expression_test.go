package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindExpressionInTokens(t *testing.T) {
	assertions := assert.New(t)

	token := Token{
		Type:     "keyword",
		Position: 1,
		Text:     "if",
	}

	expression := Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.True(FindExpressionInTokens([]Token{token}, expression))

	token = Token{
		Type:     "bracket",
		Position: 1,
		Text:     "(",
	}

	expression = Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.False(FindExpressionInTokens([]Token{token}, expression))
}

func TestFindExpressionsInTokens(t *testing.T) {
	assertions := assert.New(t)

	token := Token{
		Type:     "keyword",
		Position: 1,
		Text:     "if",
	}

	expression := Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.True(FindExpressionsInTokens([]Token{token}, []Expression{expression}))

	token2 := Token{
		Type:     "bracket",
		Position: 2,
		Text:     "(",
	}

	expression2 := Expression{
		Type: "bracket",
		Text: "(",
	}
	assertions.True(FindExpressionsInTokens([]Token{token, token2}, []Expression{expression, expression2}))

	expression2 = Expression{
		Type: "keyword",
		Text: "for",
	}
	assertions.False(FindExpressionsInTokens([]Token{token, token2}, []Expression{expression, expression2}))

	token = Token{
		Type:     "bracket",
		Position: 1,
		Text:     "(",
	}

	expression = Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.False(FindExpressionsInTokens([]Token{token}, []Expression{expression}))

	token2 = Token{
		Type:     "bracket",
		Position: 2,
		Text:     ")",
	}

	assertions.False(FindExpressionsInTokens([]Token{token, token2}, []Expression{expression}))

	assertions.False(FindExpressionsInTokens([]Token{token}, []Expression{expression, expression2}))

	assertions.False(FindExpressionsInTokens([]Token{token}, []Expression{}))
}

func TestMatchToken(t *testing.T) {
	assertions := assert.New(t)

	token := Token{
		Type:     "keyword",
		Position: 1,
		Text:     "if",
	}

	expression := Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.True(MatchToken(token, expression))

	expression = Expression{
		Type: "keyword",
		Text: "IF",
	}
	assertions.True(MatchToken(token, expression))

	expression = Expression{
		Type: "keyword",
		Text: "",
	}
	assertions.True(MatchToken(token, expression))
}
