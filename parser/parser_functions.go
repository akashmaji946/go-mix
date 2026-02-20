package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

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

	// Allow identifiers and specific keywords as function names
	if par.NextToken.Type == lexer.SET_KEY || par.NextToken.Type == lexer.MAP_KEY || par.NextToken.Type == lexer.ARRAY_KEY {
		par.advance()
	} else if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	funcNode.FuncName = IdentifierExpressionNode{
		Token: par.CurrToken,
		Name:  par.CurrToken.Literal,
		Value: &std.Nil{}, // Default value for identifier
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
			Token: par.CurrToken,
			Name:  par.CurrToken.Literal,
			Value: &std.Nil{}, // Default value for identifier
		})

		// Subsequent parameters
		for par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Consume comma
			if !par.expectAdvance(lexer.IDENTIFIER_ID) {
				return nil
			}
			funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
				Token: par.CurrToken,
				Name:  par.CurrToken.Literal,
				Value: &std.Nil{}, // Default value for identifier
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
		Value: &std.Nil{},
	}
	callNode.FunctionIdentifier = IdentifierExpressionNode{
		Token: par.CurrToken,
		Name:  par.CurrToken.Literal,
		Value: &std.Nil{}, // Default value for identifier
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
			Token: par.CurrToken,
			Name:  par.CurrToken.Literal,
			Value: &std.Nil{}, // Default value for identifier
		})

		// Subsequent parameters
		for par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Consume comma
			if !par.expectAdvance(lexer.IDENTIFIER_ID) {
				return nil
			}
			funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
				Token: par.CurrToken,
				Name:  par.CurrToken.Literal,
				Value: &std.Nil{}, // Default value for identifier
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
