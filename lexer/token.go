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
}

// NewToken(): constructor for Token
func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// TokenTypes
const (
	// Special Types
	EOF_TYPE TokenType = "EOF"

	// Arithmetic Operators
	PLUS_OP  TokenType = "+"
	MINUS_OP TokenType = "-"
	MUL_OP   TokenType = "*"
	DIV_OP   TokenType = "/"

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
	BIT_AND_OP   TokenType = "&"
	BIT_OR_OP    TokenType = "|"
	BIT_XOR_OP   TokenType = "^"
	BIT_NOT_OP   TokenType = "~"
	BIT_LEFT_OP  TokenType = "<<"
	BIT_RIGHT_OP TokenType = ">>"

	// Keywords
	FUNC_KEY     TokenType = "func"
	NEW_KEY      TokenType = "new"
	RETURN_KEY   TokenType = "return"
	VAR_KEY      TokenType = "var"
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
	"true":     TRUE_KEY,
	"false":    FALSE_KEY,
	"if":       IF_KEY,
	"else":     ELSE_KEY,
	"while":    WHILE_KEY,
	"for":      FOR_KEY,
	"break":    BREAK_KEY,
	"continue": CONTINUE_KEY,
}

// lookupIdent(): lookup the token type of an identifier
func lookupIdent(ident string) TokenType {
	if tok, ok := KEYWORDS_MAP[ident]; ok {
		return tok
	}
	return IDENTIFIER_ID
}
