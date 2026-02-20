/*
File    : go-mix/parser/parser_expressions.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// parseExpression is the entry point for parsing expressions.
// It delegates to parseInternal with minimum precedence, allowing
// all operators to be parsed.
//
// Returns:
//
//	An ExpressionNode representing the parsed expression
//
// This function uses the Pratt parsing algorithm, which handles
// operator precedence and associativity elegantly.
func (par *Parser) parseExpression() ExpressionNode {
	return par.parseInternal(MINIMUM_PRIORITY)
}

// parseParenthesizedExpression parses expressions enclosed in parentheses.
// Parentheses are used for grouping and overriding operator precedence.
//
// Syntax:
//
//	(expression)
//
// Returns:
//
//	A ParenthesizedExpressionNode containing the inner expression
//
// Examples:
//
//	(5 + 3) * 2  - Parentheses force addition before multiplication
//	(a && b) || c
func (par *Parser) parseParenthesizedExpression() ExpressionNode {
	// we are already at the LEFT_PAREN, so just advance
	par.advance()
	paren := &ParenthesizedExpressionNode{}
	paren.Expr = par.parseExpression()
	if paren.Expr == nil {
		return nil
	}
	paren.Value = parseEval(par, paren.Expr)
	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

	return paren
}

// parseBinaryExpression parses binary (infix) expressions.
// Binary expressions have the form: left operator right
//
// Parameters:
//
//	left - The already-parsed left operand
//
// Returns:
//
//	A BinaryExpressionNode representing the complete expression
//
// Supported operators:
//
//	Arithmetic: +, -, *, /, %
//	Bitwise: &, |, ^, <<, >>
//
// The function evaluates the expression during parsing for constant folding.
//
// Examples:
//
//	5 + 3, a * b, x << 2, y & 0xFF
func (par *Parser) parseBinaryExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)
	if right == nil {
		return nil
	}

	lVal := parseEval(par, left)
	rVal := parseEval(par, right)

	var val std.GoMixObject = &std.Nil{}

	if lVal.GetType() == std.IntegerType && rVal.GetType() == std.IntegerType {
		l := lVal.(*std.Integer).Value
		r := rVal.(*std.Integer).Value
		switch op.Type {
		case lexer.PLUS_OP:
			val = &std.Integer{Value: l + r}
		case lexer.MINUS_OP:
			val = &std.Integer{Value: l - r}
		case lexer.MUL_OP:
			val = &std.Integer{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				val = &std.Integer{Value: l / r}
			}
		case lexer.MOD_OP:
			if r != 0 {
				val = &std.Integer{Value: l % r}
			}
		case lexer.BIT_AND_OP:
			val = &std.Integer{Value: l & r}
		case lexer.BIT_OR_OP:
			val = &std.Integer{Value: l | r}
		case lexer.BIT_XOR_OP:
			val = &std.Integer{Value: l ^ r}
		case lexer.BIT_LEFT_OP:
			val = &std.Integer{Value: l << r}
		case lexer.BIT_RIGHT_OP:
			val = &std.Integer{Value: l >> r}
		}
	} else if (lVal.GetType() == std.IntegerType || lVal.GetType() == std.FloatType) &&
		(rVal.GetType() == std.IntegerType || rVal.GetType() == std.FloatType) {
		// Mixed arithmetic
		l := toFloat64(lVal)
		r := toFloat64(rVal)
		switch op.Type {
		case lexer.PLUS_OP:
			val = &std.Float{Value: l + r}
		case lexer.MINUS_OP:
			val = &std.Float{Value: l - r}
		case lexer.MUL_OP:
			val = &std.Float{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				val = &std.Float{Value: l / r}
			}
		}
	}

	return &BinaryExpressionNode{
		Left:      left,
		Operation: op,
		Right:     right,
		Value:     val,
	}
}

// parseUnaryExpression parses unary (prefix) expressions.
// Unary expressions have an operator followed by an operand.
//
// Returns:
//
//	A UnaryExpressionNode representing the expression
//
// Supported operators:
//
//	! (logical NOT)    - Negates boolean values
//	- (unary minus)    - Negates numbers
//	+ (unary plus)     - Identity operation (returns value unchanged)
//	~ (bitwise NOT)    - Bitwise complement
//
// Examples:
//
//	!true (false), -5 (-5), +3 (3), ~0xFF (bitwise complement)
func (par *Parser) parseUnaryExpression() ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)
	if right == nil {
		return nil
	}

	rVal := parseEval(par, right)
	var val std.GoMixObject = &std.Nil{}

	switch op.Type {
	case lexer.NOT_OP:
		// Logical NOT
		// Truthiness check
		isTrue := false
		if rVal.GetType() == std.BooleanType {
			isTrue = rVal.(*std.Boolean).Value
		} else if rVal.GetType() == std.IntegerType {
			isTrue = rVal.(*std.Integer).Value != 0
		}
		val = &std.Boolean{Value: !isTrue}

	case lexer.MINUS_OP:
		if rVal.GetType() == std.IntegerType {
			val = &std.Integer{Value: -rVal.(*std.Integer).Value}
		} else if rVal.GetType() == std.FloatType {
			val = &std.Float{Value: -rVal.(*std.Float).Value}
		}
	case lexer.PLUS_OP:
		val = rVal
	case lexer.BIT_NOT_OP:
		if rVal.GetType() == std.IntegerType {
			val = &std.Integer{Value: ^rVal.(*std.Integer).Value}
		}
	}

	return &UnaryExpressionNode{
		Operation: op,
		Right:     right,
		Value:     val,
	}
}

// parseIdentifierExpression parses identifier expressions.
// An identifier can be either a variable reference or a function call.
//
// Returns:
//
//	Either an IdentifierExpressionNode or a CallExpressionNode
//
// The function checks the next token to determine if this is a function call
// (next token is '(') or a simple variable reference.
//
// Examples:
//
//	x          - Variable reference
//	myFunc()   - Function call
func (par *Parser) parseIdentifierExpression() ExpressionNode {

	// may be an identifier expression or a function call expression
	if par.NextToken.Type == lexer.LEFT_PAREN {
		return par.parseCallExpression()
	}

	varToken := par.CurrToken

	// get the value from the environment
	val := par.Env[varToken.Literal]
	if val == nil {
		val = &std.Nil{}
	}

	// Determine if this is a const or let or var
	ident := &IdentifierExpressionNode{
		Token: varToken,
		Name:  varToken.Literal,
		Value: val,
		Type:  "var", // default type
		IsLet: false,
	}

	// Check if this identifier is a const
	if par.Consts[varToken.Literal] {
		ident.Type = "const"
	}

	// Check if this identifier is a let
	if par.LetVars[varToken.Literal] {
		ident.Type = "let"
		ident.IsLet = true
	}

	return ident
}

// parseBooleanExpression parses boolean/comparison expressions.
// These are binary expressions that produce boolean results.
//
// Parameters:
//
//	left - The already-parsed left operand
//
// Returns:
//
//	A BooleanExpressionNode with a boolean value
//
// Supported operators:
//
//	Comparison: <, >, <=, >=, ==, !=
//	Logical: &&, ||
//
// The function handles mixed-type comparisons (int/float) and
// truthiness for logical operators.
//
// Examples:
//
//	5 < 10, x >= y, a == b, true && false
func (par *Parser) parseBooleanExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)
	if right == nil {
		return nil
	}

	lVal := parseEval(par, left)
	rVal := parseEval(par, right)
	val := false

	// Comparison logic
	if lVal.GetType() == std.IntegerType && rVal.GetType() == std.IntegerType {
		l := lVal.(*std.Integer).Value
		r := rVal.(*std.Integer).Value
		switch op.Type {
		case lexer.GT_OP:
			val = l > r
		case lexer.LT_OP:
			val = l < r
		case lexer.GE_OP:
			val = l >= r
		case lexer.LE_OP:
			val = l <= r
		case lexer.EQ_OP:
			val = l == r
		case lexer.NE_OP:
			val = l != r
		case lexer.STRICT_EQ_OP:
			val = l == r
		case lexer.STRICT_NE_OP:
			val = l != r
		case lexer.AND_OP: // Logical AND/OR on integers should treat them as truthy/falsy
			val = (l != 0) && (r != 0)
		case lexer.OR_OP:
			val = (l != 0) || (r != 0)
		}
	} else if lVal.GetType() == std.BooleanType && rVal.GetType() == std.BooleanType {
		l := lVal.(*std.Boolean).Value
		r := rVal.(*std.Boolean).Value
		switch op.Type {
		case lexer.AND_OP:
			val = l && r
		case lexer.OR_OP:
			val = l || r
		case lexer.EQ_OP:
			val = l == r
		case lexer.NE_OP:
			val = l != r
		case lexer.STRICT_EQ_OP:
			val = l == r
		case lexer.STRICT_NE_OP:
			val = l != r
		}
	} else if (lVal.GetType() == std.FloatType || lVal.GetType() == std.IntegerType) &&
		(rVal.GetType() == std.FloatType || rVal.GetType() == std.IntegerType) {
		// Mixed float/integer comparison
		l := toFloat64(lVal)
		r := toFloat64(rVal)
		switch op.Type {
		case lexer.GT_OP:
			val = l > r
		case lexer.LT_OP:
			val = l < r
		case lexer.GE_OP:
			val = l >= r
		case lexer.LE_OP:
			val = l <= r
		case lexer.EQ_OP:
			val = l == r
		case lexer.NE_OP:
			val = l != r
		case lexer.STRICT_EQ_OP:
			if lVal.GetType() != rVal.GetType() {
				val = false
			} else {
				val = l == r
			}
		case lexer.STRICT_NE_OP:
			if lVal.GetType() != rVal.GetType() {
				val = true
			} else {
				val = l != r
			}
		}
	} else {
		// Fallback for other types, e.g., string comparison for equality
		switch op.Type {
		case lexer.EQ_OP:
			val = lVal.ToString() == rVal.ToString()
		case lexer.NE_OP:
			val = lVal.ToString() != rVal.ToString()
		case lexer.STRICT_EQ_OP:
			if lVal.GetType() != rVal.GetType() {
				val = false
			} else {
				if lVal.GetType() == std.StringType {
					val = lVal.(*std.String).Value == rVal.(*std.String).Value
				} else if lVal.GetType() == std.CharType {
					val = lVal.(*std.Char).Value == rVal.(*std.Char).Value
				} else if lVal.GetType() == std.NilType {
					val = true
				} else {
					val = lVal == rVal
				}
			}
		case lexer.STRICT_NE_OP:
			if lVal.GetType() != rVal.GetType() {
				val = true
			} else {
				if lVal.GetType() == std.StringType {
					val = lVal.(*std.String).Value != rVal.(*std.String).Value
				} else if lVal.GetType() == std.CharType {
					val = lVal.(*std.Char).Value != rVal.(*std.Char).Value
				} else if lVal.GetType() == std.NilType {
					val = false
				} else {
					val = lVal != rVal
				}
			}
		case lexer.AND_OP: // Treat as truthy/falsy
			isLTrue := (lVal.GetType() == std.BooleanType && lVal.(*std.Boolean).Value) || (lVal.GetType() == std.IntegerType && lVal.(*std.Integer).Value != 0)
			isRTrue := (rVal.GetType() == std.BooleanType && rVal.(*std.Boolean).Value) || (rVal.GetType() == std.IntegerType && rVal.(*std.Integer).Value != 0)
			val = isLTrue && isRTrue
		case lexer.OR_OP: // Treat as truthy/falsy
			isLTrue := (lVal.GetType() == std.BooleanType && lVal.(*std.Boolean).Value) || (lVal.GetType() == std.IntegerType && lVal.(*std.Integer).Value != 0)
			isRTrue := (rVal.GetType() == std.BooleanType && rVal.(*std.Boolean).Value) || (rVal.GetType() == std.IntegerType && rVal.(*std.Integer).Value != 0)
			val = isLTrue || isRTrue
		}
	}

	return &BooleanExpressionNode{
		Operation: op,
		Left:      left,
		Right:     right,
		Value:     &std.Boolean{Value: val},
	}
}
