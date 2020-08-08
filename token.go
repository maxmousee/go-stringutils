package stringutils

type Token struct {
	Type   				string          // verb, keyword, separator, etc
	Position	    	int             // position of the word in string
	Text				string			// text
}

type TokenType struct {
	Type			string 				// verb, keyword, separator, etc
	Words 			[]string			// words to be considered as a given token type
	CaseSensitive	bool				// whether or not this token type is case sensitive
}
