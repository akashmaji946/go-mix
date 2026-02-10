/*
File    : go-mix/lexer/lexer.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package lexer

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
	lex.IgnoreWhitespacesAndComments()

	// Match the current character to determine token type
	switch lex.Current {
	case '=':
		// Could be '=' (assignment) or '==' (equality)
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(EQ_OP, "==", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(ASSIGN_OP, "=", lex.Line, lex.Column)
		}
	case '!':
		// Could be '!' (logical NOT) or '!=' (not equal)
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(NE_OP, "!=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(NOT_OP, "!", lex.Line, lex.Column)
		}
	case '<':
		// Could be '<', '<=', '<<', or '<<='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(LE_OP, "<=", lex.Line, lex.Column)
		} else if lex.Peek() == '<' {
			lex.Advance()
			if lex.Peek() == '=' {
				lex.Advance()
				token = NewTokenWithMetadata(BIT_LEFT_ASSIGN, "<<=", lex.Line, lex.Column)
			} else {
				token = NewTokenWithMetadata(BIT_LEFT_OP, "<<", lex.Line, lex.Column)
			}
		} else {
			token = NewTokenWithMetadata(LT_OP, "<", lex.Line, lex.Column)
		}
	case '>':
		// Could be '>', '>=', '>>', or '>>='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(GE_OP, ">=", lex.Line, lex.Column)
		} else if lex.Peek() == '>' {
			lex.Advance()
			if lex.Peek() == '=' {
				lex.Advance()
				token = NewTokenWithMetadata(BIT_RIGHT_ASSIGN, ">>=", lex.Line, lex.Column)
			} else {
				token = NewTokenWithMetadata(BIT_RIGHT_OP, ">>", lex.Line, lex.Column)
			}
		} else {
			token = NewTokenWithMetadata(GT_OP, ">", lex.Line, lex.Column)
		}
	case '+':
		// Could be '+' or '+='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(PLUS_ASSIGN, "+=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(PLUS_OP, "+", lex.Line, lex.Column)
		}
	case '-':
		// Could be '-' or '-='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(MINUS_ASSIGN, "-=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MINUS_OP, "-", lex.Line, lex.Column)
		}
	case '*':
		// Could be '*' or '*='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(MUL_ASSIGN, "*=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MUL_OP, "*", lex.Line, lex.Column)
		}
	case '/':
		// Could be '/' or '/='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(DIV_ASSIGN, "/=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(DIV_OP, "/", lex.Line, lex.Column)
		}
	case '%':
		// Could be '%' or '%='
		if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(MOD_ASSIGN, "%=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(MOD_OP, "%", lex.Line, lex.Column)
		}
	case '^':
		// Could be '^' (XOR) or '^='
		if lex.Peek() == '=' {
			lex.Advance()
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
		if lex.Peek() == '.' {
			// Check if there's a third dot
			if lex.Position+2 < lex.SrcLength && lex.Src[lex.Position+2] == '.' {
				lex.Advance() // consume second dot
				lex.Advance() // consume third dot
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
		if lex.Peek() == '&' {
			lex.Advance()
			token = NewTokenWithMetadata(AND_OP, "&&", lex.Line, lex.Column)
		} else if lex.Peek() == '=' {
			lex.Advance()
			token = NewTokenWithMetadata(BIT_AND_ASSIGN, "&=", lex.Line, lex.Column)
		} else {
			token = NewTokenWithMetadata(BIT_AND_OP, "&", lex.Line, lex.Column)
		}
	case '|':
		// Could be '|' (bitwise OR), '||' (logical OR), or '|='
		if lex.Peek() == '|' {
			lex.Advance()
			token = NewTokenWithMetadata(OR_OP, "||", lex.Line, lex.Column)
		} else if lex.Peek() == '=' {
			lex.Advance()
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
	case '\'':
		// Character literal - delegate to specialized handler
		return readCharLiteral(lex)
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
	lex.Advance()

	return token
}

// Peek looks ahead to the next character in the source without consuming it.
// This is useful for lookahead when determining multi-character tokens.
//
// Returns:
//   - byte: The next character, or 0 if at end of source
//
// Example:
//
//	If current is '=' and next is '=', Peek() returns '='
func (lex *Lexer) Peek() byte {
	if lex.Position+1 >= lex.SrcLength {
		return 0 // End of source
	}
	return lex.Src[lex.Position+1]
}

// Advance moves the lexer to the next character in the source.
// It updates the Current byte, Position, and Column tracking.
// When a newline is encountered elsewhere, Line is incremented separately.
//
// After calling Advance:
//   - Position is incremented
//   - Column is incremented
//   - Current is set to the new character (or 0 if at end)
func (lex *Lexer) Advance() {
	lex.Position++
	lex.Column++

	if lex.Position >= lex.SrcLength {
		lex.Current = 0              // Null byte indicates end
		lex.Position = lex.SrcLength // Keep position at end
	} else {
		lex.Current = lex.Src[lex.Position]
	}
}

// IgnoreWhitespacesAndComments skips over whitespace and comments in the source.
// This method is called before tokenizing each meaningful token.
//
// It handles:
//   - Whitespace characters (space, tab, newline, etc.)
//   - Single-line comments (// ...)
//   - Multi-line comments (/* ... */)
//
// When a newline is encountered, the Line counter is incremented and
// Column is reset to 1.
func (lex *Lexer) IgnoreWhitespacesAndComments() {
	for {
		if isWhitespace(lex.Current) {
			// Track line numbers when encountering newlines
			if lex.Current == '\n' {
				lex.Line++
				lex.Column = 1
			}
			lex.Advance()
		} else if lex.Current == '/' && lex.Peek() == '/' {
			// Single-line comment detected
			lex.SkipSingleLineComment()
		} else if lex.Current == '/' && lex.Peek() == '*' {
			// Multi-line comment detected
			lex.SkipMultiLineComment()
		} else {
			// No more whitespace or comments
			break
		}
	}
}

// SkipSingleLineComment skips over a single-line comment (// ...).
// It advances the lexer until a newline or end of file is reached.
// The newline itself is not consumed, allowing line tracking to work correctly.
//
// Example:
//
//	Source: "// this is a comment\nvar x"
//	After skip: lexer is positioned at '\n'
func (lex *Lexer) SkipSingleLineComment() {
	// Skip the '//' characters
	lex.Advance()
	lex.Advance()

	// Skip until end of line or end of file
	for lex.Current != '\n' && lex.Current != 0 {
		lex.Advance()
	}
}

// SkipMultiLineComment skips over a multi-line comment (/* ... */).
// It advances the lexer until the closing */ is found or end of file is reached.
// Line numbers are tracked correctly even within multi-line comments.
//
// Example:
//
//	Source: "/* comment\nspanning lines */var x"
//	After skip: lexer is positioned after '*/'
func (lex *Lexer) SkipMultiLineComment() {
	// Skip the '/*' characters
	lex.Advance()
	lex.Advance()

	// Skip until we find '*/' or reach end of file
	for lex.Current != 0 {
		if lex.Current == '*' && lex.Peek() == '/' {
			// Found closing '*/' - skip it and exit
			lex.Advance()
			lex.Advance()
			break
		}
		lex.Advance()
	}
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
