/*
File    : go-mix/lexer/lexer.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// Lexer performs lexical analysis (tokenization) of GoMix source code.
// It scans through the source text character by character, identifying and
// creating tokens that represent the syntactic elements of the language.
//
// The lexer maintains state about its current position in the source code,
// including line and column numbers for error reporting. It handles:
//   - Operators (arithmetic, logical, bitwise, comparison)
//   - Keywords (if, else, func, return, etc.)
//   - Literals (numbers, strings, booleans)
//   - Identifiers (variable and function names)
//   - Structural symbols (parentheses, braces, brackets)
//   - Comments (single-line // and multi-line /* */)
//   - Whitespace (which is ignored)
//
// Fields:
//   - Src: The complete source code as a string
//   - Current: The byte at the current position being examined
//   - Position: The current index in the source string (0-indexed)
//   - SrcLength: The total length of the source string
//   - Line: The current line number in the source (1-indexed)
//   - Column: The current column number in the source (1-indexed)
type Lexer struct {
	Src       string // Entire source code in plain text format
	Current   byte   // Current character being examined
	Position  int    // Current position of pointer in the source code
	SrcLength int    // Length of source string
	Line      int    // Line number in source (1-indexed)
	Column    int    // Column number in source (1-indexed)
}

// NewLexer creates and initializes a new Lexer for the given source code.
// It sets up the initial state with the first character of the source
// and initializes position tracking to line 1, column 1.
//
// Parameters:
//   - src: The source code string to tokenize
//
// Returns:
//   - Lexer: A new lexer ready to tokenize the source code
//
// Example:
//
//	lexer := NewLexer("var x = 42")
func NewLexer(src string) Lexer {
	// Initialize current to null byte if source is empty
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

// NextToken retrieves the next token from the source code stream.
// It skips whitespace and comments, then identifies and returns the next
// meaningful token. This is the main entry point for token-by-token parsing.
//
// The method handles:
//   - All operators (arithmetic, logical, bitwise, comparison, assignment)
//   - Compound assignment operators (+=, -=, *=, etc.)
//   - Structural symbols (parentheses, braces, brackets)
//   - Delimiters (comma, semicolon, colon)
//   - String literals (with escape sequence support)
//   - Numeric literals (integers and floats)
//   - Identifiers and keywords
//
// Returns:
//   - Token: The next token in the source, or EOF_TYPE if end is reached
//
// Example:
//
//	token := lexer.NextToken()  // Returns first token
//	token = lexer.NextToken()   // Returns second token, etc.
func (lex *Lexer) NextToken() Token {

	var token Token
	// Skip any whitespace and comments before the next token
	lex.ignoreWhitespacesAndComments()

	// Match the current character to determine token type
	switch lex.Current {
	case '=':
		// Could be '=' (assignment) or '==' (equality)
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(EQ_OP, "==", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(ASSIGN_OP, "=", lex.Line, lex.Column)
		}
	case '!':
		// Could be '!' (logical NOT) or '!=' (not equal)
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(NE_OP, "!=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(NOT_OP, "!", lex.Line, lex.Column)
		}
	case '<':
		// Could be '<', '<=', '<<', or '<<='
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
		// Could be '>', '>=', '>>', or '>>='
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
		// Could be '+' or '+='
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(PLUS_ASSIGN, "+=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(PLUS_OP, "+", lex.Line, lex.Column)
		}
	case '-':
		// Could be '-' or '-='
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(MINUS_ASSIGN, "-=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MINUS_OP, "-", lex.Line, lex.Column)
		}
	case '*':
		// Could be '*' or '*='
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(MUL_ASSIGN, "*=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MUL_OP, "*", lex.Line, lex.Column)
		}
	case '/':
		// Could be '/' or '/='
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(DIV_ASSIGN, "/=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(DIV_OP, "/", lex.Line, lex.Column)
		}
	case '%':
		// Could be '%' or '%='
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(MOD_ASSIGN, "%=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MOD_OP, "%", lex.Line, lex.Column)
		}
	case '^':
		// Could be '^' (XOR) or '^='
		if lex.peek() == '=' {
			lex.advance()
			token = NewTokenWithMetadata(BIT_XOR_ASSIGN, "^=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(BIT_XOR_OP, "^", lex.Line, lex.Column)
		}
	case '~':
		// Bitwise NOT operator
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
	case '.':
		// Could be '...' (range operator)
		// Need to check if this is part of a number or a range operator
		// If the previous token was a number and current is '.', it's handled in readNumber
		// Here we only handle the case where '.' starts a token
		if lex.peek() == '.' {
			// Check if there's a third dot
			if lex.Position+2 < lex.SrcLength && lex.Src[lex.Position+2] == '.' {
				lex.advance() // consume second dot
				lex.advance() // consume third dot
				token = NewTokenWithMetadata(RANGE_OP, "...", lex.Line, lex.Column)
			} else {
				// Just two dots - invalid, treat as EOF for now
				token = NewTokenWithMetadata(EOF_TYPE, "EOF", lex.Line, lex.Column)
			}
		} else {
			// Single dot - could be member access operator
			token = NewTokenWithMetadata(DOT_OP, ".", lex.Line, lex.Column)
		}
	case '&':
		// Could be '&' (bitwise AND), '&&' (logical AND), or '&='
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
		// Could be '|' (bitwise OR), '||' (logical OR), or '|='
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
		// Null byte indicates end of file
		token = NewTokenWithMetadata(EOF_TYPE, "EOF", lex.Line, lex.Column)
	case '"':
		// String literal - delegate to specialized handler
		return readStringLiteral(lex)
	default:
		// Check for numeric literals, identifiers, or invalid characters
		if isNumeric(lex.Current) {
			return readNumber(lex)
		} else if isAlpha(lex.Current) || (lex.Current == '_') {
			return readIdentifier(lex)
		}

		// any special characters that are not recognized as valid tokens will be treated as EOF for now
		if isSpecial(lex.Current) {
			token = NewTokenWithMetadata(INVALID_TYPE, string(lex.Current), lex.Line, lex.Column)
		} else {
			// Unrecognized character - return EOF
			token = NewTokenWithMetadata(EOF_TYPE, "EOF", lex.Line, lex.Column)
		}

	}

	// Move to the next character for the next token
	lex.advance()

	return token
}

// isSpecial checks if a character is a special symbol that is not valid in GoMix.
// This includes characters that are not part of the defined token set and are not alphanumeric or whitespace.
func isSpecial(c byte) bool {
	return !isAlphanumeric(c) && !isWhitespace(c) && !strings.ContainsRune("=+-*/%&|^~!<>.,;:(){}[]\"", rune(c))
}

// readStringLiteral reads and tokenizes a string literal from the source.
// It handles escape sequences like \n, \t, \\, \", etc.
// String literals must be enclosed in double quotes (").
//
// Supported escape sequences:
//   - \n: newline
//   - \t: tab
//   - \r: carriage return
//   - \f: form feed
//   - \v: vertical tab
//   - \\: backslash
//   - \": double quote
//   - \': single quote
//   - \0: null character
//
// Parameters:
//   - lex: Pointer to the lexer instance
//
// Returns:
//   - Token: A STRING_LIT token with the string content, or INVALID_TYPE on error
//
// Example:
//
//	Source: "hello\nworld"
//	Returns: Token{Type: STRING_LIT, Literal: "hello\nworld"}
func readStringLiteral(lex *Lexer) Token {
	if lex.Current != '"' {
		// Error: expected opening quote
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed string literal", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}
	lex.advance() // Consume opening quote

	var builder strings.Builder

	// Read characters until closing quote
	for lex.Current != '"' {
		// Check for unterminated string (EOF or actual newline in source)
		if lex.Current == 0 {
			fmt.Errorf("[%d:%d] LEXER ERROR: string literal not terminated - unexpected EOF", lex.Line, lex.Column)
			return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
		}

		// Handle escape sequences
		if lex.Current == '\\' {
			lex.advance() // Consume the backslash
			escapedChar, valid := escapeChar(lex.Current)
			if !valid {
				fmt.Errorf("[%d:%d] LEXER ERROR: invalid escape sequence: \\%c", lex.Line, lex.Column, lex.Current)
				return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
			}
			builder.WriteByte(escapedChar)
			lex.advance()
			continue
		}

		// Regular character - add to string
		builder.WriteByte(lex.Current)
		lex.advance()
	}

	lex.advance() // Consume closing quote
	return NewTokenWithMetadata(STRING_LIT, builder.String(), lex.Line, lex.Column)
}

// escapeChar converts an escape sequence character to its actual byte value.
// This is used when processing escape sequences in string literals.
//
// Parameters:
//   - c: The character following the backslash in an escape sequence
//
// Returns:
//   - byte: The actual byte value of the escape sequence
//   - bool: true if the escape sequence is valid, false otherwise
//
// Example:
//
//	escapeChar('n') -> ('\n', true)
//	escapeChar('x') -> (0, false)
func escapeChar(c byte) (byte, bool) {
	switch c {
	case 'n':
		return '\n', true // Newline
	case 't':
		return '\t', true // Tab
	case 'r':
		return '\r', true // Carriage return
	case 'f':
		return '\f', true // Form feed
	case 'v':
		return '\v', true // Vertical tab
	case '\\':
		return '\\', true // Backslash
	case '"':
		return '"', true // Double quote
	case '\'':
		return '\'', true // Single quote
	case '0':
		return 0, true // Null character
	default:
		return 0, false // Invalid escape sequence
	}
}

// readNumber reads and tokenizes a numeric literal from the source.
// It supports both integer and floating-point numbers.
//
// Currently supported formats:
//   - Integers: 0, 10, 123, etc.
//   - Floats: 10.5, 0.123, 123.456, etc.
//
// TODO: Add support for additional formats:
//   - Negative numbers: -10, -10.5
//   - Scientific notation: 1e9, 1.4e9, -12E-2
//   - Hexadecimal: 0x16
//   - Octal: 0777
//
// Parameters:
//   - lex: Pointer to the lexer instance
//
// Returns:
//   - Token: An INT_LIT or FLOAT_LIT token, or INVALID_TYPE on error
//
// Example:
//
//	Source: "123.45"
//	Returns: Token{Type: FLOAT_LIT, Literal: "123.45"}
func readNumber(lex *Lexer) Token {
	position := lex.Position
	hasDot := false

	// Ensure we start with a digit
	if isNumeric(lex.Current) {
		lex.advance()
	} else {
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed number literal", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}

	// Continue reading digits and at most one decimal point
	for isNumeric(lex.Current) || lex.Current == '.' {
		if lex.Current == '.' {
			// Check if this is the start of a range operator (...)
			if lex.peek() == '.' {
				// This is a range operator, stop reading the number
				break
			}
			if hasDot {
				// Second dot encountered - stop here
				break
			}
			hasDot = true
		}
		lex.advance()
	}

	// Determine token type based on presence of decimal point
	tokenType := INT_LIT
	if hasDot {
		tokenType = FLOAT_LIT
	}

	return NewTokenWithMetadata(tokenType, lex.Src[position:lex.Position], lex.Line, lex.Column)
}

// readIdentifier reads and tokenizes an identifier or keyword from the source.
// Identifiers can be variable names, function names, or language keywords.
//
// Rules:
//   - Must start with a letter (a-z, A-Z) or underscore (_)
//   - Can contain letters, digits, or underscores
//   - Keywords are identified using the lookupIdent function
//
// Parameters:
//   - lex: Pointer to the lexer instance
//
// Returns:
//   - Token: An IDENTIFIER_ID token or a keyword token type, or INVALID_TYPE on error
//
// Example:
//
//	Source: "myVariable"
//	Returns: Token{Type: IDENTIFIER_ID, Literal: "myVariable"}
//
//	Source: "if"
//	Returns: Token{Type: IF_KEY, Literal: "if"}
func readIdentifier(lex *Lexer) Token {
	position := lex.Position

	// Ensure we start with a letter or underscore
	if isAlpha(lex.Current) || lex.Current == '_' {
		lex.advance()
	} else {
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed identifier", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}

	// Continue reading alphanumeric characters and underscores
	for isAlphanumeric(lex.Current) || lex.Current == '_' {
		lex.advance()
	}

	literal := lex.Src[position:lex.Position]

	// Check if this identifier is actually a keyword
	return NewTokenWithMetadata(lookupIdent(literal), literal, lex.Line, lex.Column)
}

// peek looks ahead to the next character in the source without consuming it.
// This is useful for lookahead when determining multi-character tokens.
//
// Returns:
//   - byte: The next character, or 0 if at end of source
//
// Example:
//
//	If current is '=' and next is '=', peek() returns '='
func (lex *Lexer) peek() byte {
	if lex.Position+1 >= lex.SrcLength {
		return 0 // End of source
	}
	return lex.Src[lex.Position+1]
}

// advance moves the lexer to the next character in the source.
// It updates the Current byte, Position, and Column tracking.
// When a newline is encountered elsewhere, Line is incremented separately.
//
// After calling advance:
//   - Position is incremented
//   - Column is incremented
//   - Current is set to the new character (or 0 if at end)
func (lex *Lexer) advance() {
	lex.Position++
	lex.Column++

	if lex.Position >= lex.SrcLength {
		lex.Current = 0              // Null byte indicates end
		lex.Position = lex.SrcLength // Keep position at end
	} else {
		lex.Current = lex.Src[lex.Position]
	}
}

// ignoreWhitespacesAndComments skips over whitespace and comments in the source.
// This method is called before tokenizing each meaningful token.
//
// It handles:
//   - Whitespace characters (space, tab, newline, etc.)
//   - Single-line comments (// ...)
//   - Multi-line comments (/* ... */)
//
// When a newline is encountered, the Line counter is incremented and
// Column is reset to 1.
func (lex *Lexer) ignoreWhitespacesAndComments() {
	for {
		if isWhitespace(lex.Current) {
			// Track line numbers when encountering newlines
			if lex.Current == '\n' {
				lex.Line++
				lex.Column = 1
			}
			lex.advance()
		} else if lex.Current == '/' && lex.peek() == '/' {
			// Single-line comment detected
			lex.skipSingleLineComment()
		} else if lex.Current == '/' && lex.peek() == '*' {
			// Multi-line comment detected
			lex.skipMultiLineComment()
		} else {
			// No more whitespace or comments
			break
		}
	}
}

// skipSingleLineComment skips over a single-line comment (// ...).
// It advances the lexer until a newline or end of file is reached.
// The newline itself is not consumed, allowing line tracking to work correctly.
//
// Example:
//
//	Source: "// this is a comment\nvar x"
//	After skip: lexer is positioned at '\n'
func (lex *Lexer) skipSingleLineComment() {
	// Skip the '//' characters
	lex.advance()
	lex.advance()

	// Skip until end of line or end of file
	for lex.Current != '\n' && lex.Current != 0 {
		lex.advance()
	}
}

// skipMultiLineComment skips over a multi-line comment (/* ... */).
// It advances the lexer until the closing */ is found or end of file is reached.
// Line numbers are tracked correctly even within multi-line comments.
//
// Example:
//
//	Source: "/* comment\nspanning lines */var x"
//	After skip: lexer is positioned after '*/'
func (lex *Lexer) skipMultiLineComment() {
	// Skip the '/*' characters
	lex.advance()
	lex.advance()

	// Skip until we find '*/' or reach end of file
	for lex.Current != 0 {
		if lex.Current == '*' && lex.peek() == '/' {
			// Found closing '*/' - skip it and exit
			lex.advance()
			lex.advance()
			break
		}
		lex.advance()
	}
}

// isWhitespace checks if the given byte is a whitespace character.
// Uses Unicode's definition of whitespace, which includes:
//   - Space, tab, newline, carriage return, form feed, vertical tab
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is whitespace, false otherwise
func isWhitespace(curr byte) bool {
	return unicode.IsSpace(rune(curr))
}

// isAlphanumeric checks if the given byte is an alphanumeric character.
// This includes both letters (a-z, A-Z) and digits (0-9).
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is a letter or digit, false otherwise
func isAlphanumeric(curr byte) bool {
	return unicode.IsLetter(rune(curr)) || unicode.IsDigit(rune(curr))
}

// isNumeric checks if the given byte is a numeric digit (0-9).
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is a digit, false otherwise
func isNumeric(curr byte) bool {
	return unicode.IsDigit(rune(curr))
}

// isAlpha checks if the given byte is an alphabetic character (a-z, A-Z).
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is a letter, false otherwise
func isAlpha(curr byte) bool {
	return unicode.IsLetter(rune(curr))
}

// isEscape checks if the given byte is an escape character.
// This function is kept for backward compatibility but is no longer used
// in readStringLiteral as escape sequences are now handled properly
// through the escapeChar function.
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is an escape character, false otherwise
func isEscape(curr byte) bool {
	return curr == '\\' || curr == '\'' || curr == 0 ||
		curr == '\n' || curr == '\t' || curr == '\r' || curr == '\f' || curr == '\v'
}

// ConsumeTokens tokenizes the entire source code and returns all tokens.
// It repeatedly calls NextToken until EOF is reached, collecting all tokens
// into a slice. This is useful for batch processing or debugging.
//
// Returns:
//   - []Token: A slice containing all tokens from the source (excluding EOF)
//
// Example:
//
//	lexer := NewLexer("var x = 42")
//	tokens := lexer.ConsumeTokens()
//	// tokens contains: [VAR_KEY, IDENTIFIER_ID, ASSIGN_OP, INT_LIT]
func (lex *Lexer) ConsumeTokens() []Token {
	tokens := make([]Token, 0)
	for {
		token := lex.NextToken()
		// Uncomment for debugging: token.Print()
		if token.Type == EOF_TYPE {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}
