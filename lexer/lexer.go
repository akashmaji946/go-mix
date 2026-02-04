package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

type Lexer struct {
	// entire source code in plain text format
	Src string
	// current character in the source code
	Current byte
	// current position of pointer in the source code
	Position int
	// length of source string
	SrcLength int
	// line number in source
	Line int
	// column number in source
	Column int
}

// NewLexer(): constructor for Lexer
func NewLexer(src string) Lexer {
	current := byte(0)
	if len(src) > 0 {
		current = src[0]
	}
	return Lexer{
		Src:       src,
		Current:   current,
		Position:  0,
		SrcLength: len(src),
		Line:      1,
		Column:    1,
	}
}

// NextToken(): gets next token from the stream
// we will skip whitespaces: '\n', '\t', '\r', '\f', '\v', ' '
func (lex *Lexer) NextToken() Token {

	var token Token
	lex.ignoreWhitespacesAndComments()

	switch lex.Current {
	case '=':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(EQ_OP, "==", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(ASSIGN_OP, "=", lex.Line, lex.Column)
		}
	case '!':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(NE_OP, "!=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(NOT_OP, "!", lex.Line, lex.Column)
		}
	case '<':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(LE_OP, "<=", lex.Line, lex.Column)
		} else if lex.peek() == '<' {
			lex.advance()
			if lex.peek() == '=' {
				lex.advance()
				token = NewTokenWithMetadata(BIT_LEFT_ASSIGN, "<<=", lex.Line, lex.Column)
			} else {
				token = NewTokenWithMetadata(BIT_LEFT_OP, "<<", lex.Line, lex.Column)
			}
		} else {
			token = NewTokenWithMetadata(LT_OP, "<", lex.Line, lex.Column)
		}
	case '>':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(GE_OP, ">=", lex.Line, lex.Column)
		} else if lex.peek() == '>' {
			lex.advance()
			if lex.peek() == '=' {
				lex.advance()
				token = NewTokenWithMetadata(BIT_RIGHT_ASSIGN, ">>=", lex.Line, lex.Column)
			} else {
				token = NewTokenWithMetadata(BIT_RIGHT_OP, ">>", lex.Line, lex.Column)
			}
		} else {
			token = NewTokenWithMetadata(GT_OP, ">", lex.Line, lex.Column)
		}
	case '+':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(PLUS_ASSIGN, "+=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(PLUS_OP, "+", lex.Line, lex.Column)
		}
	case '-':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(MINUS_ASSIGN, "-=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MINUS_OP, "-", lex.Line, lex.Column)
		}
	case '*':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(MUL_ASSIGN, "*=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MUL_OP, "*", lex.Line, lex.Column)
		}
	case '/':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(DIV_ASSIGN, "/=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(DIV_OP, "/", lex.Line, lex.Column)
		}
	case '%':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(MOD_ASSIGN, "%=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MOD_OP, "%", lex.Line, lex.Column)
		}
	case '^':
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(BIT_XOR_ASSIGN, "^=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(BIT_XOR_OP, "^", lex.Line, lex.Column)
		}
	case '~':
		token = NewTokenWithMetadata(BIT_NOT_OP, "~", lex.Line, lex.Column)
	case '(':
		token = NewTokenWithMetadata(LEFT_PAREN, "(", lex.Line, lex.Column)
	case ')':
		token = NewTokenWithMetadata(RIGHT_PAREN, ")", lex.Line, lex.Column)
	case '{':
		token = NewTokenWithMetadata(LEFT_BRACE, "{", lex.Line, lex.Column)
	case '}':
		token = NewTokenWithMetadata(RIGHT_BRACE, "}", lex.Line, lex.Column)
	case '[':
		token = NewTokenWithMetadata(LEFT_BRACKET, "[", lex.Line, lex.Column)
	case ']':
		token = NewTokenWithMetadata(RIGHT_BRACKET, "]", lex.Line, lex.Column)
	case ',':
		token = NewTokenWithMetadata(COMMA_DELIM, ",", lex.Line, lex.Column)
	case ';':
		token = NewTokenWithMetadata(SEMICOLON_DELIM, ";", lex.Line, lex.Column)
	case ':':
		token = NewTokenWithMetadata(COLON_DELIM, ":", lex.Line, lex.Column)
	case '&':
		if lex.peek() == '&' {
			lex.advance()
			token = NewTokenWithMetadata(AND_OP, "&&", lex.Line, lex.Column)
		} else if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(BIT_AND_ASSIGN, "&=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(BIT_AND_OP, "&", lex.Line, lex.Column)
		}
	case '|':
		if lex.peek() == '|' {
			lex.advance()
			token = NewTokenWithMetadata(OR_OP, "||", lex.Line, lex.Column)
		} else if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(BIT_OR_ASSIGN, "|=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(BIT_OR_OP, "|", lex.Line, lex.Column)
		}
	case 0:
		token = NewTokenWithMetadata(EOF_TYPE, "EOF", lex.Line, lex.Column)
	// match string literals
	case '"':
		return readStringLiteral(lex)
	default:
		if isNumeric(lex.Current) {
			return readNumber(lex)
		} else if isAlpha(lex.Current) || (lex.Current == '_') {
			return readIdentifier(lex)
		}
		token = NewTokenWithMetadata(EOF_TYPE, "EOF", lex.Line, lex.Column)

	}
	lex.advance()

	return token
}

// readStringLiteral(): reads a string literal in the source code
// Supports escape sequences: \n, \t, \r, \f, \v, \\, \", \', \0
func readStringLiteral(lex *Lexer) Token {
	if lex.Current != '"' {
		// TODO: do better error handling
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed string literal", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}
	lex.advance() // consume opening quote

	var builder strings.Builder

	for lex.Current != '"' {
		// Check for unterminated string (EOF or actual newline in source)
		if lex.Current == 0 {
			fmt.Errorf("[%d:%d] LEXER ERROR: string literal not terminated - unexpected EOF", lex.Line, lex.Column)
			return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
		}

		// Handle escape sequences
		if lex.Current == '\\' {
			lex.advance() // consume the backslash
			escapedChar, valid := escapeChar(lex.Current)
			if !valid {
				fmt.Errorf("[%d:%d] LEXER ERROR: invalid escape sequence: \\%c", lex.Line, lex.Column, lex.Current)
				return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
			}
			builder.WriteByte(escapedChar)
			lex.advance()
			continue
		}

		// Regular character
		builder.WriteByte(lex.Current)
		lex.advance()
	}

	lex.advance() // consume closing quote
	return NewTokenWithMetadata(STRING_LIT, builder.String(), lex.Line, lex.Column)
}

// escapeChar(): converts an escape sequence character to its actual value
// e.g., 'n' -> '\n', 't' -> '\t', etc.
func escapeChar(c byte) (byte, bool) {
	switch c {
	case 'n':
		return '\n', true
	case 't':
		return '\t', true
	case 'r':
		return '\r', true
	case 'f':
		return '\f', true
	case 'v':
		return '\v', true
	case '\\':
		return '\\', true
	case '"':
		return '"', true
	case '\'':
		return '\'', true
	case '0':
		return 0, true
	default:
		return 0, false
	}
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
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed number literal", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
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

	return NewTokenWithMetadata(tokenType, lex.Src[position:lex.Position], lex.Line, lex.Column)
}

// readIdentifier(): reads an identifier in the source code
// identifier can be a keyword also
func readIdentifier(lex *Lexer) Token {
	position := lex.Position
	if isAlpha(lex.Current) || lex.Current == '_' {
		lex.advance()
	} else {
		// TODO: do better error handling
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed identifier", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}
	for isAlphanumeric(lex.Current) || lex.Current == '_' {
		lex.advance()
	}

	literal := lex.Src[position:lex.Position]

	// lookup the token type of the identifier (maybe a keyword)
	return NewTokenWithMetadata(lookupIdent(literal), literal, lex.Line, lex.Column)
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
	lex.Column++

	if lex.Position >= lex.SrcLength {
		lex.Current = 0              // null byte
		lex.Position = lex.SrcLength // keep at end
	} else {
		lex.Current = lex.Src[lex.Position]
	}

}

// ignoreWhitespacesAndComments(): ignores all whitespaces and comments in the source code
func (lex *Lexer) ignoreWhitespacesAndComments() {
	for {
		if isWhitespace(lex.Current) {
			if lex.Current == '\n' {
				lex.Line++
				lex.Column = 1
			}
			lex.advance()
		} else if lex.Current == '/' && lex.peek() == '/' {
			// Skip single-line comment
			lex.skipSingleLineComment()
		} else if lex.Current == '/' && lex.peek() == '*' {
			// Skip multi-line comment
			lex.skipMultiLineComment()
		} else {
			break
		}
	}
}

// skipSingleLineComment(): skips a single-line comment (// ...)
func (lex *Lexer) skipSingleLineComment() {
	// Skip the // characters
	lex.advance()
	lex.advance()

	// Skip until end of line or end of file
	for lex.Current != '\n' && lex.Current != 0 {
		lex.advance()
	}
}

// skipMultiLineComment(): skips a multi-line comment (/* ... */)
func (lex *Lexer) skipMultiLineComment() {
	// Skip the /* characters
	lex.advance()
	lex.advance()

	// Skip until we find */
	for lex.Current != 0 {
		if lex.Current == '*' && lex.peek() == '/' {
			// Skip the */ characters
			lex.advance()
			lex.advance()
			break
		}
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
// This function is kept for backward compatibility but is no longer used
// in readStringLiteral as escape sequences are now handled properly
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
