/*
File    : go-mix/parser/parser_precedence.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import "github.com/akashmaji946/go-mix/lexer"

// Operator precedence constants (following C-based language standards)
// Higher number = higher precedence (binds tighter)
//
// Precedence Hierarchy (lowest to highest):
// 1. Assignment operators (right-to-left associativity)
// 2. Logical OR
// 3. Logical AND
// 4. Bitwise OR
// 5. Bitwise XOR
// 6. Bitwise AND
// 7. Equality operators
// 8. Relational operators
// 9. Shift operators
// 10. Additive operators
// 11. Multiplicative operators
// 12. Unary/Prefix operators
// 13. Parentheses
// 14. Index/Call operators (postfix)
//
// Example: In "a + b * c", multiplication has higher precedence than addition,
// so it's parsed as "a + (b * c)" rather than "(a + b) * c"
const (
	MINIMUM_PRIORITY = 0 // Base priority for starting expression parsing

	// Assignment operators (lowest precedence, right-to-left associativity)
	// Operators: = += -= *= /= %= &= |= ^= <<= >>=
	// Example: a = b = 5 is parsed as a = (b = 5)
	ASSIGN_PRIORITY = 10

	// Logical OR: ||
	// Example: a || b || c is parsed left-to-right
	OR_PRIORITY = 40

	// Logical AND: &&
	// Example: a && b has higher precedence than a || b
	AND_PRIORITY = 50

	// Bitwise OR: |
	// Example: a | b
	BIT_OR_PRIORITY = 60

	// Bitwise XOR: ^
	// Example: a ^ b
	BIT_XOR_PRIORITY = 70

	// Bitwise AND: &
	// Example: a & b
	BIT_AND_PRIORITY = 80

	// Equality operators: == !=
	// Example: a == b, a != b
	EQUALITY_PRIORITY = 90

	// Relational operators: < > <= >=
	// Example: a < b, a >= b
	RELATIONAL_PRIORITY = 100

	// Range operator: ...
	// Example: 2...5 (inclusive range)
	RANGE_PRIORITY = 105

	// Shift operators: << >>
	// Example: a << 2, b >> 1
	SHIFT_PRIORITY = 110

	// Additive operators: + -
	// Example: a + b, a - b
	PLUS_PRIORITY = 120

	// Multiplicative operators: * / %
	// Example: a * b, a / b, a % b
	MUL_PRIORITY = 130

	// Unary/Prefix operators: ! ~ + -
	// Example: !a, -b, ~c
	PREFIX_PRIORITY = 140

	// member access operator: .
	// Example: obj.field, obj.method()
	MEMBER_ACCESS_PRIORITY = 145

	// Parentheses (highest precedence for grouping)
	// Example: (a + b) * c
	PAREN_PRIORITY = 150

	// Index/Call operators (highest precedence for postfix operations)
	// Example: arr[0], func()
	INDEX_PRIORITY = 160
)

// getPrecedence returns the precedence level for a given token.
// This function is central to the Pratt parsing algorithm, determining
// how tightly operators bind to their operands.
//
// Parameters:
//
//	token - The token to get precedence for
//
// Returns:
//
//	An integer representing the precedence level (higher = tighter binding)
//	Returns -1 for tokens that are not operators
//
// The precedence values follow C-based language standards, ensuring
// familiar operator behavior for users coming from C, Java, JavaScript, etc.
func getPrecedence(token *lexer.Token) int {
	switch token.Type {

	// Parentheses - highest precedence
	case lexer.LEFT_PAREN:
		return PAREN_PRIORITY

	// Index operator - highest precedence for postfix
	case lexer.LEFT_BRACKET:
		return INDEX_PRIORITY

	// Unary/Prefix operators: ! ~
	case lexer.NOT_OP, lexer.BIT_NOT_OP:
		return PREFIX_PRIORITY

	// Multiplicative: * / %
	case lexer.MUL_OP, lexer.DIV_OP, lexer.MOD_OP:
		return MUL_PRIORITY

	// Additive: + -
	case lexer.PLUS_OP, lexer.MINUS_OP:
		return PLUS_PRIORITY

	// Shift: << >>
	case lexer.BIT_LEFT_OP, lexer.BIT_RIGHT_OP:
		return SHIFT_PRIORITY

	// Relational: < > <= >=
	case lexer.GT_OP, lexer.LT_OP, lexer.GE_OP, lexer.LE_OP:
		return RELATIONAL_PRIORITY

	// Range: ...
	case lexer.RANGE_OP:
		return RANGE_PRIORITY

	// Equality: == !=
	case lexer.EQ_OP, lexer.NE_OP:
		return EQUALITY_PRIORITY

	// Bitwise AND: &
	case lexer.BIT_AND_OP:
		return BIT_AND_PRIORITY

	// Bitwise XOR: ^
	case lexer.BIT_XOR_OP:
		return BIT_XOR_PRIORITY

	// Bitwise OR: |
	case lexer.BIT_OR_OP:
		return BIT_OR_PRIORITY

	// Logical AND: &&
	case lexer.AND_OP:
		return AND_PRIORITY

	// Logical OR: ||
	case lexer.OR_OP:
		return OR_PRIORITY

	// Assignment operators (lowest precedence)
	case lexer.ASSIGN_OP, lexer.PLUS_ASSIGN, lexer.MINUS_ASSIGN, lexer.MUL_ASSIGN, lexer.DIV_ASSIGN, lexer.MOD_ASSIGN,
		lexer.BIT_AND_ASSIGN, lexer.BIT_OR_ASSIGN, lexer.BIT_XOR_ASSIGN, lexer.BIT_LEFT_ASSIGN, lexer.BIT_RIGHT_ASSIGN:
		return ASSIGN_PRIORITY

	// Member access operator: .
	case lexer.DOT_OP:
		return MEMBER_ACCESS_PRIORITY

	default:
		return -1 // Not an operator token
	}
}

// binaryParseFunction is a function type for parsing binary expressions.
// Binary expressions have a left operand, an operator, and a right operand.
//
// Parameters:
//
//	ExpressionNode - The already-parsed left operand
//
// Returns:
//
//	ExpressionNode - The complete binary expression node
//
// Example: For "a + b", when parsing "+", the left operand "a" is passed in,
// and the function parses "b" and returns the complete "a + b" expression.
type binaryParseFunction func(ExpressionNode) ExpressionNode

// unaryParseFunction is a function type for parsing unary/prefix expressions.
// Unary expressions have an operator followed by an operand.
//
// Returns:
//
//	ExpressionNode - The parsed expression node
//
// Example: For "-5", the function parses the entire expression and returns
// a unary expression node representing the negation of 5.
type unaryParseFunction func() ExpressionNode

// registerUnaryFuncs is a helper to register a unary parsing function
// for multiple token types.
//
// Parameters:
//
//	f          - The parsing function to register
//	tokenTypes - Variable number of token types to associate with the function
//
// This allows one parsing function to handle multiple related token types.
// For example, parseUnaryExpression handles !, -, +, and ~ operators.
func (par *Parser) registerUnaryFuncs(f unaryParseFunction, tokenTypes ...lexer.TokenType) {
	for _, tokenType := range tokenTypes {
		par.UnaryFuncs[tokenType] = f
	}
}

// registerBinaryFuncs is a helper to register a binary parsing function
// for multiple token types.
//
// Parameters:
//
//	f          - The parsing function to register
//	tokenTypes - Variable number of token types to associate with the function
//
// This allows one parsing function to handle multiple related token types.
// For example, parseBinaryExpression handles +, -, *, /, and % operators.
func (par *Parser) registerBinaryFuncs(f binaryParseFunction, tokenTypes ...lexer.TokenType) {
	for _, tokenType := range tokenTypes {
		par.BinaryFuncs[tokenType] = f
	}
}
