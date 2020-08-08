package stringutils

// Token represents a word tokenized into a category of token types
type Token struct {
	Type     string // verb, keyword, separator, etc
	Position int    // position of the word in string
	Text     string // text
}

//TokenType represents a category of words, with a collection of tokens and whether or not they should be matched as
// case sensitive
type TokenType struct {
	Type          string   // verb, keyword, separator, etc
	Words         []string // words to be considered as a given token type
	CaseSensitive bool     // whether or not this token type is case sensitive
}
