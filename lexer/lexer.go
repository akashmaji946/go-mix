package lexer

import "unicode"

type Lexer struct {
	// entire source code in plain text format
	Src string
	// current character in the source code
	Current byte
	// current position of pointer in the source code
	Position int
	// length of source string
	SrcLength int
}

// NewLexer(): constructor for Lexer
func NewLexer(src string) Lexer {
	return Lexer{
		Src:       src,
		Current:   src[0],
		Position:  0,
		SrcLength: len(src),
	}
}

// NextToken(): gets next token from the stream
// we will skip whitespaces: '\n', '\t', '\r', '\f', '\v', ' '
func (lex *Lexer) NextToken() Token {

	var token Token
	lex.ignoreWhitespaces()

	switch lex.Current {
	case '=':
		if lex.peek() == '=' {
			lex.advance()
			token = NewToken(EQ_OP, "==")
		} else {
			token = NewToken(ASSIGN_OP, "=")
		}
	case '!':
		if lex.peek() == '=' {
			lex.advance()
			token = NewToken(NE_OP, "!=")
		} else {
			token = NewToken(NOT_OP, "!")
		}
	case '<':
		if lex.peek() == '=' {
			lex.advance()
			token = NewToken(LE_OP, "<=")
		} else if lex.peek() == '<' {
			lex.advance()
			token = NewToken(BIT_LEFT_OP, "<<")
		} else {
			token = NewToken(LT_OP, "<")
		}
	case '>':
		if lex.peek() == '=' {
			lex.advance()
			token = NewToken(GE_OP, ">=")
		} else if lex.peek() == '>' {
			lex.advance()
			token = NewToken(BIT_RIGHT_OP, ">>")
		} else {
			token = NewToken(GT_OP, ">")
		}
	case '+':
		token = NewToken(PLUS_OP, "+")
	case '-':
		token = NewToken(MINUS_OP, "-")
	case '*':
		token = NewToken(MUL_OP, "*")
	case '/':
		token = NewToken(DIV_OP, "/")
	case '%':
		token = NewToken(MOD_OP, "%")
	case '^':
		token = NewToken(BIT_XOR_OP, "^")
	case '~':
		token = NewToken(BIT_NOT_OP, "~")
	case '(':
		token = NewToken(LEFT_PAREN, "(")
	case ')':
		token = NewToken(RIGHT_PAREN, ")")
	case '{':
		token = NewToken(LEFT_BRACE, "{")
	case '}':
		token = NewToken(RIGHT_BRACE, "}")
	case '[':
		token = NewToken(LEFT_BRACKET, "[")
	case ']':
		token = NewToken(RIGHT_BRACKET, "]")
	case ',':
		token = NewToken(COMMA_DELIM, ",")
	case ';':
		token = NewToken(SEMICOLON_DELIM, ";")
	case ':':
		token = NewToken(COLON_DELIM, ":")
	case '&':
		if lex.peek() == '&' {
			lex.advance()
			token = NewToken(AND_OP, "&&")
		} else {
			token = NewToken(BIT_AND_OP, "&")
		}
	case '|':
		if lex.peek() == '|' {
			lex.advance()
			token = NewToken(OR_OP, "||")
		} else {
			token = NewToken(BIT_OR_OP, "|")
		}
	case 0:
		token = NewToken(EOF_TYPE, "EOF")
	// match string literals
	case '"':
		return readStringLiteral(lex)
	default:
		if isNumeric(lex.Current) {
			return readNumber(lex)
		} else if isAlpha(lex.Current) || (lex.Current == '_') {
			return readIdentifier(lex)
		}
		token = NewToken(EOF_TYPE, "EOF")

	}
	lex.advance()

	return token
}

// readStringLiteral(): reads a string literal in the source code
// TODO: add support for escape sequences
// e.g. "\n", "\t", "\r", "\f", "\v", "\\", "\"", "\'"
func readStringLiteral(lex *Lexer) Token {
	position := lex.Position
	if lex.Current != '"' {
		// TODO: do better error handling
		panic("[ERROR] Malformed string literal")
	}
	lex.advance()
	for lex.Current != '"' {
		if isEscape(lex.Current) {
			// TODO: do better error handling
			panic("[ERROR] String literal not terminated")
		}
		lex.advance()
	}
	lex.advance()
	return NewToken(STRING_LIT, lex.Src[position+1:lex.Position-1])
}

// readNumber(): reads a number in the source code
// TODO: add support for other formats
// eg. 0, 10, -10, 10.123, 1e9, 1.4e9, -12E-2, -123.123, 0x16, 0777
// readNumber(): reads a number in the source code
// TODO: add support for other formats
// eg. 0, 10, -10, 10.123, 1e9, 1.4e9, -12E-2, -123.123, 0x16, 0777
func readNumber(lex *Lexer) Token {
	position := lex.Position
	hasDot := false

	if isNumeric(lex.Current) {
		lex.advance()
	} else {
		// TODO: do better error handling
		panic("[ERROR] Malformed number literal")
	}

	for isNumeric(lex.Current) || lex.Current == '.' {
		if lex.Current == '.' {
			if hasDot {
				break
			}
			hasDot = true
		}
		lex.advance()
	}

	tokenType := INT_LIT
	if hasDot {
		tokenType = FLOAT_LIT
	}

	return NewToken(tokenType, lex.Src[position:lex.Position])
}

// readIdentifier(): reads an identifier in the source code
// identifier can be a keyword also
func readIdentifier(lex *Lexer) Token {
	position := lex.Position
	if isAlpha(lex.Current) || lex.Current == '_' {
		lex.advance()
	} else {
		// TODO: do better error handling
		panic("[ERROR] Malformed identifier")
	}
	for isAlphanumeric(lex.Current) || lex.Current == '_' {
		lex.advance()
	}

	literal := lex.Src[position:lex.Position]

	// lookup the token type of the identifier (maybe a keyword)
	return NewToken(lookupIdent(literal), literal)
}

// peek(): looks ahead to the next character
func (lex *Lexer) peek() byte {
	if lex.Position+1 >= lex.SrcLength {
		return 0
	}
	return lex.Src[lex.Position+1]
}

// advance(): get the very next character in sequence, null if end reached
func (lex *Lexer) advance() {
	lex.Position++
	if lex.Position >= lex.SrcLength {
		lex.Current = 0              // null byte
		lex.Position = lex.SrcLength // keep at end
	} else {
		lex.Current = lex.Src[lex.Position]
	}

}

// ignoreWhitespaces(): ignores all whitespaces in the source code
func (lex *Lexer) ignoreWhitespaces() {
	for isWhitespace(lex.Current) {
		lex.advance()
	}
}

// isWhitespace(): check if current byte is a whitespace or not
func isWhitespace(curr byte) bool {
	return unicode.IsSpace(rune(curr))
}

// isAlphanumeric(): check if current byte is an alphanumeric character or not
func isAlphanumeric(curr byte) bool {
	return unicode.IsLetter(rune(curr)) || unicode.IsDigit(rune(curr))
}

// isNumeric(): check if current byte is a numeric character or not
func isNumeric(curr byte) bool {
	return unicode.IsDigit(rune(curr))
}

// isAlpha(): check if current byte is an alphabetic character or not
func isAlpha(curr byte) bool {
	return unicode.IsLetter(rune(curr))
}

// isEscape(): check if current byte is an escape character or not
func isEscape(curr byte) bool {
	return curr == '\\' || curr == '\'' || curr == 0 ||
		curr == '\n' || curr == '\t' || curr == '\r' || curr == '\f' || curr == '\v'
}

// ConsumeTokens(): get all tokens from src
// stop when EOF is reached
func (lex *Lexer) ConsumeTokens() []Token {
	tokens := make([]Token, 0)
	for {
		token := lex.NextToken()
		// token.Print()
		if token.Type == EOF_TYPE {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens

}
