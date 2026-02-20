/*
File    : go-mix/parser/parser_literals.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

package parser

import (
	"fmt"
	"strconv"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// parseNumberLiteral parses integer literal expressions.
// Handles both positive and negative integers, including edge cases
// like the minimum int64 value.
//
// Returns:
//
//	An IntegerLiteralExpressionNode with the parsed value
//
// Special handling:
//   - Uses ParseUint for values that overflow ParseInt (like -9223372036854775808)
//   - Reports errors for invalid number formats
//
// Examples:
//
//	42, -17, 0, 9223372036854775807
func (par *Parser) parseNumberLiteral() ExpressionNode {
	token := par.CurrToken
	val, err := strconv.ParseInt(token.Literal, 0, 64)
	if err != nil {
		uVal, uErr := strconv.ParseUint(token.Literal, 0, 64)
		if uErr == nil {
			val = int64(uVal)
		} else {
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: could not parse number literal: %s",
				token.Line, token.Column, token.Literal)
			par.addError(msg)
			return nil
		}
	}
	return &IntegerLiteralExpressionNode{
		Token: token,
		Value: &std.Integer{Value: val},
	}
}

// parseCharLiteral parses character literal expressions.
func (par *Parser) parseCharLiteral() ExpressionNode {
	token := par.CurrToken
	// The literal in the token is the character itself
	runes := []rune(token.Literal)
	var r rune
	if len(runes) > 0 {
		r = runes[0]
	}
	return &CharLiteralExpressionNode{
		Token: token,
		Value: &std.Char{Value: r},
	}
}

// parseFloatLiteral parses floating-point literal expressions.
//
// Returns:
//
//	A FloatLiteralExpressionNode with the parsed value
//
// Examples:
//
//	3.14, -2.5, 0.001, 1.0
func (par *Parser) parseFloatLiteral() ExpressionNode {
	token := par.CurrToken
	val, err := strconv.ParseFloat(token.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: could not parse float literal: %s",
			token.Line, token.Column, token.Literal)
		par.addError(msg)
		return nil
	}
	return &FloatLiteralExpressionNode{
		Token: token,
		Value: &std.Float{Value: val},
	}
}

// parseBooleanLiteral parses boolean literal expressions.
//
// Returns:
//
//	A BooleanLiteralExpressionNode with value true or false
//
// Examples:
//
//	true, false
func (par *Parser) parseBooleanLiteral() ExpressionNode {
	token := par.CurrToken
	return &BooleanLiteralExpressionNode{
		Token: token,
		Value: &std.Boolean{Value: token.Type == lexer.TRUE_KEY},
	}
}

// parseStringLiteral parses string literal expressions.
//
// Returns:
//
//	A StringLiteralExpressionNode
//
// Examples:
//
//	"hello", "world", "Go-Mix is awesome!"
func (par *Parser) parseStringLiteral() ExpressionNode {
	return &StringLiteralExpressionNode{
		Token: par.CurrToken,
		Value: &std.String{Value: par.CurrToken.Literal},
	}
}

// parseNilLiteral parses nil literal expressions.
// Nil represents the absence of a value.
//
// Returns:
//
//	A NilLiteralExpressionNode
//
// Example:
//
//	nil
func (par *Parser) parseNilLiteral() ExpressionNode {
	return &NilLiteralExpressionNode{
		Token: par.CurrToken,
		Value: &std.Nil{},
	}
}
