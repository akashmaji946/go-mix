package parser

import (
	"fmt"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

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
	startVal := parseEval(par, left)
	endVal := parseEval(par, right)

	// Create the range value (will be nil if not both integers)
	var rangeVal std.GoMixObject = &std.Nil{}

	// Check if both are integers
	if startVal.GetType() == std.IntegerType && endVal.GetType() == std.IntegerType {
		start := startVal.(*std.Integer).Value
		end := endVal.(*std.Integer).Value
		rangeVal = &std.Range{Start: start, End: end}
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

// parseMapKeyword dispatches between map literals and function calls.
func (par *Parser) parseMapKeyword() ExpressionNode {
	if par.NextToken.Type == lexer.LEFT_PAREN {
		return par.parseCallExpression()
	}
	return par.parseMapLiteral()
}

// parseSetKeyword dispatches between set literals and function calls.
func (par *Parser) parseSetKeyword() ExpressionNode {
	if par.NextToken.Type == lexer.LEFT_PAREN {
		return par.parseCallExpression()
	}
	return par.parseSetLiteral()
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
