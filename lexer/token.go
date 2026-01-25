package lexer

// every token has a type(:TokenType) and value(:string)
type TokenType string

// <token_type, literal>
// <StringLiteral, "akash">
// <NumberLiteral, "123">
// <LeftParen, "(">
type Token struct {
	Type    TokenType
	Literal string
}

// constructor for Token
func NewToken(tokenType TokenType, literal string) *Token {
	return &Token{
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
	GT_OP TokenType = ">"
	LT_OP TokenType = "<"
	GE_OP TokenType = ">="
	LE_OP TokenType = "<="
	EQ_OP TokenType = "=="
	NE_OP TokenType = "!="

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
