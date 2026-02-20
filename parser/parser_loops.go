package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

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
			val := parseEval(par, expr)
			par.Env[identifier.Literal] = val

			// Save the type for let variables
			if varToken.Type == lexer.LET_KEY {
				par.LetTypes[identifier.Literal] = val.GetType()
			}

			declStmt := &DeclarativeStatementNode{
				VarToken:   varToken,
				Identifier: IdentifierExpressionNode{Token: identifier, Name: identifier.Literal, Value: val, Type: typ, IsLet: isLet},
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
				val := parseEval(par, expr)
				par.Env[identifier.Literal] = val

				// Save the type for let variables
				if varToken.Type == lexer.LET_KEY {
					par.LetTypes[identifier.Literal] = val.GetType()
				}

				declStmt := &DeclarativeStatementNode{
					VarToken:   varToken,
					Identifier: IdentifierExpressionNode{Token: identifier, Name: identifier.Literal, Value: val, Type: typ, IsLet: isLet},
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
		Value:        &std.Nil{},
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
		Value:      &std.Nil{},
	}
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
		Token: par.CurrToken,
		Name:  par.CurrToken.Literal,
		Value: &std.Nil{},
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
		Value:        &std.Nil{},
	}
}
