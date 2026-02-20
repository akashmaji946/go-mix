package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

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
	val := parseEval(par, expr)

	// save the type for let variables
	if varToken.Type == lexer.LET_KEY {
		par.LetTypes[identifier.Literal] = val.GetType()
	}

	// save the value in the environment
	par.Env[identifier.Literal] = val

	return &DeclarativeStatementNode{
		VarToken:   varToken,
		Identifier: IdentifierExpressionNode{Token: identifier, Name: identifier.Literal, Value: val, Type: typ, IsLet: isLet},
		Expr:       expr,
		Value:      val,
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
			block.Value = parseEval(par, exprNode)
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
			block.Value = &std.Nil{}
		}
	} else {
		block.Value = &std.Nil{} // Default value for an empty block
	}

	return block
}
