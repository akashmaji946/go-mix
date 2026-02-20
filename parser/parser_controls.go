package parser

import (
	"fmt"
	"strings"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

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
//	 	return;
//		return 42;
//		return x + y;
//		return func() { return 5; }();
func (par *Parser) parseReturnStatement() ExpressionNode {
	returnToken := par.CurrToken
	par.advance()

	// Handle empty return (return;)
	if par.CurrToken.Type == lexer.SEMICOLON_DELIM {
		return &ReturnStatementNode{
			ReturnToken: returnToken,
			Expr:        &NilLiteralExpressionNode{Token: lexer.Token{Type: lexer.NIL_LIT, Literal: "nil"}, Value: &std.Nil{}},
			Value:       &std.Nil{},
		}
	}

	expr := par.parseExpression()
	if expr == nil {
		return nil
	}
	// evaluating the expression
	val := parseEval(par, expr)
	return &ReturnStatementNode{
		ReturnToken: returnToken,
		Expr:        expr,
		Value:       val,
	}
}

// parseBreakStatement parses a break statement.
func (par *Parser) parseBreakStatement() StatementNode {
	stmt := &BreakStatementNode{Token: par.CurrToken}
	return stmt
}

// parseContinueStatement parses a continue statement.
func (par *Parser) parseContinueStatement() StatementNode {
	stmt := &ContinueStatementNode{Token: par.CurrToken}
	return stmt
}

// parseImportStatement parses an import statement.
// Syntax: import packageName; or import "packageName";
func (par *Parser) parseImportStatement() StatementNode {
	importToken := par.CurrToken

	// Expect the package name (identifier or string literal)
	if par.NextToken.Type != lexer.IDENTIFIER_ID && par.NextToken.Type != lexer.STRING_LIT {
		par.addError(fmt.Sprintf("[%d:%d] PARSER ERROR: expected Identifier or StringLiteral, got %s",
			par.NextToken.Line, par.NextToken.Column, par.NextToken.Type))
		return nil
	}
	par.advance()
	packageName := par.CurrToken.Literal
	// If it's a string literal, remove the quotes
	if par.CurrToken.Type == lexer.STRING_LIT {
		packageName = strings.Trim(packageName, `"`)
	}
	alias := "" // Default: no alias

	// Check for optional "as" keyword for aliasing
	if par.NextToken.Type == lexer.IDENTIFIER_ID && par.NextToken.Literal == "as" {
		par.advance() // Move to "as"
		par.advance() // Move to alias identifier

		// Expect the alias name (identifier)
		if par.CurrToken.Type != lexer.IDENTIFIER_ID {
			par.addError(fmt.Sprintf("expected identifier for alias, got %s", par.CurrToken.Literal))
			return nil
		}
		alias = par.CurrToken.Literal
	}

	// Expect semicolon
	if !par.expectAdvance(lexer.SEMICOLON_DELIM) {
		return nil
	}

	return &ImportStatementNode{
		Token: importToken,
		Name:  packageName,
		Alias: alias,
	}
}
