package parser

import (
	"fmt"
	"strconv"

	"github.com/akashmaji946/go-mix/lexer"
)

// operator precedence
const (
	MINIMUM_PRIORITY = 0
	OTHER_PRIORITY   = 1
	PLUS_PRIORITY    = 2
	MUL_PRIORITY     = 3
	PREFIX_PRIORITY  = 4

	PAREN_PRIORITY = 5
)

// get the precedence of the operator
func getPrecedence(token *lexer.Token) int {
	switch token.Type {
	case lexer.LEFT_PAREN:
		return PAREN_PRIORITY
	case lexer.NOT_OP:
		return PREFIX_PRIORITY
	case lexer.PLUS_OP, lexer.MINUS_OP:
		return PLUS_PRIORITY
	case lexer.MUL_OP, lexer.DIV_OP:
		return MUL_PRIORITY
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

	// register the functions
	par.registerUnaryFuncs(par.parseParenthesizedExpression, lexer.LEFT_PAREN)
	par.registerUnaryFuncs(par.parseNumberLiteral, lexer.NUMBER_ID)
	par.registerUnaryFuncs(par.parseUnaryExpression, lexer.NOT_OP, lexer.MINUS_OP)
	par.registerUnaryFuncs(par.parseBooleanLiteral, lexer.TRUE_KEY, lexer.FALSE_KEY)

	par.registerBinaryFuncs(par.parseBinaryExpression, lexer.PLUS_OP, lexer.MINUS_OP, lexer.MUL_OP, lexer.DIV_OP)

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

	// TestParser_Parse_OneNumberExpression
	// root := &RootNode{}
	// root.Statements = make([]StatementNode, 1)
	// par.advance()
	// par.advance()
	// token := par.CurrToken
	// root.Statements[0] = &NumberLiteralExpressionNode{
	// 	Token: token,
	// 	Value: 12,
	// }
	// return root

	// TestParser_Parse_AddExpression
	// root := &RootNode{}
	// root.Statements = make([]StatementNode, 1)
	// par.advance()
	// par.advance()
	// left := par.CurrToken
	// par.advance()
	// _ = par.CurrToken
	// par.advance()
	// right := par.CurrToken
	// root.Statements[0] = &AddExpressionNode{
	// 	Left:  &NumberLiteralExpressionNode{Token: left, Value: 12},
	// 	Right: &NumberLiteralExpressionNode{Token: right, Value: 13},
	// 	Value: 25,
	// }
	// return root

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
		if lastStmt, ok := root.Statements[len(root.Statements)-1].(ExpressionNode); ok {
			root.Value = eval(lastStmt)
		}
	}

	return root
}

// parse a statement
// currently only supports expression statements (will be extended)
func (par *Parser) parseStatement() StatementNode {
	switch par.CurrToken {

	// TODO: add for statements like
	// var a = 10;
	// for (a < 10)

	default:
		return par.parseExpression()
	}
}

// parse an expression
// currently only supports binary expressions and unary expressions (will be extended)
func (par *Parser) parseExpression() ExpressionNode {

	// switch par.CurrToken.Type {
	// case lexer.NUMBER_ID:
	// 	par.advance()
	// 	leftToken := par.CurrToken

	// 	par.expectAdvance(lexer.PLUS_OP)

	// 	par.expectAdvance(lexer.NUMBER_ID)

	// 	rightToken := par.CurrToken

	// 	leftValue, _ := strconv.Atoi(leftToken.Literal)
	// 	rightValue, _ := strconv.Atoi(rightToken.Literal)

	// 	return &BinaryExpressionNode{
	// 		Left:  &NumberLiteralExpressionNode{Token: leftToken, Value: leftValue},
	// 		Right: &NumberLiteralExpressionNode{Token: rightToken, Value: rightValue},
	// 		Value: leftValue + rightValue,
	// 	}
	// default:
	// 	return nil
	// }

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
	val, err := strconv.ParseInt(token.Literal, 10, 32)
	if err != nil {
		msg := fmt.Sprintf("[ERROR] could not parse number literal: %s", token.Literal)
		fmt.Println(msg)
		panic(msg)
	}
	return &NumberLiteralExpressionNode{
		Token: token,
		Value: int(val),
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

// parse a binary expression
func (par *Parser) parseBinaryExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)

	lVal := eval(left)
	rVal := eval(right)
	val := 0
	switch op.Type {
	case lexer.PLUS_OP:
		val = lVal + rVal
	case lexer.MINUS_OP:
		val = lVal - rVal
	case lexer.MUL_OP:
		val = lVal * rVal
	case lexer.DIV_OP:
		val = lVal / rVal
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

// evaluate the expression
func eval(node ExpressionNode) int {
	switch n := node.(type) {
	case *NumberLiteralExpressionNode:
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
		return eval(n.Expr)
	}
	return 0
}

// parse the expression
// Pratt parsing algorithm
func (par *Parser) parseInternal(currPrecedence int) ExpressionNode {
	unary, has := par.UnaryFuncs[par.CurrToken.Type]
	if !has {
		panic("[ERROR] could not parse binary expression")
	}
	left := unary()
	for par.NextToken.Type != lexer.EOF_TYPE && getPrecedence(&par.NextToken) >= currPrecedence {
		binary, has := par.BinaryFuncs[par.NextToken.Type]
		par.advance()
		if !has {
			panic("[ERROR] could not parse binary expression")
		}
		left = binary(left)
	}
	return left
}

// parse a unary expression
func (par *Parser) parseUnaryExpression() ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)

	val := 0
	rVal := eval(right)
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
