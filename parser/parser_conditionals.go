/*
File    : go-mix/parser/parser_conditionals.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

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
				elseBlock.Value = parseEval(par, exprNode)
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
					elseBlock.Value = &std.Nil{}
				}
			} else {
				elseBlock.Value = &std.Nil{}
			}
			ifNode.ElseBlock = *elseBlock
		} else {
			ifNode.ElseBlock = *par.parseBlockStatement()
		}
	} else {
		ifNode.ElseBlock = BlockStatementNode{Value: &std.Nil{}} // Default empty else block value
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
		ConditionValue: &std.Nil{},
	}
}

// parseSwitchStatement parses a switch statement.
// Syntax: switch (expression) { case value: statements... default: statements... }
func (par *Parser) parseSwitchStatement() StatementNode {
	// Create the switch statement node
	switchNode := &SwitchStatementNode{
		Token: par.CurrToken,
		Cases: make([]SwitchCaseNode, 0),
	}

	// Expect 'switch' keyword (current token)
	if par.CurrToken.Type != lexer.SWITCH_KEY {
		par.Errors = append(par.Errors, "Expected 'switch' keyword")
		return nil
	}

	// Move to next token (should be the expression or left paren)
	par.advance()

	// Optional parentheses around the expression
	hasParen := false
	if par.CurrToken.Type == lexer.LEFT_PAREN {
		hasParen = true
		par.advance()
	}

	// Parse the switch expression
	if par.CurrToken.Type == lexer.EOF_TYPE {
		par.Errors = append(par.Errors, "Unexpected end of file after 'switch'")
		return nil
	}

	switchNode.Expression = par.parseExpression()

	// If we had an opening paren, expect a closing paren
	if hasParen {
		if !par.expectAdvance(lexer.RIGHT_PAREN) {
			return nil
		}
	}

	// Expect opening brace
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}
	par.advance() // Move past the {

	// Parse case clauses and optional default clause
	for par.CurrToken.Type != lexer.RIGHT_BRACE && par.CurrToken.Type != lexer.EOF_TYPE {
		switch par.CurrToken.Type {
		case lexer.CASE_KEY:
			caseNode := par.parseCaseClause()
			if caseNode != nil {
				switchNode.Cases = append(switchNode.Cases, *caseNode)
			}

		case lexer.DEFAULT_KEY:
			if switchNode.Default != nil {
				par.Errors = append(par.Errors, "Switch statement can only have one default clause")
			}
			defaultNode := par.parseDefaultClause()
			switchNode.Default = defaultNode

		default:
			par.Errors = append(par.Errors, "Expected 'case' or 'default' in switch body, got "+string(par.CurrToken.Type))
			par.advance()
		}
	}

	// Expect closing brace
	if par.CurrToken.Type != lexer.RIGHT_BRACE {
		par.Errors = append(par.Errors, "Expected '}' to end switch body")
		return nil
	}

	return switchNode
}

// parseCaseClause parses a single case clause in a switch statement.
// Syntax: case expression: statements...
func (par *Parser) parseCaseClause() *SwitchCaseNode {
	caseNode := &SwitchCaseNode{
		Token: par.CurrToken,
	}

	// Expect 'case' keyword (current token)
	if par.CurrToken.Type != lexer.CASE_KEY {
		par.Errors = append(par.Errors, "Expected 'case' keyword")
		return nil
	}
	par.advance()

	// Parse the case value expression
	if par.CurrToken.Type == lexer.EOF_TYPE {
		par.Errors = append(par.Errors, "Unexpected end of file after 'case'")
		return nil
	}

	caseNode.Value = par.parseExpression()

	// Expect colon (current token should be colon after parseExpression)
	if !par.expectAdvance(lexer.COLON_DELIM) {
		return nil
	}
	par.advance() // Move past the :

	// Parse statements until we hit another case, default, or closing brace
	caseNode.Body = par.parseCaseBody()

	return caseNode
}

// parseDefaultClause parses the default clause in a switch statement.
// Syntax: default: statements...
func (par *Parser) parseDefaultClause() *SwitchDefaultNode {
	defaultNode := &SwitchDefaultNode{
		Token: par.CurrToken,
	}

	// Expect 'default' keyword (current token)
	if par.CurrToken.Type != lexer.DEFAULT_KEY {
		par.Errors = append(par.Errors, "Expected 'default' keyword")
		return nil
	}
	par.advance()

	// Expect colon
	if par.CurrToken.Type != lexer.COLON_DELIM {
		par.Errors = append(par.Errors, "Expected ':' after 'default'")
		return nil
	}
	par.advance()

	// Parse statements until we hit another case or closing brace
	defaultNode.Body = par.parseCaseBody()

	return defaultNode
}

// parseCaseBody parses the statements within a case or default clause.
// It stops when it encounters another case, default, or closing brace.
func (par *Parser) parseCaseBody() BlockStatementNode {
	body := BlockStatementNode{
		Statements: make([]StatementNode, 0),
	}

	// Parse statements until we hit another case, default, or closing brace
	for par.CurrToken.Type != lexer.CASE_KEY &&
		par.CurrToken.Type != lexer.DEFAULT_KEY &&
		par.CurrToken.Type != lexer.RIGHT_BRACE &&
		par.CurrToken.Type != lexer.EOF_TYPE {

		stmt := par.parseStatement()
		if stmt != nil {
			body.Statements = append(body.Statements, stmt)
		}

		// Advance to the next token to continue parsing the next statement.
		// This is crucial to prevent an infinite loop if parseStatement() doesn't
		// advance the token, ensuring the loop always makes progress.
		par.advance()
	}

	// computes the value of the block node
	// by evaluating the last statement
	if len(body.Statements) > 0 {
		lastStmt := body.Statements[len(body.Statements)-1]
		if exprNode, ok := lastStmt.(ExpressionNode); ok {
			body.Value = parseEval(par, exprNode)
		} else if declNode, ok := lastStmt.(*DeclarativeStatementNode); ok {
			body.Value = declNode.Value
		} else if returnNode, ok := lastStmt.(*ReturnStatementNode); ok {
			body.Value = returnNode.Value
		} else if blockNode, ok := lastStmt.(*BlockStatementNode); ok {
			body.Value = blockNode.Value
		} else if funcNode, ok := lastStmt.(*FunctionStatementNode); ok {
			body.Value = funcNode.Value
		} else if forLoopNode, ok := lastStmt.(*ForLoopStatementNode); ok {
			body.Value = forLoopNode.Value
		} else if whileLoopNode, ok := lastStmt.(*WhileLoopStatementNode); ok {
			body.Value = whileLoopNode.Value
		} else {
			body.Value = &std.Nil{}
		}
	} else {
		body.Value = &std.Nil{} // Default value for an empty block
	}
	return body
}
