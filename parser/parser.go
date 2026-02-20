/*
File    : go-mix/parser/parser.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

/*
Package parser implements a Pratt parser (also known as top-down operator precedence parser)
for the Go-Mix programming language.

The parser converts a stream of tokens from the lexer into an Abstract Syntax Tree (AST).
It handles:
- Expressions (binary, unary, literals, identifiers)
- Statements (declarations, assignments, control flow)
- Functions (declarations and calls)
- Loops (for and while)
- Arrays (literals, indexing, slicing)
- Operator precedence and associativity

Key Features:
- Pratt parsing algorithm for efficient expression parsing
- Operator precedence handling (following C-based language standards)
- Error collection (doesn't panic on first error)
- Support for var, let, and const declarations
- Type tracking for let variables
- Compound assignment operators (+=, -=, etc.)

The parser maintains an environment to track variable values during parsing,
which enables constant folding and early error detection.
*/
package parser

import (
	"fmt"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// Parser represents the parser state and configuration.
// It maintains all the information needed to parse Go-Mix source code
// into an Abstract Syntax Tree (AST).
type Parser struct {
	Lex       lexer.Lexer // Lexer instance for tokenizing source code
	CurrToken lexer.Token // Current token being processed
	NextToken lexer.Token // Next token (for lookahead)

	// Function maps for Pratt parsing
	// These maps associate token types with their parsing functions
	UnaryFuncs  map[lexer.TokenType]unaryParseFunction  // Prefix/unary operators and literals
	BinaryFuncs map[lexer.TokenType]binaryParseFunction // Binary/infix operators

	// Environment and variable tracking
	Env map[string]std.GoMixObject // Variable environment (name -> value)

	// Track which variables are const (immutable after declaration)
	Consts map[string]bool

	// Track which variables are declared with let (statically typed)
	LetVars map[string]bool

	// Track the type of let variables for type checking
	LetTypes map[string]std.GoMixType

	// Collect parsing errors instead of panicking
	// This allows reporting multiple errors in a single parse
	Errors []string
}

// NewParser creates and initializes a new Parser instance.
// This is the main entry point for creating a parser.
//
// Parameters:
//
//	src - The Go-Mix source code to parse
//
// Returns:
//
//	A pointer to a fully initialized Parser instance
//
// The parser is ready to use immediately after creation.
// Call Parse() to begin parsing the source code.
func NewParser(src string) *Parser {
	// Create a lexer for the source code
	lex := lexer.NewLexer(src)

	// Create the parser with the lexer
	par := &Parser{
		Lex: lex,
	}

	// Initialize all parser state (maps, tokens, etc.)
	par.init()

	return par
}

// init initializes the parser's internal state.
// This function sets up:
// 1. Function maps for Pratt parsing
// 2. Variable environment and tracking maps
// 3. Error collection
// 4. Initial token lookahead
//
// The function registers parsing functions for all supported token types,
// establishing the grammar of the Go-Mix language.
func (par *Parser) init() {
	// Initialize all maps
	par.UnaryFuncs = make(map[lexer.TokenType]unaryParseFunction)
	par.BinaryFuncs = make(map[lexer.TokenType]binaryParseFunction)
	par.Env = make(map[string]std.GoMixObject)
	par.Consts = make(map[string]bool)
	par.LetVars = make(map[string]bool)
	par.LetTypes = make(map[string]std.GoMixType)
	par.Errors = make([]string, 0)

	// Register unary/prefix parsing functions
	// These handle tokens that can start an expression

	// Parenthesized expressions: (expr)
	par.registerUnaryFuncs(par.parseParenthesizedExpression, lexer.LEFT_PAREN)

	// Numeric literals: 42, 3.14
	par.registerUnaryFuncs(par.parseNumberLiteral, lexer.INT_LIT)
	par.registerUnaryFuncs(par.parseFloatLiteral, lexer.FLOAT_LIT)

	// Char literals: 'a'
	par.registerUnaryFuncs(par.parseCharLiteral, lexer.CHAR_LIT)

	// Return statements: return expr
	par.registerUnaryFuncs(par.parseReturnStatement, lexer.RETURN_KEY)

	// Identifiers: variable names, function names
	par.registerUnaryFuncs(par.parseIdentifierExpression, lexer.IDENTIFIER_ID, lexer.THIS_KEY, lexer.SELF_KEY, lexer.ARRAY_KEY)

	// Boolean literals: true, false
	par.registerUnaryFuncs(par.parseBooleanLiteral, lexer.TRUE_KEY, lexer.FALSE_KEY)

	// String literals: "hello"
	par.registerUnaryFuncs(par.parseStringLiteral, lexer.STRING_LIT)

	// Nil literal: nil
	par.registerUnaryFuncs(par.parseNilLiteral, lexer.NIL_LIT)

	// Register binary/infix parsing functions
	// These handle operators that appear between two expressions

	// Arithmetic operators: +, -, *, /, %
	par.registerBinaryFuncs(par.parseBinaryExpression, lexer.PLUS_OP, lexer.MINUS_OP, lexer.MUL_OP, lexer.DIV_OP, lexer.MOD_OP)

	// Bitwise operators: &, |, ^, <<, >>
	par.registerBinaryFuncs(par.parseBinaryExpression, lexer.BIT_AND_OP, lexer.BIT_OR_OP, lexer.BIT_XOR_OP, lexer.BIT_LEFT_OP, lexer.BIT_RIGHT_OP)

	// Unary operators: !, -, +, ~
	par.registerUnaryFuncs(par.parseUnaryExpression, lexer.NOT_OP, lexer.MINUS_OP, lexer.PLUS_OP, lexer.BIT_NOT_OP)

	// Control flow: if statements
	par.registerUnaryFuncs(par.parseIfStatement, lexer.IF_KEY)

	// Function expressions: func(params) { body }
	par.registerUnaryFuncs(par.parseFunctionAssignment, lexer.FUNC_KEY)

	// Boolean/comparison operators: &&, ||, <, >, <=, >=, ==, !=
	par.registerBinaryFuncs(par.parseBooleanExpression, lexer.AND_OP, lexer.OR_OP, lexer.GT_OP, lexer.LT_OP, lexer.GE_OP, lexer.LE_OP, lexer.EQ_OP, lexer.NE_OP, lexer.STRICT_EQ_OP, lexer.STRICT_NE_OP)

	// Assignment operators: =, +=, -=, *=, /=, %=, &=, |=, ^=, <<=, >>=
	par.registerBinaryFuncs(par.parseAssignmentExpression, lexer.ASSIGN_OP, lexer.PLUS_ASSIGN, lexer.MINUS_ASSIGN, lexer.MUL_ASSIGN, lexer.DIV_ASSIGN, lexer.MOD_ASSIGN,
		lexer.BIT_AND_ASSIGN, lexer.BIT_OR_ASSIGN, lexer.BIT_XOR_ASSIGN, lexer.BIT_LEFT_ASSIGN, lexer.BIT_RIGHT_ASSIGN)

	// Array literals: [1, 2, 3]
	par.registerUnaryFuncs(par.parseArrayExpressionNode, lexer.LEFT_BRACKET)

	// Map literals: map{key: value}
	par.registerUnaryFuncs(par.parseMapKeyword, lexer.MAP_KEY)

	// Set literals: set{1, 2, 3}
	par.registerUnaryFuncs(par.parseSetKeyword, lexer.SET_KEY)

	// Array indexing and slicing: arr[0], arr[1:3]
	par.registerBinaryFuncs(par.parseIndexExpression, lexer.LEFT_BRACKET)

	// Range operator: 2...5
	par.registerBinaryFuncs(par.parseRangeExpression, lexer.RANGE_OP)

	// new keyword for struct instantiation: new Name(args)
	par.registerUnaryFuncs(par.parseNewCallExpression, lexer.NEW_KEY)

	// enum keyword for enum declarations: enum Name { MEMBER1, MEMBER2 }
	par.registerUnaryFuncs(par.parseEnumDeclaration, lexer.ENUM_KEY)

	// memebr access operator: obj.field or obj.method()
	par.registerBinaryFuncs(par.parseMemberAccess, lexer.DOT_OP)

	// Prime the token lookahead by advancing twice
	// After this, CurrToken and NextToken are both valid
	par.advance()
	par.advance()
}

// advance moves the parser forward by one token.
// This implements the token lookahead mechanism:
// - CurrToken becomes NextToken
// - NextToken is fetched from the lexer
//
// This two-token lookahead allows the parser to make decisions
// based on the current token and peek at what's coming next.
func (par *Parser) advance() {
	par.CurrToken = par.NextToken
	par.NextToken = par.Lex.NextToken()
}

// expectAdvance checks if the next token matches the expected type,
// and if so, advances the parser.
//
// Parameters:
//
//	expected - The token type we expect to see next
//
// Returns:
//
//	true if the next token matches and we advanced, false otherwise
//
// This is a common pattern in parsing: "I expect a semicolon next,
// and if it's there, move past it."
func (par *Parser) expectAdvance(expected lexer.TokenType) bool {
	if !par.expectNext(expected) {
		return false
	}
	par.advance()
	return true
}

// expectNext checks if the next token matches the expected type.
// If not, it adds an error message to the error list.
//
// Parameters:
//
//	expected - The token type we expect to see next
//
// Returns:
//
//	true if the next token matches, false otherwise
//
// This function doesn't advance the parser, it only checks.
// Use expectAdvance() if you want to check and advance in one step.
func (par *Parser) expectNext(expected lexer.TokenType) bool {
	if par.NextToken.Type != expected {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected %s, got %s",
			par.NextToken.Line, par.NextToken.Column, expected, par.NextToken.Type)
		par.addError(msg)
		return false
	}
	return true
}

// addError adds an error message to the parser's error list.
// The parser collects errors instead of panicking, allowing it to
// report multiple errors in a single parse.
//
// Parameters:
//
//	msg - The error message to add
func (par *Parser) addError(msg string) {
	par.Errors = append(par.Errors, msg)
}

// HasErrors returns true if there are parsing errors.
// This should be checked after parsing to determine if the parse was successful.
//
// Returns:
//
//	true if there are any errors, false if parsing was successful
func (par *Parser) HasErrors() bool {
	return len(par.Errors) > 0
}

// GetErrors returns all parsing errors collected during parsing.
// This allows the caller to display all errors to the user.
//
// Returns:
//
//	A slice of error message strings
func (par *Parser) GetErrors() []string {
	return par.Errors
}

// Parse is the main parsing function that converts source code into an AST.
// It repeatedly parses statements until reaching the end of the file (EOF),
// building up a RootNode that contains all the parsed statements.
//
// Returns:
//
//	A pointer to a RootNode containing all parsed statements and the final value
//
// The function also computes the value of the root node by evaluating the last
// statement, which allows the REPL to display the result of the last expression.
//
// Example:
//
//	For input "var x = 5; x + 10", the root value would be 15 (the result of x + 10)
func (par *Parser) Parse() *RootNode {

	// Create the root node that will hold all statements
	root := &RootNode{}
	root.Statements = make([]StatementNode, 0)

	// Parse statements until we reach the end of file
	for par.CurrToken.Type != lexer.EOF_TYPE {
		stmt := par.parseStatement()
		if stmt != nil {
			root.Statements = append(root.Statements, stmt)
		}
		par.advance()
	}

	// Compute the value of the root node by evaluating the last statement
	// This allows the REPL to display the result of the last expression
	if len(root.Statements) > 0 {
		lastStmt := root.Statements[len(root.Statements)-1]
		// Try to extract a value from different statement types
		if exprNode, ok := lastStmt.(ExpressionNode); ok {
			root.Value = parseEval(par, exprNode)
		} else if declNode, ok := lastStmt.(*DeclarativeStatementNode); ok {
			root.Value = declNode.Value
		} else if returnNode, ok := lastStmt.(*ReturnStatementNode); ok {
			root.Value = returnNode.Value
		} else if blockNode, ok := lastStmt.(*BlockStatementNode); ok {
			root.Value = blockNode.Value
		} else if funcNode, ok := lastStmt.(*FunctionStatementNode); ok {
			root.Value = funcNode.Value
		} else if forLoopNode, ok := lastStmt.(*ForLoopStatementNode); ok {
			root.Value = forLoopNode.Value
		} else if whileLoopNode, ok := lastStmt.(*WhileLoopStatementNode); ok {
			root.Value = whileLoopNode.Value
		} else {
			root.Value = &std.Nil{}
		}
	} else {
		// Empty program evaluates to nil
		root.Value = &std.Nil{}
	}

	return root
}

// parseInternal is the core of the Pratt parsing algorithm.
// It parses expressions while respecting operator precedence.
//
// Parameters:
//
//	currPrecedence - The minimum precedence level for operators to parse
//
// Returns:
//
//	An ExpressionNode representing the parsed expression
//
// Algorithm:
//  1. Parse a prefix expression (unary operator or primary expression)
//  2. While the next operator has higher precedence than currPrecedence:
//     a. Parse the operator as an infix expression
//     b. The result becomes the new left operand
//  3. Return the final expression
//
// This elegant algorithm handles operator precedence and associativity
// without needing separate precedence levels or recursive descent for each level.
func (par *Parser) parseInternal(currPrecedence int) ExpressionNode {
	unary, has := par.UnaryFuncs[par.CurrToken.Type]
	if !has {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected token: %s",
			par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
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
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected operator: %s",
				par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
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

// parseStatement parses a single statement.
// This is the main dispatcher that determines what type of statement to parse
// based on the current token.
//
// Returns:
//
//	A StatementNode representing the parsed statement, or nil for empty statements
//
// Supported statement types:
//   - Variable declarations (var, let, const)
//   - Block statements ({ ... })
//   - If statements
//   - Function declarations
//   - For loops
//   - While loops
//   - Expression statements (any expression followed by semicolon)
func (par *Parser) parseStatement() StatementNode {
	switch par.CurrToken.Type {

	// ignore semicolons
	case lexer.SEMICOLON_DELIM:
		return nil

	// var a = (expression); // dynamically typed variable declaration
	case lexer.VAR_KEY:
		return par.parseDeclarativeStatement()
	// let a = (expression);  // statically typed variable declaration
	case lexer.LET_KEY:
		return par.parseDeclarativeStatement()
	// const a = (expression); // immutable variable declaration
	case lexer.CONST_KEY:
		return par.parseDeclarativeStatement()

	// {.....}
	case lexer.LEFT_BRACE:
		return par.parseBlockStatement()
	// if (condition) { ... } [ [else if () {...} ] [else if () {...} ]... [else { ... }] ]
	case lexer.IF_KEY:
		return par.parseIfStatement()

	// func functionName(params) { ... }
	case lexer.FUNC_KEY:
		return par.parseFunctionStatement()

	// for (init; condition; update) { ... }
	case lexer.FOR_KEY:
		return par.parseForLoop()

	// while (condition) { ... }
	case lexer.WHILE_KEY:
		return par.parseWhileLoop()

	// foreach item in iterable { ... }
	case lexer.FOREACH_KEY:
		return par.parseForeachLoop()

	// struct StructName { field1; field2; ... func foo(params){...} ... }
	case lexer.STRUCT_KEY:
		return par.parseStructDeclaration()

	// break;
	case lexer.BREAK_KEY:
		return par.parseBreakStatement()

	// continue;
	case lexer.CONTINUE_KEY:
		return par.parseContinueStatement()

	// import "module" as alias;
	case lexer.IMPORT_KEY:
		return par.parseImportStatement()

	// switch (expression) { case value: ... default: ... }
	case lexer.SWITCH_KEY:
		return par.parseSwitchStatement()

	default:
		return par.parseExpression()
	}
}

// parseEval evaluates an expression node during parsing.
// This enables constant folding and early error detection.
//
// Parameters:
//
//	par  - The parser instance (for accessing environment)
//	node - The expression node to evaluate
//
// Returns:
//
//	The evaluated GoMixObject value
//
// Note: This is a simplified evaluator used during parsing.
// The full evaluator in the parseEval package handles more complex cases.
func parseEval(par *Parser, node ExpressionNode) std.GoMixObject {
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
		return parseEval(par, n.Expr)
	case *IdentifierExpressionNode:
		if val, ok := par.Env[n.Name]; ok {
			return val
		}
		return &std.Nil{}
	case *ReturnStatementNode:
		return parseEval(par, n.Expr)
	case *BooleanExpressionNode:
		return n.Value
	case *BlockStatementNode:
		var result std.GoMixObject = &std.Nil{}
		for _, stmt := range n.Statements {
			if expr, ok := stmt.(ExpressionNode); ok {
				val := parseEval(par, expr)
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
		cond := parseEval(par, n.Condition)
		n.ConditionValue = cond
		// Check truthiness
		var isTrue bool
		if cond.GetType() == std.BooleanType {
			isTrue = cond.(*std.Boolean).Value
		} else if cond.GetType() == std.IntegerType {
			isTrue = cond.(*std.Integer).Value != 0
		} else {
			isTrue = false
		}

		if isTrue {
			return parseEval(par, &n.ThenBlock)
		}
		return parseEval(par, &n.ElseBlock)
	case *StringLiteralExpressionNode:
		return n.Value
	case *FunctionStatementNode:
		return n.Value
	case *CallExpressionNode:
		return n.Value
	case *FloatLiteralExpressionNode:
		return n.Value
	}
	return &std.Nil{}
}
