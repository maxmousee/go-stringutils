package stringutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsExpressionInTokens(t *testing.T) {
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
	assertions.True(ContainsExpressionInTokens(expression, []Token{token}))

	token = Token{
		Type:     "bracket",
		Position: 1,
		Text:     "(",
	}

	expression = Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.False(ContainsExpressionInTokens(expression, []Token{token}))
}

func TestContainsExpressionsInTokens(t *testing.T) {
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
	assertions.True(ContainsExpressionsInTokens([]Expression{expression}, []Token{token}))

	token2 := Token{
		Type:     "bracket",
		Position: 2,
		Text:     "(",
	}

	expression2 := Expression{
		Type: "bracket",
		Text: "(",
	}
	assertions.True(ContainsExpressionsInTokens([]Expression{expression, expression2}, []Token{token, token2}))

	expression2 = Expression{
		Type: "keyword",
		Text: "for",
	}
	assertions.False(ContainsExpressionsInTokens([]Expression{expression, expression2}, []Token{token, token2}))

	token = Token{
		Type:     "bracket",
		Position: 1,
		Text:     "(",
	}

	expression = Expression{
		Type: "keyword",
		Text: "if",
	}
	assertions.False(ContainsExpressionsInTokens([]Expression{expression}, []Token{token}))

	token2 = Token{
		Type:     "bracket",
		Position: 2,
		Text:     ")",
	}

	assertions.False(ContainsExpressionsInTokens([]Expression{expression}, []Token{token, token2}))

	assertions.False(ContainsExpressionsInTokens([]Expression{expression, expression2}, []Token{token}))

	assertions.False(ContainsExpressionsInTokens([]Expression{}, []Token{token}))
}

func TestContainsComplexExpressionsInTokens(t *testing.T) {
	assertions := assert.New(t)
	expression := Expression{
		Type: "bracket",
		Text: "(",
	}

	expression2 := Expression{
		Type: "bracket",
		Text: ")",
	}

	token := Token{
		Type:     "bracket",
		Position: 1,
		Text:     "{",
	}
	token2 := Token{
		Type:     "bracket",
		Position: 2,
		Text:     "(",
	}
	token3 := Token{
		Type:     "bracket",
		Position: 3,
		Text:     ")",
	}
	token4 := Token{
		Type:     "bracket",
		Position: 4,
		Text:     "{",
	}

	assertions.True(ContainsExpressionsInTokens([]Expression{expression, expression2}, []Token{token, token2, token3, token4}))
}

func TestContainsComplexExpressionsInTokensOnSecondTime(t *testing.T) {
	assertions := assert.New(t)
	expression := Expression{
		Type: "bracket",
		Text: "(",
	}

	expression2 := Expression{
		Type: "bracket",
		Text: ")",
	}

	token := Token{
		Type:     "bracket",
		Position: 1,
		Text:     "(",
	}
	token2 := Token{
		Type:     "bracket",
		Position: 2,
		Text:     "(",
	}
	token3 := Token{
		Type:     "bracket",
		Position: 3,
		Text:     ")",
	}
	token4 := Token{
		Type:     "bracket",
		Position: 4,
		Text:     "{",
	}

	assertions.True(ContainsExpressionsInTokens([]Expression{expression, expression2}, []Token{token, token2, token3, token4}))
}

func TestNotContainsComplexExpressionsInTokens(t *testing.T) {
	assertions := assert.New(t)
	expression := Expression{
		Type: "bracket",
		Text: "{",
	}

	expression2 := Expression{
		Type: "keyword",
		Text: "var",
	}

	token := Token{
		Type:     "bracket",
		Position: 1,
		Text:     "{",
	}
	token2 := Token{
		Type:     "bracket",
		Position: 2,
		Text:     "(",
	}
	token3 := Token{
		Type:     "bracket",
		Position: 3,
		Text:     ")",
	}
	token4 := Token{
		Type:     "bracket",
		Position: 4,
		Text:     "{",
	}

	assertions.False(ContainsExpressionsInTokens([]Expression{expression, expression2}, []Token{token, token2, token3, token4}))
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
