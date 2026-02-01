package parser

import (
	"fmt"
	"strconv"

	"github.com/akashmaji946/go-mix/lexer"
)

// operator precedence
const (
	MINIMUM_PRIORITY = 0

	ASSIGN_PRIORITY = 10

	OR_PRIORITY  = 20
	AND_PRIORITY = 30

	RELATIONAL_PRIORITY = 40

	PLUS_PRIORITY = 50 // + - | ^
	MUL_PRIORITY  = 60 // * / % << >> & &^

	PREFIX_PRIORITY = 70

	PAREN_PRIORITY = 100
)

// get the precedence of the operator
func getPrecedence(token *lexer.Token) int {
	switch token.Type {

	case lexer.LEFT_PAREN:
		return PAREN_PRIORITY
	case lexer.NOT_OP:
		return PREFIX_PRIORITY

	// Mul level: * / % & << >>
	case lexer.MUL_OP, lexer.DIV_OP, lexer.MOD_OP, lexer.BIT_AND_OP, lexer.BIT_LEFT_OP, lexer.BIT_RIGHT_OP:
		return MUL_PRIORITY

	// Add level: + - | ^
	case lexer.PLUS_OP, lexer.MINUS_OP, lexer.BIT_OR_OP, lexer.BIT_XOR_OP:
		return PLUS_PRIORITY

	case lexer.GT_OP, lexer.LT_OP, lexer.GE_OP, lexer.LE_OP, lexer.EQ_OP, lexer.NE_OP:
		return RELATIONAL_PRIORITY

	case lexer.AND_OP:
		return AND_PRIORITY
	case lexer.OR_OP:
		return OR_PRIORITY

	case lexer.ASSIGN_OP:
		return ASSIGN_PRIORITY

	default:
		return -1
	}
}

// binaryParseFunction(): parse binary expression
// first parameter is the parsed left expression
// parses the right expression and returns the parsed expression
type binaryParseFunction func(ExpressionNode) ExpressionNode

// unaryParseFunction(): parse unary expression
// parses the right expression and returns the parsed expression
type unaryParseFunction func() ExpressionNode

type Parser struct {
	Lex       lexer.Lexer
	CurrToken lexer.Token
	NextToken lexer.Token

	UnaryFuncs  map[lexer.TokenType]unaryParseFunction
	BinaryFuncs map[lexer.TokenType]binaryParseFunction

	Env map[string]int64
}

func NewParser(src string) *Parser {
	lex := lexer.NewLexer(src)
	par := &Parser{
		Lex: lex,
	}
	par.init()
	return par
}

func (par *Parser) init() {
	par.UnaryFuncs = make(map[lexer.TokenType]unaryParseFunction)
	par.BinaryFuncs = make(map[lexer.TokenType]binaryParseFunction)
	par.Env = make(map[string]int64)

	// register the functions
	par.registerUnaryFuncs(par.parseParenthesizedExpression, lexer.LEFT_PAREN)
	par.registerUnaryFuncs(par.parseNumberLiteral, lexer.INT_LIT)
	par.registerUnaryFuncs(par.parseFloatLiteral, lexer.FLOAT_LIT)
	par.registerUnaryFuncs(par.parseReturnStatement, lexer.RETURN_KEY)
	par.registerUnaryFuncs(par.parseIdentifierExpression, lexer.IDENTIFIER_ID)

	par.registerUnaryFuncs(par.parseBooleanLiteral, lexer.TRUE_KEY, lexer.FALSE_KEY)

	par.registerUnaryFuncs(par.parseStringLiteral, lexer.STRING_LIT)
	par.registerUnaryFuncs(par.parseNilLiteral, lexer.NIL_LIT)

	par.registerBinaryFuncs(par.parseBinaryExpression, lexer.PLUS_OP, lexer.MINUS_OP, lexer.MUL_OP, lexer.DIV_OP, lexer.MOD_OP)
	par.registerBinaryFuncs(par.parseBinaryExpression, lexer.BIT_AND_OP, lexer.BIT_OR_OP, lexer.BIT_XOR_OP, lexer.BIT_LEFT_OP, lexer.BIT_RIGHT_OP)

	par.registerUnaryFuncs(par.parseUnaryExpression, lexer.NOT_OP, lexer.MINUS_OP, lexer.BIT_NOT_OP)

	par.registerBinaryFuncs(par.parseBooleanExpression, lexer.AND_OP, lexer.OR_OP, lexer.GT_OP, lexer.LT_OP, lexer.GE_OP, lexer.LE_OP, lexer.EQ_OP, lexer.NE_OP)
	par.registerUnaryFuncs(par.parseUnaryExpression, lexer.NOT_OP, lexer.MINUS_OP)
	par.registerBinaryFuncs(par.parseAssignmentExpression, lexer.ASSIGN_OP)

	par.advance()
	par.advance()
}

// helper to register unary functions
func (par *Parser) registerUnaryFuncs(f unaryParseFunction, tokenTypes ...lexer.TokenType) {
	for _, tokenType := range tokenTypes {
		par.UnaryFuncs[tokenType] = f
	}
}

// helper to register binary functions
func (par *Parser) registerBinaryFuncs(f binaryParseFunction, tokenTypes ...lexer.TokenType) {
	for _, tokenType := range tokenTypes {
		par.BinaryFuncs[tokenType] = f
	}
}

// advance the parser to the next token
func (par *Parser) advance() {
	par.CurrToken = par.NextToken
	par.NextToken = par.Lex.NextToken()
}

// expect the next token to be of the expected type and advance
func (par *Parser) expectAdvance(expected lexer.TokenType) {
	if par.NextToken.Type != expected {
		msg := fmt.Sprintf("[ERROR] expected %s, got %s", expected, par.NextToken.Type)
		fmt.Println(msg)
		panic(msg)
	}
	par.advance()
}

// parse the source code
// parse each statement and return the root node
func (par *Parser) Parse() *RootNode {

	// a real parser
	root := &RootNode{}
	root.Statements = make([]StatementNode, 0)
	for par.CurrToken.Type != lexer.EOF_TYPE {
		stmt := par.parseStatement()
		if stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}
		par.advance()
	}

	// computes the value of the root node
	// by evaluating the last statement
	if len(root.Statements) > 0 {
		lastStmt := root.Statements[len(root.Statements)-1]
		if exprNode, ok := lastStmt.(ExpressionNode); ok {
			root.Value = eval(par, exprNode)
		} else if declNode, ok := lastStmt.(*DeclarativeStatementNode); ok {
			root.Value = declNode.Value
		} else if returnNode, ok := lastStmt.(*ReturnStatementNode); ok {
			root.Value = returnNode.Value
		} else if blockNode, ok := lastStmt.(*BlockStatementNode); ok {
			root.Value = blockNode.Value
		} else if funcNode, ok := lastStmt.(*FunctionStatementNode); ok {
			root.Value = funcNode.Value
		}
	}

	return root
}

// parse a statement
// currently only supports expression statements (will be extended)
func (par *Parser) parseStatement() StatementNode {
	switch par.CurrToken.Type {

	// ignore semicolons
	case lexer.SEMICOLON_DELIM:
		return nil

	// TODO: add for statements like

	// var a = 10;
	case lexer.VAR_KEY:
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

	default:
		return par.parseExpression()
	}
}

// parse an expression
// currently only supports binary expressions and unary expressions (will be extended)
func (par *Parser) parseExpression() ExpressionNode {

	return par.parseInternal(MINIMUM_PRIORITY)
}

// parse a parenthesis expression
func (par *Parser) parseParenthesizedExpression() ExpressionNode {
	// we are already at the LEFT_PAREN, so just advance
	par.advance()
	paren := &ParenthesizedExpressionNode{}
	paren.Expr = par.parseExpression()
	par.expectAdvance(lexer.RIGHT_PAREN)
	// par.advance()
	return paren
}

// parse a number literal
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
			msg := fmt.Sprintf("[ERROR] could not parse number literal: %s", token.Literal)
			fmt.Println(msg)
			panic(msg)
		}
	}
	return &IntegerLiteralExpressionNode{
		Token: token,
		Value: val,
	}
}

// parse a float literal
func (par *Parser) parseFloatLiteral() ExpressionNode {
	token := par.CurrToken
	val, err := strconv.ParseFloat(token.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("[ERROR] could not parse float literal: %s", token.Literal)
		fmt.Println(msg)
		panic(msg)
	}
	return &FloatLiteralExpressionNode{
		Token: token,
		Value: val,
	}
}

// parse a boolean literal
func (par *Parser) parseBooleanLiteral() ExpressionNode {
	token := par.CurrToken
	return &BooleanLiteralExpressionNode{
		Token: token,
		Value: token.Type == lexer.TRUE_KEY,
	}
}

// parse a null literal
func (par *Parser) parseNilLiteral() ExpressionNode {
	return &NilLiteralExpressionNode{
		Token: par.CurrToken,
		Value: nil,
	}
}

// parse a binary expression
func (par *Parser) parseBinaryExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)

	lVal := eval(par, left)
	rVal := eval(par, right)
	var val int64 = 0
	switch op.Type {
	// arithmetic operators
	case lexer.PLUS_OP:
		val = lVal + rVal
	case lexer.MINUS_OP:
		val = lVal - rVal
	case lexer.MUL_OP:
		val = lVal * rVal
	case lexer.DIV_OP:
		val = lVal / rVal
	case lexer.MOD_OP:
		val = lVal % rVal
	// bitwise operators
	case lexer.BIT_AND_OP:
		val = lVal & rVal
	case lexer.BIT_OR_OP:
		val = lVal | rVal
	case lexer.BIT_XOR_OP:
		val = lVal ^ rVal
	case lexer.BIT_NOT_OP:
		val = ^rVal
	case lexer.BIT_LEFT_OP:
		val = lVal << rVal
	case lexer.BIT_RIGHT_OP:
		val = lVal >> rVal
	}

	binaryExpr := &BinaryExpressionNode{
		Left:      left,
		Operation: op,
		Right:     right,
		Value:     val,
	}
	// par.advance()
	return binaryExpr

}

// parse a unary expression
func (par *Parser) parseUnaryExpression() ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)

	var val int64 = 0
	rVal := eval(par, right)
	switch op.Type {
	case lexer.NOT_OP:
		if rVal == 0 {
			val = 1
		} else {
			val = 0
		}
	case lexer.MINUS_OP:
		val = -rVal
	}

	return &UnaryExpressionNode{
		Operation: op,
		Right:     right,
		Value:     val,
	}
}

// parse a declarative statement
func (par *Parser) parseDeclarativeStatement() StatementNode {
	varToken := par.CurrToken
	par.expectAdvance(lexer.IDENTIFIER_ID)
	identifier := par.CurrToken
	par.expectAdvance(lexer.ASSIGN_OP)
	par.advance()
	expr := par.parseExpression()

	// evaluating the expression
	val := eval(par, expr)

	// save the value in the environment
	par.Env[identifier.Literal] = val

	return &DeclarativeStatementNode{
		VarToken:   varToken,
		Identifier: identifier,
		Expr:       expr,
		Value:      val,
	}
}

// parse an identifier expression
func (par *Parser) parseIdentifierExpression() ExpressionNode {

	// may be an identifier expression or a function call expression
	if par.NextToken.Type == lexer.LEFT_PAREN {
		return par.parseCallExpression()
	}

	varToken := par.CurrToken

	// get the value from the environment
	val := par.Env[varToken.Literal]

	return &IdentifierExpressionNode{
		Name:  varToken.Literal,
		Value: val,
	}
}

// parse a call expression
func (par *Parser) parseCallExpression() ExpressionNode {
	callNode := &CallExpressionNode{}
	callNode.FunctionIdentifier = IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: 0,
	}

	par.expectAdvance(lexer.LEFT_PAREN)
	// if there are arguments, parse them
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		par.advance()
		for {
			callNode.Arguments = append(callNode.Arguments, par.parseExpression())
			if par.NextToken.Type == lexer.COMMA_DELIM {
				par.advance()
				par.advance()
			} else {
				break
			}
		}
	}

	par.expectAdvance(lexer.RIGHT_PAREN)
	return callNode
}

// parse a return statement
func (par *Parser) parseReturnStatement() ExpressionNode {
	returnToken := par.CurrToken
	par.advance()
	expr := par.parseExpression()
	// evaluating the expression
	val := eval(par, expr)
	return &ReturnStatementNode{
		ReturnToken: returnToken,
		Expr:        expr,
		Value:       val,
	}
}

// parse a boolean expression
func (par *Parser) parseBooleanExpression(left ExpressionNode) ExpressionNode {
	// we are already at the left expression
	// so just parse the right expression
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)

	lVal := eval(par, left)
	rVal := eval(par, right)
	val := false
	switch op.Type {
	case lexer.AND_OP:
		if lVal != 0 && rVal != 0 {
			val = true
		}
	case lexer.OR_OP:
		if lVal != 0 || rVal != 0 {
			val = true
		}
	case lexer.GT_OP:
		if lVal > rVal {
			val = true
		}
	case lexer.LT_OP:
		if lVal < rVal {
			val = true
		}
	case lexer.GE_OP:
		if lVal >= rVal {
			val = true
		}
	case lexer.LE_OP:
		if lVal <= rVal {
			val = true
		}
	case lexer.EQ_OP:
		if lVal == rVal {
			val = true
		}
	case lexer.NE_OP:
		if lVal != rVal {
			val = true
		}
	}
	return &BooleanExpressionNode{
		Operation: op,
		Left:      left,
		Right:     right,
		Value:     val,
	}
}

// parse a block statement
func (par *Parser) parseBlockStatement() *BlockStatementNode {
	block := &BlockStatementNode{}
	block.Statements = make([]Node, 0)
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
		}
	}

	return block
}

// parse an assignment expression
func (par *Parser) parseAssignmentExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)

	val := eval(par, right)
	par.Env[left.(*IdentifierExpressionNode).Name] = val

	return &AssignmentExpressionNode{
		Operation: op,
		Left:      left.(*IdentifierExpressionNode).Name,
		Right:     right,
		Value:     val,
	}
}

// parse an if statement
func (par *Parser) parseIfStatement() StatementNode {
	ifNode := NewIfStatement()
	ifNode.IfToken = par.CurrToken
	par.expectAdvance(lexer.LEFT_PAREN)
	ifNode.Condition = par.parseInternal(MINIMUM_PRIORITY)
	par.expectAdvance(lexer.LEFT_BRACE)
	ifNode.ThenBlock = *par.parseBlockStatement()
	if par.NextToken.Type == lexer.ELSE_KEY {
		par.advance() // consume closing brace of if block
		par.advance() // consume else
		if par.CurrToken.Type == lexer.IF_KEY {
			// else if case
			// treat it as a nested if statement
			// wrap it in a block statement
			elseBlock := &BlockStatementNode{}
			elseBlock.Statements = make([]Node, 0)
			nestedIf := par.parseIfStatement()
			elseBlock.Statements = append(elseBlock.Statements, nestedIf)
			if exprNode, ok := nestedIf.(ExpressionNode); ok {
				elseBlock.Value = eval(par, exprNode)
			}
			ifNode.ElseBlock = *elseBlock
		} else {
			ifNode.ElseBlock = *par.parseBlockStatement()
		}
	}
	return ifNode
}

// parse a function statement
func (par *Parser) parseFunctionStatement() StatementNode {
	funcNode := NewFunctionStatementNode()
	funcNode.FuncToken = par.CurrToken
	par.expectAdvance(lexer.IDENTIFIER_ID)
	funcNode.FuncName = IdentifierExpressionNode{
		Name:  par.CurrToken.Literal,
		Value: 0,
	}
	par.expectAdvance(lexer.LEFT_PAREN)

	// Handle empty parameters case
	if par.NextToken.Type != lexer.RIGHT_PAREN {
		// First parameter
		par.expectAdvance(lexer.IDENTIFIER_ID)
		funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
			Name:  par.CurrToken.Literal,
			Value: 0,
		})

		// Subsequent parameters
		for par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Consume comma
			par.expectAdvance(lexer.IDENTIFIER_ID)
			funcNode.FuncParams = append(funcNode.FuncParams, &IdentifierExpressionNode{
				Name:  par.CurrToken.Literal,
				Value: 0,
			})
		}
	}
	par.expectAdvance(lexer.RIGHT_PAREN)

	par.expectAdvance(lexer.LEFT_BRACE)
	funcNode.FuncBody = *par.parseBlockStatement()
	funcNode.Value = funcNode.FuncBody.Value
	return funcNode
}

// parse a string literal
func (par *Parser) parseStringLiteral() ExpressionNode {
	return &StringLiteralExpressionNode{
		Token: par.CurrToken,
		Value: par.CurrToken.Literal,
	}
}

// evaluate the expression
func eval(par *Parser, node ExpressionNode) int64 {
	switch n := node.(type) {
	case *IntegerLiteralExpressionNode:
		return n.Value
	case *BinaryExpressionNode:
		return n.Value
	case *UnaryExpressionNode:
		return n.Value
	case *BooleanLiteralExpressionNode:
		if n.Value {
			return 1
		}
		return 0
	case *ParenthesizedExpressionNode:
		return eval(par, n.Expr)
	case *IdentifierExpressionNode:
		return par.Env[n.Name]
	case *ReturnStatementNode:
		return n.Value
	case *BooleanExpressionNode:
		if n.Value {
			return 1
		}
		return 0
	case *BlockStatementNode:
		return n.Value
	case *AssignmentExpressionNode:
		return n.Value
	case *IfExpressionNode:
		cond := eval(par, n.Condition)
		n.ConditionValue = cond
		if cond != 0 {
			return eval(par, &n.ThenBlock)
		}
		return eval(par, &n.ElseBlock)
	case *StringLiteralExpressionNode:
		return 0
	case *FunctionStatementNode:
		return n.Value
	case *CallExpressionNode:
		return n.Value
	}
	return 0
}

// parse the expression
// Pratt parsing algorithm
func (par *Parser) parseInternal(currPrecedence int) ExpressionNode {
	unary, has := par.UnaryFuncs[par.CurrToken.Type]
	if !has {
		msg := fmt.Sprintf("[ERROR] could not parse unary expression: %s", par.CurrToken.Literal)
		fmt.Println(msg)
		panic(msg)
	}
	left := unary()
	for par.NextToken.Type != lexer.EOF_TYPE && getPrecedence(&par.NextToken) >= currPrecedence {
		binary, has := par.BinaryFuncs[par.NextToken.Type]
		par.advance()
		if !has {
			msg := fmt.Sprintf("[ERROR] could not parse binary expression: %s", par.NextToken.Literal)
			fmt.Println(msg)
			panic(msg)
		}
		left = binary(left)
	}
	return left
}
