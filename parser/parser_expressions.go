/*
File    : go-mix/parser/parser_expressions.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"fmt"
	"strconv"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
)

// parseStatement parses a single statement.
// This is the main dispatcher that determines what type of statement to parse
// based on the current token.
//
// Returns:
//
//	A StatementNode representing the parsed statement, or nil for empty statements
//
// Supported statement types:
//   - Variable declarations (var, let, const)
//   - Block statements ({ ... })
//   - If statements
//   - Function declarations
//   - For loops
//   - While loops
//   - Expression statements (any expression followed by semicolon)
func (par *Parser) parseStatement() StatementNode {
	switch par.CurrToken.Type {

	// ignore semicolons
	case lexer.SEMICOLON_DELIM:
		return nil

	// TODO: add for statements like

	// var a = 10;
	case lexer.VAR_KEY:
		return par.parseDeclarativeStatement()
	case lexer.LET_KEY:
		return par.parseDeclarativeStatement()
	case lexer.CONST_KEY:
		return par.parseDeclarativeStatement()
	// var a = (true && true);

	// for (a < 10)

	// {.....}
	case lexer.LEFT_BRACE:
		return par.parseBlockStatement()

	case lexer.IF_KEY:
		return par.parseIfStatement()

	case lexer.FUNC_KEY:
		return par.parseFunctionStatement()

	case lexer.FOR_KEY:
		return par.parseForLoop()

	case lexer.WHILE_KEY:
		return par.parseWhileLoop()

	case lexer.FOREACH_KEY:
		return par.parseForeachLoop()

	case lexer.STRUCT_KEY:
		return par.parseStructDeclaration()

	default:
		return par.parseExpression()
	}
}

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
	paren.Value = eval(par, paren.Expr)
	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

	return paren
}

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
	val, err := strconv.ParseInt(token.Literal, 10, 64)
	if err != nil {
		// try unsigned int for the edge case of -9223372036854775808
		// which is 9223372036854775808 in unsigned int
		// strconv.ParseInt fails for this value, but ParseUint succeeds
		// and we can cast it to int64
		uVal, uErr := strconv.ParseUint(token.Literal, 10, 64)
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
		Value: &objects.Integer{Value: val},
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
		Value: &objects.Float{Value: val},
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
		Value: &objects.Boolean{Value: token.Type == lexer.TRUE_KEY},
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
		Value: &objects.String{Value: par.CurrToken.Literal},
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
		Value: &objects.Nil{},
	}
}

// toFloat64 converts a GoMixObject to float64.
// This helper function is used for mixed-type arithmetic operations.
//
// Parameters:
//
//	obj - The object to convert (Integer or Float)
//
// Returns:
//
//	The float64 representation of the value, or 0 if not a number
func toFloat64(obj objects.GoMixObject) float64 {
	if obj.GetType() == objects.IntegerType {
		return float64(obj.(*objects.Integer).Value)
	} else if obj.GetType() == objects.FloatType {
		return obj.(*objects.Float).Value
	}
	return 0
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
	lVal := eval(par, left)
	// Evaluate the right operand
	rVal := eval(par, right)

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
func evaluateCompoundBinaryOp(lVal, rVal objects.GoMixObject, binaryOp lexer.TokenType) objects.GoMixObject {
	if lVal.GetType() == objects.IntegerType && rVal.GetType() == objects.IntegerType {
		l := lVal.(*objects.Integer).Value
		r := rVal.(*objects.Integer).Value
		switch binaryOp {
		case lexer.PLUS_OP:
			return &objects.Integer{Value: l + r}
		case lexer.MINUS_OP:
			return &objects.Integer{Value: l - r}
		case lexer.MUL_OP:
			return &objects.Integer{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				return &objects.Integer{Value: l / r}
			}
		case lexer.MOD_OP:
			if r != 0 {
				return &objects.Integer{Value: l % r}
			}
		case lexer.BIT_AND_OP:
			return &objects.Integer{Value: l & r}
		case lexer.BIT_OR_OP:
			return &objects.Integer{Value: l | r}
		case lexer.BIT_XOR_OP:
			return &objects.Integer{Value: l ^ r}
		case lexer.BIT_LEFT_OP:
			return &objects.Integer{Value: l << r}
		case lexer.BIT_RIGHT_OP:
			return &objects.Integer{Value: l >> r}
		}
	} else if (lVal.GetType() == objects.IntegerType || lVal.GetType() == objects.FloatType) &&
		(rVal.GetType() == objects.IntegerType || rVal.GetType() == objects.FloatType) {
		// Mixed arithmetic
		l := toFloat64(lVal)
		r := toFloat64(rVal)
		switch binaryOp {
		case lexer.PLUS_OP:
			return &objects.Float{Value: l + r}
		case lexer.MINUS_OP:
			return &objects.Float{Value: l - r}
		case lexer.MUL_OP:
			return &objects.Float{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				return &objects.Float{Value: l / r}
			}
		}
	}
	return &objects.Nil{}
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

	lVal := eval(par, left)
	rVal := eval(par, right)

	var val objects.GoMixObject = &objects.Nil{}

	if lVal.GetType() == objects.IntegerType && rVal.GetType() == objects.IntegerType {
		l := lVal.(*objects.Integer).Value
		r := rVal.(*objects.Integer).Value
		switch op.Type {
		case lexer.PLUS_OP:
			val = &objects.Integer{Value: l + r}
		case lexer.MINUS_OP:
			val = &objects.Integer{Value: l - r}
		case lexer.MUL_OP:
			val = &objects.Integer{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				val = &objects.Integer{Value: l / r}
			}
		case lexer.MOD_OP:
			if r != 0 {
				val = &objects.Integer{Value: l % r}
			}
		case lexer.BIT_AND_OP:
			val = &objects.Integer{Value: l & r}
		case lexer.BIT_OR_OP:
			val = &objects.Integer{Value: l | r}
		case lexer.BIT_XOR_OP:
			val = &objects.Integer{Value: l ^ r}
		case lexer.BIT_LEFT_OP:
			val = &objects.Integer{Value: l << r}
		case lexer.BIT_RIGHT_OP:
			val = &objects.Integer{Value: l >> r}
		}
	} else if (lVal.GetType() == objects.IntegerType || lVal.GetType() == objects.FloatType) &&
		(rVal.GetType() == objects.IntegerType || rVal.GetType() == objects.FloatType) {
		// Mixed arithmetic
		l := toFloat64(lVal)
		r := toFloat64(rVal)
		switch op.Type {
		case lexer.PLUS_OP:
			val = &objects.Float{Value: l + r}
		case lexer.MINUS_OP:
			val = &objects.Float{Value: l - r}
		case lexer.MUL_OP:
			val = &objects.Float{Value: l * r}
		case lexer.DIV_OP:
			if r != 0 {
				val = &objects.Float{Value: l / r}
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

	rVal := eval(par, right)
	var val objects.GoMixObject = &objects.Nil{}

	switch op.Type {
	case lexer.NOT_OP:
		// Logical NOT
		// Truthiness check
		isTrue := false
		if rVal.GetType() == objects.BooleanType {
			isTrue = rVal.(*objects.Boolean).Value
		} else if rVal.GetType() == objects.IntegerType {
			isTrue = rVal.(*objects.Integer).Value != 0
		}
		val = &objects.Boolean{Value: !isTrue}

	case lexer.MINUS_OP:
		if rVal.GetType() == objects.IntegerType {
			val = &objects.Integer{Value: -rVal.(*objects.Integer).Value}
		} else if rVal.GetType() == objects.FloatType {
			val = &objects.Float{Value: -rVal.(*objects.Float).Value}
		}
	case lexer.PLUS_OP:
		val = rVal
	case lexer.BIT_NOT_OP:
		if rVal.GetType() == objects.IntegerType {
			val = &objects.Integer{Value: ^rVal.(*objects.Integer).Value}
		}
	}

	return &UnaryExpressionNode{
		Operation: op,
		Right:     right,
		Value:     val,
	}
}

// parseDeclarativeStatement parses variable declaration statements.
// Handles var, let, and const declarations.
//
// Syntax:
//
//	var identifier = expression;
//	let identifier = expression;   (statically typed)
//	const identifier = expression; (immutable)
//
// Returns:
//
//	A DeclarativeStatementNode
//
// Behavior:
//   - var: Mutable, dynamically typed
//   - let: Mutable, statically typed (type inferred from initial value)
//   - const: Immutable, type checking enforced by evaluator
//
// Examples:
//
//	var x = 10;
//	let name = "Alice";
//	const PI = 3.14159;
func (par *Parser) parseDeclarativeStatement() StatementNode {
	varToken := par.CurrToken
	if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	identifier := par.CurrToken
	typ := "var"
	isLet := false
	if varToken.Type == lexer.CONST_KEY {
		typ = "const"
		par.Consts[identifier.Literal] = true
	} else if varToken.Type == lexer.LET_KEY {
		typ = "let"
		isLet = true
		par.LetVars[identifier.Literal] = true
	}
	if !par.expectAdvance(lexer.ASSIGN_OP) {
		return nil
	}
	par.advance()
	expr := par.parseExpression()
	if expr == nil {
		return nil
	}

	// evaluating the expression
	val := eval(par, expr)

	// save the type for let variables
	if varToken.Type == lexer.LET_KEY {
		par.LetTypes[identifier.Literal] = val.GetType()
	}

	// save the value in the environment
	par.Env[identifier.Literal] = val

	return &DeclarativeStatementNode{
		VarToken:   varToken,
		Identifier: IdentifierExpressionNode{Name: identifier.Literal, Value: val, Type: typ, IsLet: isLet},
		Expr:       expr,
		Value:      val,
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
		val = &objects.Nil{}
	}

	// Determine if this is a const or let or var
	ident := &IdentifierExpressionNode{
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

// parseReturnStatement parses return statements.
//
// Syntax:
//
//	return expression;
//
// Returns:
//
//	A ReturnStatementNode containing the return value expression
//
// Examples:
//
//	return 42;
//	return x + y;
//	return func() { return 5; }();
func (par *Parser) parseReturnStatement() ExpressionNode {
	returnToken := par.CurrToken
	par.advance()
	expr := par.parseExpression()
	if expr == nil {
		return nil
	}
	// evaluating the expression
	val := eval(par, expr)
	return &ReturnStatementNode{
		ReturnToken: returnToken,
		Expr:        expr,
		Value:       val,
	}
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

	lVal := eval(par, left)
	rVal := eval(par, right)
	val := false

	// Comparison logic
	if lVal.GetType() == objects.IntegerType && rVal.GetType() == objects.IntegerType {
		l := lVal.(*objects.Integer).Value
		r := rVal.(*objects.Integer).Value
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
		case lexer.AND_OP: // Logical AND/OR on integers should treat them as truthy/falsy
			val = (l != 0) && (r != 0)
		case lexer.OR_OP:
			val = (l != 0) || (r != 0)
		}
	} else if lVal.GetType() == objects.BooleanType && rVal.GetType() == objects.BooleanType {
		l := lVal.(*objects.Boolean).Value
		r := rVal.(*objects.Boolean).Value
		switch op.Type {
		case lexer.AND_OP:
			val = l && r
		case lexer.OR_OP:
			val = l || r
		case lexer.EQ_OP:
			val = l == r
		case lexer.NE_OP:
			val = l != r
		}
	} else if (lVal.GetType() == objects.FloatType || lVal.GetType() == objects.IntegerType) &&
		(rVal.GetType() == objects.FloatType || rVal.GetType() == objects.IntegerType) {
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
		}
	} else {
		// Fallback for other types, e.g., string comparison for equality
		switch op.Type {
		case lexer.EQ_OP:
			val = lVal.ToString() == rVal.ToString()
		case lexer.NE_OP:
			val = lVal.ToString() != rVal.ToString()
		case lexer.AND_OP: // Treat as truthy/falsy
			isLTrue := (lVal.GetType() == objects.BooleanType && lVal.(*objects.Boolean).Value) || (lVal.GetType() == objects.IntegerType && lVal.(*objects.Integer).Value != 0)
			isRTrue := (rVal.GetType() == objects.BooleanType && rVal.(*objects.Boolean).Value) || (rVal.GetType() == objects.IntegerType && rVal.(*objects.Integer).Value != 0)
			val = isLTrue && isRTrue
		case lexer.OR_OP: // Treat as truthy/falsy
			isLTrue := (lVal.GetType() == objects.BooleanType && lVal.(*objects.Boolean).Value) || (lVal.GetType() == objects.IntegerType && lVal.(*objects.Integer).Value != 0)
			isRTrue := (rVal.GetType() == objects.BooleanType && rVal.(*objects.Boolean).Value) || (rVal.GetType() == objects.IntegerType && rVal.(*objects.Integer).Value != 0)
			val = isLTrue || isRTrue
		}
	}

	return &BooleanExpressionNode{
		Operation: op,
		Left:      left,
		Right:     right,
		Value:     &objects.Boolean{Value: val},
	}
}

// parseBlockStatement parses block statements (code blocks).
// A block is a sequence of statements enclosed in curly braces.
//
// Syntax:
//
//	{ statement1; statement2; ... }
//
// Returns:
//
//	A BlockStatementNode containing all statements in the block
//
// The block's value is determined by the last statement in the block,
// allowing blocks to be used as expressions.
//
// Examples:
//
//	{ var x = 5; x + 10; }  - Block value is 15
//	{ println("Hello"); }   - Block value is nil
func (par *Parser) parseBlockStatement() *BlockStatementNode {
	block := &BlockStatementNode{}
	block.Statements = make([]StatementNode, 0)
	par.advance()
	for par.CurrToken.Type != lexer.RIGHT_BRACE && par.CurrToken.Type != lexer.EOF_TYPE {
		stmt := par.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		par.advance()
	}

	// computes the value of the block node
	// by evaluating the last statement
	if len(block.Statements) > 0 {
		lastStmt := block.Statements[len(block.Statements)-1]
		if exprNode, ok := lastStmt.(ExpressionNode); ok {
			block.Value = eval(par, exprNode)
		} else if declNode, ok := lastStmt.(*DeclarativeStatementNode); ok {
			block.Value = declNode.Value
		} else if returnNode, ok := lastStmt.(*ReturnStatementNode); ok {
			block.Value = returnNode.Value
		} else if blockNode, ok := lastStmt.(*BlockStatementNode); ok {
			block.Value = blockNode.Value
		} else if funcNode, ok := lastStmt.(*FunctionStatementNode); ok {
			block.Value = funcNode.Value
		} else if forLoopNode, ok := lastStmt.(*ForLoopStatementNode); ok {
			block.Value = forLoopNode.Value
		} else if whileLoopNode, ok := lastStmt.(*WhileLoopStatementNode); ok {
			block.Value = whileLoopNode.Value
		} else {
			block.Value = &objects.Nil{}
		}
	} else {
		block.Value = &objects.Nil{} // Default value for an empty block
	}

	return block
}

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

	if !isIdent && !isIndex {
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
		val := eval(par, right)
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
		Value:     &objects.Nil{},
	}
}

// parseFunctionStatement parses named function declarations.
//
// Syntax:
//
//	func functionName(param1, param2, ...) { body }
//
// Returns:
//
//	A FunctionStatementNode
//
// Unlike parseFunctionAssignment (anonymous functions), this creates
// a named function that can be called by name.
//
// Examples:
//
//	func add(a, b) { return a + b; }
//	func greet() { println("Hello!"); }
func (par *Parser) parseFunctionStatement() StatementNode {
	funcNode := NewFunctionStatementNode()
	funcNode.FuncToken = par.CurrToken
	if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	funcNode.FuncName = IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: &objects.Nil{}, // Default value for identifier
	}
	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}

	// Handle empty parameters case
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		// First parameter
		if !par.expectAdvance(lexer.IDENTIFIER_ID) {
			return nil
		}
		funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
			Name:  par.CurrToken.Literal,
			Value: &objects.Nil{}, // Default value for identifier
		})

		// Subsequent parameters
		for par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Consume comma
			if !par.expectAdvance(lexer.IDENTIFIER_ID) {
				return nil
			}
			funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
				Name:  par.CurrToken.Literal,
				Value: &objects.Nil{}, // Default value for identifier
			})
		}
	}
	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}
	funcNode.FuncBody = *par.parseBlockStatement()
	funcNode.Value = funcNode.FuncBody.Value
	return funcNode
}

// parseCallExpression parses function call expressions.
//
// Syntax:
//
//	functionName(arg1, arg2, ...)
//	functionName()  (no arguments)
//
// Returns:
//
//	A CallExpressionNode containing the function name and arguments
//
// Examples:
//
//	print("Hello")
//	add(5, 3)
//	factorial(10)
func (par *Parser) parseCallExpression() ExpressionNode {
	callNode := &CallExpressionNode{
		Value: &objects.Nil{},
	}
	callNode.FunctionIdentifier = IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: &objects.Nil{}, // Default value for identifier
	}

	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}
	// if there are arguments, parse them
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		par.advance()
		for {
			arg := par.parseExpression()
			if arg == nil {
				return nil
			}
			callNode.Arguments = append(callNode.Arguments, arg)
			if par.NextToken.Type == lexer.COMMA_DELIM {
				par.advance()
				par.advance()
			} else {
				break
			}
		}
	}

	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}
	return callNode
}

// parseFunctionAssignment parses anonymous function expressions.
// These are function literals that can be assigned to variables or passed as arguments.
//
// Syntax:
//
//	func(param1, param2, ...) { body }
//
// Returns:
//
//	A FunctionStatementNode representing the function expression
//
// Examples:
//
//	var add = func(a, b) { return a + b; };
//	var greet = func() { println("Hello!"); };
//	map(arr, func(x) { return x * 2; });
func (par *Parser) parseFunctionAssignment() ExpressionNode {
	funcNode := NewFunctionStatementNode()
	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}

	// Handle empty parameters case
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		// First parameter
		if !par.expectAdvance(lexer.IDENTIFIER_ID) {
			return nil
		}
		funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
			Name:  par.CurrToken.Literal,
			Value: &objects.Nil{}, // Default value for identifier
		})

		// Subsequent parameters
		for par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Consume comma
			if !par.expectAdvance(lexer.IDENTIFIER_ID) {
				return nil
			}
			funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
				Name:  par.CurrToken.Literal,
				Value: &objects.Nil{}, // Default value for identifier
			})
		}
	}
	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}
	funcNode.FuncBody = *par.parseBlockStatement()
	funcNode.Value = funcNode.FuncBody.Value
	return funcNode
}

// parseIfStatement parses if statements with optional else/else-if clauses.
//
// Syntax:
//
//	if (condition) { thenBlock }
//	if (condition) { thenBlock } else { elseBlock }
//	if (condition) { thenBlock } else if (condition2) { elseBlock }
//
// Returns:
//
//	An IfExpressionNode (despite the name, it's used as a statement)
//
// The function handles chained else-if by treating them as nested if statements
// within the else block.
//
// Examples:
//
//	if (x > 0) { println("positive"); }
//	if (x > 0) { println("positive"); } else { println("non-positive"); }
func (par *Parser) parseIfStatement() ExpressionNode {
	ifNode := NewIfStatement()
	ifNode.IfToken = par.CurrToken
	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}
	ifNode.Condition = par.parseInternal(MINIMUM_PRIORITY)
	if ifNode.Condition == nil {
		return nil
	}
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}
	ifNode.ThenBlock = *par.parseBlockStatement()
	if par.NextToken.Type == lexer.ELSE_KEY {
		par.advance() // consume closing brace of if block
		par.advance() // consume else
		if par.CurrToken.Type == lexer.IF_KEY {
			// else if case
			// treat it as a nested if statement
			// wrap it in a block statement
			elseBlock := &BlockStatementNode{}
			elseBlock.Statements = make([]StatementNode, 0)
			nestedIf := par.parseIfStatement()
			if nestedIf == nil {
				return nil
			}
			elseBlock.Statements = append(elseBlock.Statements, nestedIf)
			if exprNode, ok := nestedIf.(ExpressionNode); ok {
				elseBlock.Value = eval(par, exprNode)
			} else if stmtNode, ok := nestedIf.(StatementNode); ok {
				// If it's a statement, its value might be in its Value field
				if declNode, ok := stmtNode.(*DeclarativeStatementNode); ok {
					elseBlock.Value = declNode.Value
				} else if returnNode, ok := stmtNode.(*ReturnStatementNode); ok {
					elseBlock.Value = returnNode.Value
				} else if blockNode, ok := stmtNode.(*BlockStatementNode); ok {
					elseBlock.Value = blockNode.Value
				} else if funcNode, ok := stmtNode.(*FunctionStatementNode); ok {
					elseBlock.Value = funcNode.Value
				} else {
					elseBlock.Value = &objects.Nil{}
				}
			} else {
				elseBlock.Value = &objects.Nil{}
			}
			ifNode.ElseBlock = *elseBlock
		} else {
			ifNode.ElseBlock = *par.parseBlockStatement()
		}
	} else {
		ifNode.ElseBlock = BlockStatementNode{Value: &objects.Nil{}} // Default empty else block value
	}
	return ifNode
}

// parseIfExpression parses if expressions (used internally).
// This is similar to parseIfStatement but returns an expression node.
//
// Syntax:
//
//	if (condition) { thenBlock } else { elseBlock }
//
// Returns:
//
//	An IfExpressionNode
//
// Note: The else block is optional and defaults to an empty block.
func (par *Parser) parseIfExpression() ExpressionNode {
	ifToken := par.CurrToken
	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}
	par.advance()
	condition := par.parseExpression()
	if condition == nil {
		return nil
	}
	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}
	thenBlock := par.parseBlockStatement()

	var elseBlock *BlockStatementNode
	if par.CurrToken.Type == lexer.ELSE_KEY {
		par.advance()
		if !par.expectAdvance(lexer.LEFT_BRACE) {
			return nil
		}
		elseBlock = par.parseBlockStatement()
	} else {
		elseBlock = &BlockStatementNode{} // Default empty block
	}

	return &IfExpressionNode{
		IfToken:        ifToken,
		Condition:      condition,
		ThenBlock:      *thenBlock,
		ElseBlock:      *elseBlock,
		ConditionValue: &objects.Nil{},
	}
}

// parseForLoop parses for loop statements.
// Go-Mix for loops follow C-style syntax with three parts.
//
// Syntax:
//
//	for (initializer; condition; update) { body }
//
// Returns:
//
//	A ForLoopStatementNode
//
// Features:
//   - Multiple initializers (comma-separated)
//   - Variable declarations in initializer (var, let, const)
//   - Multiple updates (comma-separated)
//   - Optional condition (infinite loop if omitted)
//
// Examples:
//
//	for (var i = 0; i < 10; i += 1) { println(i); }
//	for (var i = 0, j = 10; i < j; i += 1, j -= 1) { ... }
func (par *Parser) parseForLoop() StatementNode {
	forToken := par.CurrToken

	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}

	// Parse initializers (can be var/const declarations or assignment expressions)
	initializers := make([]StatementNode, 0)

	// Parse first initializer if present
	if par.NextToken.Type != lexer.SEMICOLON_DELIM {
		par.advance()

		// Check if this is a variable declaration (var, let, or const)
		if par.CurrToken.Type == lexer.VAR_KEY || par.CurrToken.Type == lexer.LET_KEY || par.CurrToken.Type == lexer.CONST_KEY {
			// Parse variable declaration(s)
			varToken := par.CurrToken

			// Parse first variable declaration
			if !par.expectAdvance(lexer.IDENTIFIER_ID) {
				return nil
			}
			identifier := par.CurrToken
			typ := "var"
			isLet := false
			if varToken.Type == lexer.CONST_KEY {
				typ = "const"
				par.Consts[identifier.Literal] = true
			} else if varToken.Type == lexer.LET_KEY {
				typ = "let"
				isLet = true
				par.LetVars[identifier.Literal] = true
			}

			if !par.expectAdvance(lexer.ASSIGN_OP) {
				return nil
			}
			par.advance()
			expr := par.parseExpression()
			if expr == nil {
				return nil
			}

			// Evaluate and store in environment
			val := eval(par, expr)
			par.Env[identifier.Literal] = val

			// Save the type for let variables
			if varToken.Type == lexer.LET_KEY {
				par.LetTypes[identifier.Literal] = val.GetType()
			}

			declStmt := &DeclarativeStatementNode{
				VarToken:   varToken,
				Identifier: IdentifierExpressionNode{Name: identifier.Literal, Value: val, Type: typ, IsLet: isLet},
				Expr:       expr,
				Value:      val,
			}
			initializers = append(initializers, declStmt)

			// Parse additional variable declarations separated by commas
			for par.NextToken.Type == lexer.COMMA_DELIM {
				par.advance() // consume comma

				if !par.expectAdvance(lexer.IDENTIFIER_ID) {
					return nil
				}
				identifier := par.CurrToken
				typ := "var"
				isLet := false
				if varToken.Type == lexer.CONST_KEY {
					typ = "const"
					par.Consts[identifier.Literal] = true
				} else if varToken.Type == lexer.LET_KEY {
					typ = "let"
					isLet = true
					par.LetVars[identifier.Literal] = true
				}

				if !par.expectAdvance(lexer.ASSIGN_OP) {
					return nil
				}
				par.advance()
				expr := par.parseExpression()
				if expr == nil {
					return nil
				}

				// Evaluate and store in environment
				val := eval(par, expr)
				par.Env[identifier.Literal] = val

				// Save the type for let variables
				if varToken.Type == lexer.LET_KEY {
					par.LetTypes[identifier.Literal] = val.GetType()
				}

				declStmt := &DeclarativeStatementNode{
					VarToken:   varToken,
					Identifier: IdentifierExpressionNode{Name: identifier.Literal, Value: val, Type: typ, IsLet: isLet},
					Expr:       expr,
					Value:      val,
				}
				initializers = append(initializers, declStmt)
			}
		} else {
			// Parse regular expression initializer
			init := par.parseExpression()
			if init == nil {
				return nil
			}
			initializers = append(initializers, init)

			// Parse additional initializers separated by commas
			for par.NextToken.Type == lexer.COMMA_DELIM {
				par.advance() // consume comma
				par.advance() // move to next expression
				init := par.parseExpression()
				if init == nil {
					return nil
				}
				initializers = append(initializers, init)
			}
		}
	}

	if !par.expectAdvance(lexer.SEMICOLON_DELIM) {
		return nil
	}

	// Parse condition if present
	var condition ExpressionNode
	if par.NextToken.Type != lexer.SEMICOLON_DELIM {
		par.advance()
		condition = par.parseExpression()
		if condition == nil {
			return nil
		}
	}

	if !par.expectAdvance(lexer.SEMICOLON_DELIM) {
		return nil
	}

	// Parse updates (comma-separated assignment expressions)
	updates := make([]ExpressionNode, 0)
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		par.advance()
		update := par.parseExpression()
		if update == nil {
			return nil
		}
		updates = append(updates, update)

		// Parse additional updates separated by commas
		for par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // consume comma
			par.advance() // move to next expression
			update := par.parseExpression()
			if update == nil {
				return nil
			}
			updates = append(updates, update)
		}
	}

	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	body := par.parseBlockStatement()

	return &ForLoopStatementNode{
		ForToken:     forToken,
		Initializers: initializers,
		Condition:    condition,
		Updates:      updates,
		Body:         *body,
		Value:        &objects.Nil{},
	}
}

// parseWhileLoop parses while loop statements.
//
// Syntax:
//
//	while (condition) { body }
//	while (condition1, condition2, ...) { body }  (multiple conditions)
//
// Returns:
//
//	A WhileLoopStatementNode
//
// Multiple conditions are evaluated as a logical AND.
//
// Examples:
//
//	while (x < 10) { x += 1; }
//	while (x < 10, y > 0) { x += 1; y -= 1; }
func (par *Parser) parseWhileLoop() StatementNode {
	whileToken := par.CurrToken

	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}

	// Parse conditions (comma-separated)
	conditions := make([]ExpressionNode, 0)

	// Parse first condition
	par.advance()
	condition := par.parseExpression()
	if condition == nil {
		return nil
	}
	conditions = append(conditions, condition)

	// Parse additional conditions separated by commas
	for par.NextToken.Type == lexer.COMMA_DELIM {
		par.advance() // consume comma
		par.advance() // move to next expression
		condition := par.parseExpression()
		if condition == nil {
			return nil
		}
		conditions = append(conditions, condition)
	}

	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	body := par.parseBlockStatement()

	return &WhileLoopStatementNode{
		WhileToken: whileToken,
		Conditions: conditions,
		Body:       *body,
		Value:      &objects.Nil{},
	}
}

// parseArrayExpressionNode parses array literal expressions.
// Array literals are enclosed in square brackets with comma-separated elements.
//
// Syntax:
//
//	[element1, element2, element3, ...]
//	[]  (empty array)
//
// Returns:
//
//	An ArrayExpressionNode containing all parsed elements
//
// Examples:
//
//	[1, 2, 3]
//	["hello", "world"]
//	[1 + 2, 3 * 4, func() { return 5; }()]
func (par *Parser) parseArrayExpressionNode() ExpressionNode {
	arrayNode := &ArrayExpressionNode{}
	arrayElements := make([]ExpressionNode, 0)
	arrayNode.Elements = arrayElements

	// current token must be [
	if par.CurrToken.Type != lexer.LEFT_BRACKET {
		return nil
	}
	par.advance()
	if par.CurrToken.Type == lexer.RIGHT_BRACKET {
		return arrayNode
	}
	for par.CurrToken.Type != lexer.RIGHT_BRACKET {
		expr := par.parseExpression()
		arrayNode.Elements = append(arrayNode.Elements, expr)
		// After parsing expression, check if next token is ] or ,
		if par.NextToken.Type == lexer.RIGHT_BRACKET {
			par.advance() // move to ]
			break
		}
		if par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // move to ,
			par.advance() // move past , to next element
		} else {
			// If next token is neither ] nor ,, report error and try to continue
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected , or ], got %s",
				par.NextToken.Line, par.NextToken.Column, par.NextToken.Type)
			par.addError(msg)
			par.advance()
		}
	}
	return arrayNode
}

// parseIndexExpression parses array indexing and slicing operations.
// This function handles three cases:
// 1. Regular indexing: arr[index]
// 2. Slicing with start and end: arr[start:end]
// 3. Slicing with omitted bounds: arr[:end], arr[start:], arr[:]
//
// Parameters:
//
//	left - The array expression being indexed/sliced
//
// Returns:
//
//	Either an IndexExpressionNode or SliceExpressionNode
//
// Examples:
//
//	arr[0]      - Get first element
//	arr[-1]     - Get last element (negative indexing)
//	arr[1:3]    - Slice from index 1 to 3 (exclusive)
//	arr[:3]     - Slice from start to index 3
//	arr[1:]     - Slice from index 1 to end
//	arr[:]      - Copy entire array
func (par *Parser) parseIndexExpression(left ExpressionNode) ExpressionNode {
	// current token is [
	par.advance() // move past [

	// Check for empty slice [:] or [:end]
	if par.CurrToken.Type == lexer.COLON_DELIM {
		// This is a slice with no start: arr[:end] or arr[:]
		sliceNode := &SliceExpressionNode{
			Left:  left,
			Start: nil,
		}
		par.advance() // move past :

		// Check if there's an end index
		if par.CurrToken.Type != lexer.RIGHT_BRACKET {
			sliceNode.End = par.parseExpression()
			if sliceNode.End == nil {
				return nil
			}
			// After parsing end expression, NextToken should be ]
			if !par.expectAdvance(lexer.RIGHT_BRACKET) {
				return nil
			}
		} else {
			// CurrToken is already ], don't need to advance
			// The calling code in parseInternal will handle advancing
		}
		return sliceNode
	}

	// Parse the first expression (could be index or start of slice)
	firstExpr := par.parseExpression()
	if firstExpr == nil {
		return nil
	}

	// After parseExpression, check NextToken for colon (since parseExpression stops before operators it doesn't handle)
	if par.NextToken.Type == lexer.COLON_DELIM {
		// This is a slice: arr[start:end] or arr[start:]
		sliceNode := &SliceExpressionNode{
			Left:  left,
			Start: firstExpr,
		}
		par.advance() // move to :
		par.advance() // move past :

		// Check if there's an end index (skip any semicolons)
		for par.CurrToken.Type == lexer.SEMICOLON_DELIM {
			par.advance()
		}

		if par.CurrToken.Type != lexer.RIGHT_BRACKET {
			// Parse end expression
			sliceNode.End = par.parseExpression()
			if sliceNode.End == nil {
				return nil
			}
			// After parseExpression, NextToken should be ]
			if !par.expectAdvance(lexer.RIGHT_BRACKET) {
				return nil
			}
		} else {
			// CurrToken is already ], don't need to advance
			// The calling code in parseInternal will handle advancing
		}
		return sliceNode
	}

	// This is a regular index expression
	indexNode := &IndexExpressionNode{
		Left:  left,
		Index: firstExpr,
	}

	if !par.expectAdvance(lexer.RIGHT_BRACKET) {
		return nil
	}
	return indexNode
}

// parseInternal is the core of the Pratt parsing algorithm.
// It parses expressions while respecting operator precedence.
//
// Parameters:
//
//	currPrecedence - The minimum precedence level for operators to parse
//
// Returns:
//
//	An ExpressionNode representing the parsed expression
//
// Algorithm:
//  1. Parse a prefix expression (unary operator or primary expression)
//  2. While the next operator has higher precedence than currPrecedence:
//     a. Parse the operator as an infix expression
//     b. The result becomes the new left operand
//  3. Return the final expression
//
// This elegant algorithm handles operator precedence and associativity
// without needing separate precedence levels or recursive descent for each level.
func (par *Parser) parseInternal(currPrecedence int) ExpressionNode {
	unary, has := par.UnaryFuncs[par.CurrToken.Type]
	if !has {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected token: %s",
			par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
		par.addError(msg)
		return nil
	}
	left := unary()
	if left == nil {
		return nil
	}
	for par.NextToken.Type != lexer.EOF_TYPE && getPrecedence(&par.NextToken) >= currPrecedence {
		binary, has := par.BinaryFuncs[par.NextToken.Type]
		par.advance()
		if !has {
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected operator: %s",
				par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
			par.addError(msg)
			return nil
		}
		left = binary(left)
		if left == nil {
			return nil
		}
	}
	return left
}

// parseRangeExpression parses range expressions with the ... operator.
// Range expressions create inclusive ranges from start to end.
//
// Parameters:
//
//	left - The already-parsed left operand (start of range)
//
// Returns:
//
//	A RangeExpressionNode representing the range
//
// Syntax:
//
//	start...end  (creates range from start to end, inclusive)
//
// Examples:
//
//	2...5    - Range from 2 to 5 (inclusive)
//	1...10   - Range from 1 to 10 (inclusive)
//	x...y    - Range from x to y (inclusive)
func (par *Parser) parseRangeExpression(left ExpressionNode) ExpressionNode {
	// Current token is RANGE_OP (...)
	par.advance() // Move past ...

	// Parse the right operand (end of range)
	right := par.parseInternal(getPrecedence(&lexer.Token{Type: lexer.RANGE_OP}) + 1)
	if right == nil {
		return nil
	}

	// Evaluate both operands
	startVal := eval(par, left)
	endVal := eval(par, right)

	// Create the range value (will be nil if not both integers)
	var rangeVal objects.GoMixObject = &objects.Nil{}

	// Check if both are integers
	if startVal.GetType() == objects.IntegerType && endVal.GetType() == objects.IntegerType {
		start := startVal.(*objects.Integer).Value
		end := endVal.(*objects.Integer).Value
		rangeVal = &objects.Range{Start: start, End: end}
	}

	return &RangeExpressionNode{
		Start: left,
		End:   right,
		Value: rangeVal,
	}
}

// parseMapLiteral parses map literal expressions.
// Map literals use the syntax: map{key1: value1, key2: value2, ...}
//
// Syntax:
//
//	map{key: value, key: value, ...}
//	map{}  (empty map)
//
// Returns:
//
//	A MapExpressionNode containing all parsed key-value pairs
//
// Examples:
//
//	map{10: 20, 30: 40}
//	map{"name": "John", "age": 25}
//	map{1: "one", 2: "two", 3: "three"}
func (par *Parser) parseMapLiteral() ExpressionNode {
	mapNode := &MapExpressionNode{
		Keys:   make([]ExpressionNode, 0),
		Values: make([]ExpressionNode, 0),
	}

	// Current token is MAP_KEY
	// Expect opening brace
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	// Check for empty map
	if par.NextToken.Type == lexer.RIGHT_BRACE {
		par.advance() // Move to }
		return mapNode
	}

	// Parse key-value pairs
	par.advance() // Move to first key
	for {
		// Parse key expression
		key := par.parseExpression()
		if key == nil {
			return nil
		}

		// Expect colon
		if !par.expectAdvance(lexer.COLON_DELIM) {
			return nil
		}

		// Parse value expression
		par.advance() // Move past colon
		value := par.parseExpression()
		if value == nil {
			return nil
		}

		// Add key-value pair
		mapNode.Keys = append(mapNode.Keys, key)
		mapNode.Values = append(mapNode.Values, value)

		// Check what comes next
		if par.NextToken.Type == lexer.RIGHT_BRACE {
			par.advance() // Move to }
			break
		}

		if par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Move to ,
			par.advance() // Move past , to next key
		} else {
			// Error: expected , or }
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected , or }, got %s",
				par.NextToken.Line, par.NextToken.Column, par.NextToken.Type)
			par.addError(msg)
			return nil
		}
	}

	return mapNode
}

// parseSetLiteral parses set literal expressions.
// Set literals use the syntax: set{value1, value2, value3, ...}
// Sets automatically remove duplicates and maintain unique values.
//
// Syntax:
//
//	set{value, value, value, ...}
//	set{}  (empty set)
//
// Returns:
//
//	A SetExpressionNode containing all parsed element expressions
//
// Examples:
//
//	set{1, 2, 3, 4, 5}
//	set{"apple", "banana", "cherry"}
//	set{1, 2, 2, 3}  // Duplicates will be removed during evaluation
func (par *Parser) parseSetLiteral() ExpressionNode {
	setNode := &SetExpressionNode{
		Elements: make([]ExpressionNode, 0),
	}

	// Current token is SET_KEY
	// Expect opening brace
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	// Check for empty set
	if par.NextToken.Type == lexer.RIGHT_BRACE {
		par.advance() // Move to }
		return setNode
	}

	// Parse elements
	par.advance() // Move to first element
	for {
		// Parse element expression
		elem := par.parseExpression()
		if elem == nil {
			return nil
		}

		// Add element
		setNode.Elements = append(setNode.Elements, elem)

		// Check what comes next
		if par.NextToken.Type == lexer.RIGHT_BRACE {
			par.advance() // Move to }
			break
		}

		if par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Move to ,
			par.advance() // Move past , to next element
		} else {
			// Error: expected , or }
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected , or }, got %s",
				par.NextToken.Line, par.NextToken.Column, par.NextToken.Type)
			par.addError(msg)
			return nil
		}
	}

	return setNode
}

// parseForeachLoop parses foreach loop statements.
// Foreach loops iterate over ranges or arrays.
//
// Syntax:
//
//	foreach identifier in iterable { body }
//
// Returns:
//
//	A ForeachLoopStatementNode
//
// Examples:
//
//	foreach i in 2...10 { print(i); }
//	foreach item in array { print(item); }
//	foreach x in myRange { body }
func (par *Parser) parseForeachLoop() StatementNode {
	foreachToken := par.CurrToken

	// Expect iterator identifier
	if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	iterator := IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: &objects.Nil{},
	}

	// Expect 'in' keyword
	if !par.expectAdvance(lexer.IN_KEY) {
		return nil
	}

	// Parse the iterable expression (range or array)
	par.advance()
	iterable := par.parseExpression()
	if iterable == nil {
		return nil
	}

	// Expect opening brace for body
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	// Parse the loop body
	body := par.parseBlockStatement()

	return &ForeachLoopStatementNode{
		ForeachToken: foreachToken,
		Iterator:     iterator,
		Iterable:     iterable,
		Body:         *body,
		Value:        &objects.Nil{},
	}
}

// parseStructDeclaration parses struct declarations.
//
// Syntax:
func (par *Parser) parseStructDeclaration() StatementNode {
	structToken := par.CurrToken

	// Expect struct name
	if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	structName := IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: &objects.Nil{},
	}

	// Expect opening brace for struct body
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	// Parse struct methods
	methods := make([]*FunctionStatementNode, 0)
	for par.NextToken.Type != lexer.RIGHT_BRACE {
		par.advance()
		if par.CurrToken.Type == lexer.FUNC_KEY {
			method := par.parseFunctionStatement()
			if method == nil {
				return nil
			}
			methods = append(methods, method.(*FunctionStatementNode))
		} else {
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected 'func' in struct body, got %s",
				par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Type)
			par.addError(msg)
			return nil
		}
	}

	// Expect closing brace for struct body
	if !par.expectAdvance(lexer.RIGHT_BRACE) {
		return nil
	}

	return &StructDeclarationNode{
		StructToken: structToken,
		StructName:  structName,
		Methods:     methods,
		Value:       &objects.Nil{},
	}
}

// parseNewCallExpression parses expressions for creating new instances of structs.
//
// Syntax:
//
//	new StructName(arg1, arg2, ...)
//
// Returns:
//
//	A NewCallExpressionNode representing the instantiation of a struct
//
// Examples:
//
//	new MyStruct(10, "hello")
//	new Point(5, 5)
func (par *Parser) parseNewCallExpression() ExpressionNode {
	// Current token is NEW_KEY

	newCallNode := &NewCallExpressionNode{
		NewToken: par.CurrToken,
		Value:    &objects.Nil{},
	}
	if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	newCallNode.StructName = IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: &objects.Nil{}, // Default value for identifier
	}

	if !par.expectAdvance(lexer.LEFT_PAREN) {
		return nil
	}
	// if there are arguments, parse them
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		par.advance()
		for {
			arg := par.parseExpression()
			if arg == nil {
				return nil
			}
			newCallNode.Arguments = append(newCallNode.Arguments, arg)
			if par.NextToken.Type == lexer.COMMA_DELIM {
				par.advance()
				par.advance()
			} else {
				break
			}
		}
	}

	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}
	return newCallNode
}

// eval evaluates an expression node during parsing.
// This enables constant folding and early error detection.
//
// Parameters:
//
//	par  - The parser instance (for accessing environment)
//	node - The expression node to evaluate
//
// Returns:
//
//	The evaluated GoMixObject value
//
// Note: This is a simplified evaluator used during parsing.
// The full evaluator in the eval package handles more complex cases.
func eval(par *Parser, node ExpressionNode) objects.GoMixObject {
	switch n := node.(type) {
	case *IntegerLiteralExpressionNode:
		return n.Value
	case *BinaryExpressionNode:
		return n.Value
	case *UnaryExpressionNode:
		return n.Value
	case *BooleanLiteralExpressionNode:
		return n.Value
	case *ParenthesizedExpressionNode:
		return eval(par, n.Expr)
	case *IdentifierExpressionNode:
		if val, ok := par.Env[n.Name]; ok {
			return val
		}
		return &objects.Nil{}
	case *ReturnStatementNode:
		return eval(par, n.Expr)
	case *BooleanExpressionNode:
		return n.Value
	case *BlockStatementNode:
		var result objects.GoMixObject = &objects.Nil{}
		for _, stmt := range n.Statements {
			if expr, ok := stmt.(ExpressionNode); ok {
				val := eval(par, expr)
				if _, isRet := stmt.(*ReturnStatementNode); isRet {
					return val
				}
				result = val
			}
		}
		return result
	case *AssignmentExpressionNode:
		return n.Value
	case *IfExpressionNode:
		cond := eval(par, n.Condition)
		n.ConditionValue = cond
		// Check truthiness
		var isTrue bool
		if cond.GetType() == objects.BooleanType {
			isTrue = cond.(*objects.Boolean).Value
		} else if cond.GetType() == objects.IntegerType {
			isTrue = cond.(*objects.Integer).Value != 0
		} else {
			isTrue = false
		}

		if isTrue {
			return eval(par, &n.ThenBlock)
		}
		return eval(par, &n.ElseBlock)
	case *StringLiteralExpressionNode:
		return n.Value
	case *FunctionStatementNode:
		return n.Value
	case *CallExpressionNode:
		return n.Value
	case *FloatLiteralExpressionNode:
		return n.Value
	}
	return &objects.Nil{}
}
