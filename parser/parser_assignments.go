package parser

import (
	"fmt"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// parseAssignmentExpression parses assignment expressions.
// Handles both simple assignments (=) and compound assignments (+=, -=, etc.).
// Supports identifier assignment (x = 10) and index assignment (a[0] = 11).
//
// Parameters:
//
//	left - The left-hand side expression (must be an identifier or index expression)
//
// Returns:
//
//	An AssignmentExpressionNode
//
// Supported operators:
//
//	Simple: =
//	Compound: +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=
//
// Examples:
//
//	x = 10
//	count += 1
//	value *= 2
//	a[0] = 11
//	map["key"] = value
func (par *Parser) parseAssignmentExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)
	if right == nil {
		return nil
	}

	// Check if left is a valid assignment target (identifier or index expression)
	_, isIdent := left.(*IdentifierExpressionNode)
	_, isIndex := left.(*IndexExpressionNode)

	isMember := false
	if binNode, ok := left.(*BinaryExpressionNode); ok {
		if binNode.Operation.Type == lexer.DOT_OP {
			isMember = true
		}
	}

	if !isIdent && !isIndex && !isMember {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: invalid assignment target", par.CurrToken.Line, par.CurrToken.Column)
		par.addError(msg)
		return nil
	}

	// Handle identifier assignment
	if isIdent {
		ident := left.(*IdentifierExpressionNode)

		// Check if this is a compound assignment (+=, -=, *=, etc.)
		if binaryOp, isCompound := getCompoundBinaryOp(op.Type); isCompound {
			return par.parseCompoundAssignment(left, ident, op, right, binaryOp)
		}

		// Regular assignment - evaluate and store
		val := parseEval(par, right)
		par.Env[ident.Name] = val
		return &AssignmentExpressionNode{
			Operation: op,
			Left:      ident,
			Right:     right,
			Value:     val,
		}
	}

	// Handle index assignment (e.g., a[0] = 11)
	// For index expressions, we don't evaluate at parse time - let the evaluator handle it
	return &AssignmentExpressionNode{
		Operation: op,
		Left:      left,
		Right:     right,
		Value:     &std.Nil{},
	}
}

// parseCompoundAssignment handles compound assignment expressions (+=, -=, *=, etc.).
// It transforms them into regular assignments with a binary expression on the right side.
//
// Parameters:
//
//	left     - The left-hand side expression (the variable being assigned to)
//	ident    - The identifier node for the variable
//	op       - The compound assignment operator token
//	right    - The right-hand side expression
//	binaryOp - The underlying binary operator
//
// Returns:
//
//	An AssignmentExpressionNode representing the expanded assignment
//
// Transformation:
//
//	a += 5  becomes  a = a + 5
//	x *= 2  becomes  x = x * 2
//
// The function evaluates the expression and updates the parser's environment.
func (par *Parser) parseCompoundAssignment(left ExpressionNode, ident *IdentifierExpressionNode, op lexer.Token, right ExpressionNode, binaryOp lexer.TokenType) ExpressionNode {
	// Evaluate the left operand (get current value from environment)
	lVal := parseEval(par, left)
	// Evaluate the right operand
	rVal := parseEval(par, right)

	// Compute the binary expression value
	binaryVal := evaluateCompoundBinaryOp(lVal, rVal, binaryOp)

	// Update the environment with the new value
	par.Env[ident.Name] = binaryVal

	// Create a binary expression: left op right (e.g., a + 1)
	binaryExpr := &BinaryExpressionNode{
		Left:      left,
		Operation: lexer.Token{Type: binaryOp, Literal: string(binaryOp), Line: op.Line, Column: op.Column},
		Right:     right,
		Value:     binaryVal,
	}

	// Create assignment: left = binaryExpr (e.g., a = a + 1)
	assignOp := lexer.Token{Type: lexer.ASSIGN_OP, Literal: "=", Line: op.Line, Column: op.Column}
	return &AssignmentExpressionNode{
		Operation: assignOp,
		Left:      ident,
		Right:     binaryExpr,
		Value:     binaryVal,
	}
}

// evaluateCompoundBinaryOp evaluates a binary operation for compound assignments.
// This function is used to compute the result of compound assignment operators
// like +=, -=, *=, etc.
//
// Parameters:
//
//	lVal     - The left operand value
//	rVal     - The right operand value
//	binaryOp - The binary operator type (e.g., PLUS_OP for +=)
//
// Returns:
//
//	The computed result as a GoMixObject
//
// Examples:
//
//	For "a += 5", this computes a + 5
//	For "x *= 2", this computes x * 2
func evaluateCompoundBinaryOp(lVal, rVal std.GoMixObject, binaryOp lexer.TokenType) std.GoMixObject {
	if lVal.GetType() == std.IntegerType && rVal.GetType() == std.IntegerType {
		l := lVal.(*std.Integer).Value
		r := rVal.(*std.Integer).Value
		switch binaryOp {
		case lexer.PLUS_OP:
			return &std.Integer{Value: l + r}
		case lexer.MINUS_OP:
			return &std.Integer{Value: l - r}
		case lexer.MUL_OP:
			return &std.Integer{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				return &std.Integer{Value: l / r}
			}
		case lexer.MOD_OP:
			if r != 0 {
				return &std.Integer{Value: l % r}
			}
		case lexer.BIT_AND_OP:
			return &std.Integer{Value: l & r}
		case lexer.BIT_OR_OP:
			return &std.Integer{Value: l | r}
		case lexer.BIT_XOR_OP:
			return &std.Integer{Value: l ^ r}
		case lexer.BIT_LEFT_OP:
			return &std.Integer{Value: l << r}
		case lexer.BIT_RIGHT_OP:
			return &std.Integer{Value: l >> r}
		}
	} else if (lVal.GetType() == std.IntegerType || lVal.GetType() == std.FloatType) &&
		(rVal.GetType() == std.IntegerType || rVal.GetType() == std.FloatType) {
		// Mixed arithmetic
		l := toFloat64(lVal)
		r := toFloat64(rVal)
		switch binaryOp {
		case lexer.PLUS_OP:
			return &std.Float{Value: l + r}
		case lexer.MINUS_OP:
			return &std.Float{Value: l - r}
		case lexer.MUL_OP:
			return &std.Float{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				return &std.Float{Value: l / r}
			}
		}
	}
	return &std.Nil{}
}

// getCompoundBinaryOp returns the corresponding binary operator for a compound assignment operator.
// This function maps compound assignments to their underlying binary operations.
//
// Parameters:
//
//	opType - The compound assignment operator type
//
// Returns:
//
//	The corresponding binary operator type and true if it's a compound assignment,
//	otherwise returns empty string and false
//
// Mappings:
//
//	+= -> +, -= -> -, *= -> *, /= -> /, %= -> %
//	&= -> &, |= -> |, ^= -> ^, <<= -> <<, >>= -> >>
func getCompoundBinaryOp(opType lexer.TokenType) (lexer.TokenType, bool) {
	switch opType {
	case lexer.PLUS_ASSIGN:
		return lexer.PLUS_OP, true
	case lexer.MINUS_ASSIGN:
		return lexer.MINUS_OP, true
	case lexer.MUL_ASSIGN:
		return lexer.MUL_OP, true
	case lexer.DIV_ASSIGN:
		return lexer.DIV_OP, true
	case lexer.MOD_ASSIGN:
		return lexer.MOD_OP, true
	case lexer.BIT_AND_ASSIGN:
		return lexer.BIT_AND_OP, true
	case lexer.BIT_OR_ASSIGN:
		return lexer.BIT_OR_OP, true
	case lexer.BIT_XOR_ASSIGN:
		return lexer.BIT_XOR_OP, true
	case lexer.BIT_LEFT_ASSIGN:
		return lexer.BIT_LEFT_OP, true
	case lexer.BIT_RIGHT_ASSIGN:
		return lexer.BIT_RIGHT_OP, true
	default:
		return "", false
	}
}
