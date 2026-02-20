/*
File    : go-mix/parser/switch_parser.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
)

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

	return body
}
