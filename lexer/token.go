/*
File    : go-mix/lexer/token.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package lexer

import "fmt"

// TokenType represents the type of a lexical token in the GoMix language.
// It is defined as a string to allow for easy comparison and debugging.
// Each token type corresponds to a specific syntactic element in the language,
// such as operators, keywords, literals, or structural symbols.
type TokenType string

// TokenType Constants:
// These constants define all possible token types in the GoMix language.
// They are organized into logical groups for clarity and maintainability.
const (
	// Special Types
	// EOF_TYPE marks the end of the input stream
	EOF_TYPE TokenType = "EOF"
	// INVALID_TYPE represents an unrecognized or malformed token
	INVALID_TYPE TokenType = "INVALID"

	// Arithmetic Operators
	// Basic arithmetic operations
	PLUS_OP  TokenType = "+" // Addition operator
	MINUS_OP TokenType = "-" // Subtraction operator
	MUL_OP   TokenType = "*" // Multiplication operator
	DIV_OP   TokenType = "/" // Division operator
	MOD_OP   TokenType = "%" // Modulo operator

	// Compound assignment operators (arithmetic)
	PLUS_ASSIGN  TokenType = "+=" // Add and assign (x += y)
	MINUS_ASSIGN TokenType = "-=" // Subtract and assign (x -= y)
	MUL_ASSIGN   TokenType = "*=" // Multiply and assign (x *= y)
	DIV_ASSIGN   TokenType = "/=" // Divide and assign (x /= y)
	MOD_ASSIGN   TokenType = "%=" // Modulo and assign (x %= y)

	// Logical/Comparison Operators
	GT_OP     TokenType = ">"  // Greater than
	LT_OP     TokenType = "<"  // Less than
	GE_OP     TokenType = ">=" // Greater than or equal to
	LE_OP     TokenType = "<=" // Less than or equal to
	EQ_OP     TokenType = "==" // Equality comparison
	NE_OP     TokenType = "!=" // Not equal comparison
	ASSIGN_OP TokenType = "="  // Assignment operator
	NOT_OP    TokenType = "!"  // Logical NOT operator

	// Boolean Operators
	AND_OP TokenType = "&&" // Logical AND
	OR_OP  TokenType = "||" // Logical OR

	// Bitwise Operators
	BIT_NOT_OP TokenType = "~" // Bitwise NOT (complement)

	// Bitwise binary operators
	BIT_AND_OP   TokenType = "&"  // Bitwise AND
	BIT_OR_OP    TokenType = "|"  // Bitwise OR
	BIT_XOR_OP   TokenType = "^"  // Bitwise XOR
	BIT_LEFT_OP  TokenType = "<<" // Bitwise left shift
	BIT_RIGHT_OP TokenType = ">>" // Bitwise right shift

	// Compound assignment operators (bitwise)
	BIT_AND_ASSIGN   TokenType = "&="  // Bitwise AND and assign
	BIT_OR_ASSIGN    TokenType = "|="  // Bitwise OR and assign
	BIT_XOR_ASSIGN   TokenType = "^="  // Bitwise XOR and assign
	BIT_LEFT_ASSIGN  TokenType = "<<=" // Left shift and assign
	BIT_RIGHT_ASSIGN TokenType = ">>=" // Right shift and assign

	// Keywords
	// Language keywords for control flow and declarations
	FUNC_KEY     TokenType = "func"     // Function declaration keyword
	NEW_KEY      TokenType = "new"      // Object instantiation keyword
	RETURN_KEY   TokenType = "return"   // Return statement keyword
	VAR_KEY      TokenType = "var"      // Variable declaration (mutable)
	LET_KEY      TokenType = "let"      // Variable declaration (type-locked)
	CONST_KEY    TokenType = "const"    // Constant declaration (immutable)
	TRUE_KEY     TokenType = "true"     // Boolean true literal
	FALSE_KEY    TokenType = "false"    // Boolean false literal
	IF_KEY       TokenType = "if"       // Conditional if keyword
	ELSE_KEY     TokenType = "else"     // Conditional else keyword
	WHILE_KEY    TokenType = "while"    // While loop keyword
	FOR_KEY      TokenType = "for"      // For loop keyword
	FOREACH_KEY  TokenType = "foreach"  // Foreach loop keyword
	IN_KEY       TokenType = "in"       // In keyword for foreach loops
	BREAK_KEY    TokenType = "break"    // Loop break keyword
	CONTINUE_KEY TokenType = "continue" // Loop continue keyword

	// Data Structure Literals
	ARRAY_KEY  TokenType = "array"  // Array literal keyword
	MAP_KEY    TokenType = "map"    // Map literal keyword
	SET_KEY    TokenType = "set"    // Set literal keyword
	STRUCT_KEY TokenType = "struct" // Struct declaration keyword

	// Identifiers
	// Token types for user-defined names and character classes
	IDENTIFIER_ID TokenType = "Identifier" // User-defined identifier (variable/function name)
	NUMBER_ID     TokenType = "[0-9]"      // Numeric character class
	CHAR_ID       TokenType = "[a-zA-Z]"   // Alphabetic character class

	// Literals
	// Token types for literal values in the source code
	INT_LIT    TokenType = "IntLiteral"    // Integer literal (e.g., 42, -10)
	FLOAT_LIT  TokenType = "FloatLiteral"  // Floating-point literal (e.g., 3.14, -0.5)
	STRING_LIT TokenType = "StringLiteral" // String literal (e.g., "hello")
	BOOL_LIT   TokenType = "BoolLiteral"   // Boolean literal (true or false)
	NIL_LIT    TokenType = "NilLiteral"    // Nil/null literal

	// Structural Tokens
	// Brackets and braces for grouping and scoping
	LEFT_PAREN    TokenType = "(" // Left parenthesis - function calls, grouping
	RIGHT_PAREN   TokenType = ")" // Right parenthesis
	LEFT_BRACE    TokenType = "{" // Left brace - code blocks, scopes
	RIGHT_BRACE   TokenType = "}" // Right brace
	LEFT_BRACKET  TokenType = "[" // Left bracket - array indexing, slicing
	RIGHT_BRACKET TokenType = "]" // Right bracket

	// Delimiters
	// Punctuation for separating elements
	COMMA_DELIM     TokenType = "," // Comma - separates parameters, array elements
	SEMICOLON_DELIM TokenType = ";" // Semicolon - statement terminator
	COLON_DELIM     TokenType = ":" // Colon - used in slicing operations

	// Range Operator
	RANGE_OP TokenType = "..." // Range operator - creates inclusive ranges (e.g., 2...5)

	// Object member access operator
	DOT_OP   TokenType = "."    // Dot operator - access struct fields and methods
	THIS_KEY TokenType = "this" // 'this' keyword for referring to the current struct instance member
	SELF_KEY TokenType = "self" // 'self' keyword for referring to the current class member

)

// KEYWORDS_MAP is a lookup table that maps keyword strings to their token types.
// This map is used during lexical analysis to distinguish between keywords
// (reserved words with special meaning) and regular identifiers (user-defined names).
//
// The map includes all language keywords such as control flow statements,
// declaration keywords, and boolean literals.
//
// Usage:
//
//	When the lexer encounters an identifier-like token, it checks this map
//	to determine if it's a keyword or a user-defined identifier.
var KEYWORDS_MAP = map[string]TokenType{
	"func":     FUNC_KEY,     // Function declaration
	"new":      NEW_KEY,      // Object creation
	"return":   RETURN_KEY,   // Return from function
	"var":      VAR_KEY,      // Mutable variable
	"let":      LET_KEY,      // Type-locked variable
	"const":    CONST_KEY,    // Immutable constant
	"true":     TRUE_KEY,     // Boolean true
	"false":    FALSE_KEY,    // Boolean false
	"if":       IF_KEY,       // Conditional if
	"else":     ELSE_KEY,     // Conditional else
	"while":    WHILE_KEY,    // While loop
	"for":      FOR_KEY,      // For loop
	"foreach":  FOREACH_KEY,  // Foreach loop
	"in":       IN_KEY,       // In keyword for foreach
	"break":    BREAK_KEY,    // Break from loop
	"continue": CONTINUE_KEY, // Continue to next iteration
	"array":    ARRAY_KEY,    // Array literal
	"struct":   STRUCT_KEY,   // Struct declaration
	"map":      MAP_KEY,      // Map literal
	"set":      SET_KEY,      // Set literal
	"nil":      NIL_LIT,      // Nil/null value
	"this":     THIS_KEY,     // 'this' keyword
	"self":     SELF_KEY,     // 'self' keyword
}

// Token represents a single lexical token in the GoMix source code.
// It contains the token's type, its literal string representation from the source,
// and metadata about its position in the source file (line and column numbers).
//
// Fields:
//   - Type: The category of the token (e.g., operator, keyword, literal)
//   - Literal: The actual string from the source code that this token represents
//   - Line: The line number where this token appears in the source (1-indexed)
//   - Column: The column number where this token starts in the source (1-indexed)
//
// Example:
//
//	For the source code "var x = 123" at line 5, column 10:
//	Token{Type: VAR_KEY, Literal: "var", Line: 5, Column: 10}
type Token struct {
	Type    TokenType // The type/category of this token
	Literal string    // The actual text from source code
	Line    int       // Line number in source file (1-indexed)
	Column  int       // Column number in source file (1-indexed)
}

// NewToken creates a new Token with the specified type and literal value.
// This is a basic constructor that does not set line/column metadata.
// Use NewTokenWithMetadata if position information is needed.
//
// Parameters:
//   - tokenType: The type of token to create
//   - literal: The string representation of the token from source code
//
// Returns:
//   - Token: A new token with the specified type and literal, but no position info
//
// Example:
//
//	token := NewToken(PLUS_OP, "+")
func NewToken(tokenType TokenType, literal string) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
	}
}

// NewTokenWithMetadata creates a new Token with full metadata including position.
// This constructor should be used during lexical analysis to preserve source location
// information, which is essential for error reporting and debugging.
//
// Parameters:
//   - tokenType: The type of token to create
//   - literal: The string representation of the token from source code
//   - line: The line number where the token appears (1-indexed)
//   - column: The column number where the token starts (1-indexed)
//
// Returns:
//   - Token: A new token with complete type, literal, and position information
//
// Example:
//
//	token := NewTokenWithMetadata(INT_LIT, "42", 10, 5)
func NewTokenWithMetadata(tokenType TokenType, literal string, line int, column int) Token {
	return Token{
		Type:    tokenType,
		Literal: literal,
		Line:    line,
		Column:  column,
	}
}

// Print outputs a human-readable representation of the token to standard output.
// The format is "literal:type", which shows both the actual text and its classification.
// This is primarily used for debugging and development purposes.
//
// Example output:
//
//	For Token{Type: PLUS_OP, Literal: "+"}:
//	Output: "+:+"
func (tok *Token) Print() {
	fmt.Printf("%s:%v\n", tok.Literal, tok.Type)
}

// lookupIdent determines the token type for an identifier string.
// It checks if the identifier is a reserved keyword by looking it up in KEYWORDS_MAP.
// If found, it returns the corresponding keyword token type; otherwise, it returns
// IDENTIFIER_ID to indicate a user-defined identifier.
//
// This function is essential for the lexer to correctly classify tokens during
// the tokenization process, ensuring that keywords are not treated as regular
// variable or function names.
//
// Parameters:
//   - ident: The identifier string to look up
//
// Returns:
//   - TokenType: The keyword token type if ident is a keyword, otherwise IDENTIFIER_ID
//
// Example:
//
//	lookupIdent("if")    -> IF_KEY
//	lookupIdent("myVar") -> IDENTIFIER_ID
func lookupIdent(ident string) TokenType {
	// Check if the identifier is a keyword
	if tok, ok := KEYWORDS_MAP[ident]; ok {
		return tok
	}
	// Not a keyword, so it's a user-defined identifier
	return IDENTIFIER_ID
}
