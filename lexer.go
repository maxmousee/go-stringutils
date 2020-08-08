package stringutils

// Tokenize splits a string into an array of strings/words, and categorizes them into tokens
// If the string is empty, it returns an empty slice
func Tokenize(input string, tokenTypes []TokenType) []Token {
	var tokens []Token
	words := WordSplit(input)
	for index, aWord := range words {
		aToken := TokenizeWord(aWord, index, tokenTypes)
		tokens = append(tokens, aToken)
	}
	return tokens
}

// TokenizeWord categorizes a given word into tokens of a given set of token types
// If no match is found, it returns a token with "" as type
func TokenizeWord(word string, position int, tokenTypes []TokenType) Token {
	aTokenType := LookupType(word, tokenTypes)
	return Token{
		Type:     aTokenType.Type,
		Position: position,
		Text:     word,
	}
}

// LookupType looks up a token type for a given word
// If no match is found, it returns a token type with "" as type
func LookupType(word string, tokenTypes []TokenType) TokenType {
	for _, aToken := range tokenTypes {
		if aToken.CaseSensitive {
			if EqualsAny(word, aToken.Words) {
				return aToken
			}
		} else {
			if EqualsAnyIgnoreCase(word, aToken.Words) {
				return aToken
			}
		}
	}
	return TokenType{
		Type:          "",
		Words:         []string{""},
		CaseSensitive: false,
	}
}