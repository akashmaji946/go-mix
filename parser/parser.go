package parser

import (
	"fmt"
	"strconv"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
)

// operator precedence (following C-based language standards)
// Higher number = higher precedence (binds tighter)
const (
	MINIMUM_PRIORITY = 0

	// Assignment operators (lowest precedence, right-to-left associativity)
	// = += -= *= /= %= &= |= ^= <<= >>=
	ASSIGN_PRIORITY = 10

	// Logical OR: ||
	OR_PRIORITY = 40

	// Logical AND: &&
	AND_PRIORITY = 50

	// Bitwise OR: |
	BIT_OR_PRIORITY = 60

	// Bitwise XOR: ^
	BIT_XOR_PRIORITY = 70

	// Bitwise AND: &
	BIT_AND_PRIORITY = 80

	// Equality operators: == !=
	EQUALITY_PRIORITY = 90

	// Relational operators: < > <= >=
	RELATIONAL_PRIORITY = 100

	// Shift operators: << >>
	SHIFT_PRIORITY = 110

	// Additive operators: + -
	PLUS_PRIORITY = 120

	// Multiplicative operators: * / %
	MUL_PRIORITY = 130

	// Unary/Prefix operators: ! ~ + -
	PREFIX_PRIORITY = 140

	// Parentheses (highest precedence)
	PAREN_PRIORITY = 150
)

// get the precedence of the operator
func getPrecedence(token *lexer.Token) int {
	switch token.Type {

	// Parentheses - highest precedence
	case lexer.LEFT_PAREN:
		return PAREN_PRIORITY

	// Unary/Prefix operators: ! ~
	case lexer.NOT_OP, lexer.BIT_NOT_OP:
		return PREFIX_PRIORITY

	// Multiplicative: * / %
	case lexer.MUL_OP, lexer.DIV_OP, lexer.MOD_OP:
		return MUL_PRIORITY

	// Additive: + -
	case lexer.PLUS_OP, lexer.MINUS_OP:
		return PLUS_PRIORITY

	// Shift: << >>
	case lexer.BIT_LEFT_OP, lexer.BIT_RIGHT_OP:
		return SHIFT_PRIORITY

	// Relational: < > <= >=
	case lexer.GT_OP, lexer.LT_OP, lexer.GE_OP, lexer.LE_OP:
		return RELATIONAL_PRIORITY

	// Equality: == !=
	case lexer.EQ_OP, lexer.NE_OP:
		return EQUALITY_PRIORITY

	// Bitwise AND: &
	case lexer.BIT_AND_OP:
		return BIT_AND_PRIORITY

	// Bitwise XOR: ^
	case lexer.BIT_XOR_OP:
		return BIT_XOR_PRIORITY

	// Bitwise OR: |
	case lexer.BIT_OR_OP:
		return BIT_OR_PRIORITY

	// Logical AND: &&
	case lexer.AND_OP:
		return AND_PRIORITY

	// Logical OR: ||
	case lexer.OR_OP:
		return OR_PRIORITY

	// Assignment operators (lowest precedence)
	case lexer.ASSIGN_OP, lexer.PLUS_ASSIGN, lexer.MINUS_ASSIGN, lexer.MUL_ASSIGN, lexer.DIV_ASSIGN, lexer.MOD_ASSIGN,
		lexer.BIT_AND_ASSIGN, lexer.BIT_OR_ASSIGN, lexer.BIT_XOR_ASSIGN, lexer.BIT_LEFT_ASSIGN, lexer.BIT_RIGHT_ASSIGN:
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

	Env map[string]objects.GoMixObject
	// Track which variables are const
	Consts map[string]bool
	// Track which variables are declared with let (statically typed)
	LetVars map[string]bool
	// Track the type of let variables
	LetTypes map[string]objects.GoMixType
	// Collect parsing errors instead of panicking
	Errors []string
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
	par.Env = make(map[string]objects.GoMixObject)
	par.Consts = make(map[string]bool)
	par.LetVars = make(map[string]bool)
	par.LetTypes = make(map[string]objects.GoMixType)
	par.Errors = make([]string, 0)

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

	par.registerUnaryFuncs(par.parseUnaryExpression, lexer.NOT_OP, lexer.MINUS_OP, lexer.PLUS_OP, lexer.BIT_NOT_OP)
	par.registerUnaryFuncs(par.parseIfStatement, lexer.IF_KEY)
	par.registerUnaryFuncs(par.parseFunctionAssignment, lexer.FUNC_KEY)

	par.registerBinaryFuncs(par.parseBooleanExpression, lexer.AND_OP, lexer.OR_OP, lexer.GT_OP, lexer.LT_OP, lexer.GE_OP, lexer.LE_OP, lexer.EQ_OP, lexer.NE_OP)
	par.registerUnaryFuncs(par.parseUnaryExpression, lexer.NOT_OP, lexer.MINUS_OP)
	par.registerBinaryFuncs(par.parseAssignmentExpression, lexer.ASSIGN_OP, lexer.PLUS_ASSIGN, lexer.MINUS_ASSIGN, lexer.MUL_ASSIGN, lexer.DIV_ASSIGN, lexer.MOD_ASSIGN,
		lexer.BIT_AND_ASSIGN, lexer.BIT_OR_ASSIGN, lexer.BIT_XOR_ASSIGN, lexer.BIT_LEFT_ASSIGN, lexer.BIT_RIGHT_ASSIGN)

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
func (par *Parser) expectAdvance(expected lexer.TokenType) bool {
	if !par.expectNext(expected) {
		return false
	}
	par.advance()
	return true
}

// expect the next token to be of the expected type
func (par *Parser) expectNext(expected lexer.TokenType) bool {
	if par.NextToken.Type != expected {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected %s, got %s", par.NextToken.Line, par.NextToken.Column, expected, par.NextToken.Type)
		par.addError(msg)
		return false
	}
	return true
}

// addError adds an error message to the parser's error list
func (par *Parser) addError(msg string) {
	par.Errors = append(par.Errors, msg)
}

// HasErrors returns true if there are parsing errors
func (par *Parser) HasErrors() bool {
	return len(par.Errors) > 0
}

// GetErrors returns all parsing errors
func (par *Parser) GetErrors() []string {
	return par.Errors
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
		} else if forLoopNode, ok := lastStmt.(*ForLoopNode); ok {
			root.Value = forLoopNode.Value
		} else if whileLoopNode, ok := lastStmt.(*WhileLoopNode); ok {
			root.Value = whileLoopNode.Value
		} else {
			root.Value = &objects.Nil{}
		}
	} else {
		root.Value = &objects.Nil{}
	}

	return root
}

// parseFunctionAssignment
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
	if paren.Expr == nil {
		return nil
	}
	paren.Value = eval(par, paren.Expr)
	if !par.expectAdvance(lexer.RIGHT_PAREN) {
		return nil
	}

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
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: could not parse number literal: %s", token.Line, token.Column, token.Literal)
			par.addError(msg)
			return nil
		}
	}
	return &IntegerLiteralExpressionNode{
		Token: token,
		Value: &objects.Integer{Value: val},
	}
}

// parse a float literal
func (par *Parser) parseFloatLiteral() ExpressionNode {
	token := par.CurrToken
	val, err := strconv.ParseFloat(token.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: could not parse float literal: %s", token.Line, token.Column, token.Literal)
		par.addError(msg)
		return nil
	}
	return &FloatLiteralExpressionNode{
		Token: token,
		Value: &objects.Float{Value: val},
	}
}

// parse a boolean literal
func (par *Parser) parseBooleanLiteral() ExpressionNode {
	token := par.CurrToken
	return &BooleanLiteralExpressionNode{
		Token: token,
		Value: &objects.Boolean{Value: token.Type == lexer.TRUE_KEY},
	}
}

// parse a null literal
func (par *Parser) parseNilLiteral() ExpressionNode {
	return &NilLiteralExpressionNode{
		Token: par.CurrToken,
		Value: &objects.Nil{},
	}
}

// evaluate the expression
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

// helper to cast to int64 for internal usage if needed, or handle objects
// actually we should perform operations on objects now.

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

// parse an if expression
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

func toFloat64(obj objects.GoMixObject) float64 {
	if obj.GetType() == objects.IntegerType {
		return float64(obj.(*objects.Integer).Value)
	} else if obj.GetType() == objects.FloatType {
		return obj.(*objects.Float).Value
	}
	return 0
}

// evaluateCompoundBinaryOp evaluates a binary operation for compound assignments
// It takes the left and right operand values and the binary operator type
// Returns the computed result as a GoMixObject
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

// getCompoundBinaryOp returns the corresponding binary operator for a compound assignment operator
// Returns the binary operator type and true if it's a compound assignment, otherwise returns empty and false
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

// parseCompoundAssignment handles compound assignment expressions (+=, -=, *=, etc.)
// It transforms them into regular assignments with a binary expression on the right side
// e.g., a += 5 becomes a = a + 5
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
		Left:      *ident,
		Right:     binaryExpr,
		Value:     binaryVal,
	}
}

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

// parse a declarative statement
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

// parse an identifier expression
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

// parse a call expression
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

// parse a return statement
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

// parse a boolean expression
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

// parse a block statement
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
		} else if forLoopNode, ok := lastStmt.(*ForLoopNode); ok {
			block.Value = forLoopNode.Value
		} else if whileLoopNode, ok := lastStmt.(*WhileLoopNode); ok {
			block.Value = whileLoopNode.Value
		} else {
			block.Value = &objects.Nil{}
		}
	} else {
		block.Value = &objects.Nil{} // Default value for an empty block
	}

	return block
}

// parse an assignment expression
func (par *Parser) parseAssignmentExpression(left ExpressionNode) ExpressionNode {
	op := par.CurrToken
	par.advance()
	right := par.parseInternal(getPrecedence(&op) + 1)
	if right == nil {
		return nil
	}

	// if left is const - evaluator will handle this error
	// just continue parsing without adding to parser errors
	ident, ok := left.(*IdentifierExpressionNode)
	if !ok {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: invalid assignment target", par.CurrToken.Line, par.CurrToken.Column)
		par.addError(msg)
		return nil
	}

	// Check if this is a compound assignment (+=, -=, *=, etc.)
	if binaryOp, isCompound := getCompoundBinaryOp(op.Type); isCompound {
		return par.parseCompoundAssignment(left, ident, op, right, binaryOp)
	}

	// Regular assignment - evaluate and store
	val := eval(par, right)
	par.Env[ident.Name] = val
	return &AssignmentExpressionNode{
		Operation: op,
		Left:      *ident,
		Right:     right,
		Value:     val,
	}
}

// parse an if statement
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

// parse a function statement
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

// parse a for loop statement
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

	return &ForLoopNode{
		ForToken:     forToken,
		Initializers: initializers,
		Condition:    condition,
		Updates:      updates,
		Body:         *body,
		Value:        &objects.Nil{},
	}
}

// parse a while loop statement
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

	return &WhileLoopNode{
		WhileToken: whileToken,
		Conditions: conditions,
		Body:       *body,
		Value:      &objects.Nil{},
	}
}

// parse a string literal
func (par *Parser) parseStringLiteral() ExpressionNode {
	return &StringLiteralExpressionNode{
		Token: par.CurrToken,
		Value: &objects.String{Value: par.CurrToken.Literal},
	}
}

// parse the expression
// Pratt parsing algorithm
func (par *Parser) parseInternal(currPrecedence int) ExpressionNode {
	unary, has := par.UnaryFuncs[par.CurrToken.Type]
	if !has {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected token: %s", par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
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
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected operator: %s", par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
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
