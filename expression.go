package stringutils

// Expression represents tokens to be matched in a list of tokens
// they might have a specific word or not, in case the word is empty, any token of a given Type will be matched
type Expression struct {
	Type string // verb, keyword, separator, etc
	Text string // words to be matched with a token (empty for any)
}

// ContainsExpressionsInTokens checks if a given Expression is present in a list of tokens
func ContainsExpressionsInTokens(expressions []Expression, tokens []Token) bool {
	if len(expressions) > len(tokens) {
		return false
	}
	if len(expressions) > 1 {
		return ContainsExpressionsInTokens(expressions[1:], tokens[1:])
	} else if len(expressions) == 1 {
		if len(tokens) > 1 {
			return ContainsExpressionInTokens(expressions[0], tokens[1:])
		}
		return ContainsExpressionInTokens(expressions[0], tokens)
	}
	return false
}

// ContainsExpressionInTokens checks if a given Expression is present in a list of tokens
func ContainsExpressionInTokens(expression Expression, tokens []Token) bool {
	for _, token := range tokens {
		if MatchToken(token, expression) {
			return true
		}
	}
	return false
}

// MatchToken checks if a given Expression matches a given Token
func MatchToken(token Token, expression Expression) bool {
	if len(expression.Text) == 0 {
		return EqualsIgnoreCase(token.Type, expression.Type)
	}
	return EqualsIgnoreCase(token.Type, expression.Type) && EqualsIgnoreCase(token.Text, expression.Text)
}
