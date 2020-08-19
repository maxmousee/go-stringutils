package stringutils

// Expression represents tokens to be matched in a list of tokens
// they might have a specific word or not, in case the word is empty, any token of a given Type will be matched
type Expression struct {
	Type string // verb, keyword, separator, etc
	Text string // words to be matched with a token (empty for any)
}

// FindExpressionsInTokens checks if a given Expression is present in a list of tokens
func FindExpressionsInTokens(tokens []Token, expressions []Expression) bool {
	if len(expressions) > len(tokens) {
		return false
	}
	if len(expressions) > 1 {
		return FindExpressionsInTokens(tokens[1:], expressions[1:])
	} else if len(expressions) == 1 {
		if len(tokens) > 1 {
			return FindExpressionInTokens(tokens[1:], expressions[0])
		} else {
			return FindExpressionInTokens(tokens, expressions[0])
		}
	}
	return false
}

// FindExpressionInTokens checks if a given Expression is present in a list of tokens
func FindExpressionInTokens(tokens []Token, expression Expression) bool {
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
	} else {
		return EqualsIgnoreCase(token.Type, expression.Type) && EqualsIgnoreCase(token.Text, expression.Text)
	}
}
