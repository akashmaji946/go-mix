package lexer

import "fmt"

// every Token has a type(:TokenType) and value(:string)
type TokenType string

// <token_type, literal>
// <StringLiteral, "akash">
// <NumberLiteral, "123">
// <LeftParen, "(">
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// NewToken(): constructor for Token
func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// NewToken(): constructor for Token
func NewTokenWithMetadata(tokenType TokenType, literal string, line int, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    line,
		Column:  column,
	}
}

// TokenTypes
const (
	// Special Types
	EOF_TYPE     TokenType = "EOF"
	INVALID_TYPE TokenType = "INVALID"

	// Arithmetic Operators
	PLUS_OP      TokenType = "+"
	MINUS_OP     TokenType = "-"
	MUL_OP       TokenType = "*"
	DIV_OP       TokenType = "/"
	MOD_OP       TokenType = "%"
	PLUS_ASSIGN  TokenType = "+="
	MINUS_ASSIGN TokenType = "-="
	MUL_ASSIGN   TokenType = "*="
	DIV_ASSIGN   TokenType = "/="
	MOD_ASSIGN   TokenType = "%="

	// Logical Operators
	GT_OP     TokenType = ">"
	LT_OP     TokenType = "<"
	GE_OP     TokenType = ">="
	LE_OP     TokenType = "<="
	EQ_OP     TokenType = "=="
	NE_OP     TokenType = "!="
	ASSIGN_OP TokenType = "="
	NOT_OP    TokenType = "!"

	// Boolean Operators
	AND_OP TokenType = "&&"
	OR_OP  TokenType = "||"

	// Bitwise Operators
	BIT_NOT_OP TokenType = "~"

	BIT_AND_OP       TokenType = "&"
	BIT_OR_OP        TokenType = "|"
	BIT_XOR_OP       TokenType = "^"
	BIT_LEFT_OP      TokenType = "<<"
	BIT_RIGHT_OP     TokenType = ">>"
	BIT_AND_ASSIGN   TokenType = "&="
	BIT_OR_ASSIGN    TokenType = "|="
	BIT_XOR_ASSIGN   TokenType = "^="
	BIT_LEFT_ASSIGN  TokenType = "<<="
	BIT_RIGHT_ASSIGN TokenType = ">>="

	// Keywords
	FUNC_KEY     TokenType = "func"
	NEW_KEY      TokenType = "new"
	RETURN_KEY   TokenType = "return"
	VAR_KEY      TokenType = "var"
	LET_KEY      TokenType = "let"
	CONST_KEY    TokenType = "const"
	TRUE_KEY     TokenType = "true"
	FALSE_KEY    TokenType = "false"
	IF_KEY       TokenType = "if"
	ELSE_KEY     TokenType = "else"
	WHILE_KEY    TokenType = "while"
	FOR_KEY      TokenType = "for"
	BREAK_KEY    TokenType = "break"
	CONTINUE_KEY TokenType = "continue"

	// Identifiers
	IDENTIFIER_ID TokenType = "Identifier"
	NUMBER_ID     TokenType = "[0-9]"
	CHAR_ID       TokenType = "[a-zA-Z]"

	// Literals
	INT_LIT    TokenType = "IntLiteral"
	FLOAT_LIT  TokenType = "FloatLiteral"
	STRING_LIT TokenType = "StringLiteral"
	BOOL_LIT   TokenType = "BoolLiteral"
	NIL_LIT    TokenType = "NilLiteral"

	// Structural Tokens
	LEFT_PAREN    TokenType = "("
	RIGHT_PAREN   TokenType = ")"
	LEFT_BRACE    TokenType = "{"
	RIGHT_BRACE   TokenType = "}"
	LEFT_BRACKET  TokenType = "["
	RIGHT_BRACKET TokenType = "]"

	// Delimiters
	COMMA_DELIM     TokenType = ","
	SEMICOLON_DELIM TokenType = ";"
	COLON_DELIM     TokenType = ":"
)

// Print(): prints the token
func (tok *Token) Print() {
	fmt.Printf("%s:%v\n", tok.Literal, tok.Type)
}

// KEYWORDS_MAP: map of keywords to their token types
var KEYWORDS_MAP = map[string]TokenType{
	"func":     FUNC_KEY,
	"new":      NEW_KEY,
	"return":   RETURN_KEY,
	"var":      VAR_KEY,
	"let":      LET_KEY,
	"const":    CONST_KEY,
	"true":     TRUE_KEY,
	"false":    FALSE_KEY,
	"if":       IF_KEY,
	"else":     ELSE_KEY,
	"while":    WHILE_KEY,
	"for":      FOR_KEY,
	"break":    BREAK_KEY,
	"continue": CONTINUE_KEY,
	"nil":      NIL_LIT,
}

// lookupIdent(): lookup the token type of an identifier
func lookupIdent(ident string) TokenType {
	if tok, ok := KEYWORDS_MAP[ident]; ok {
		return tok
	}
	return IDENTIFIER_ID
}
