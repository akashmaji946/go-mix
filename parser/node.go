/*
File    : go-mix/parser/node.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/objects"
)

// NodeVisitor: implements the Visitor design pattern for traversing the Abstract Syntax Tree (AST)
// Each Visit method processes a specific node type, enabling operations like evaluation, printing, or transformation
type NodeVisitor interface {
	// Base node visitors for generic expression and statement handling
	VisitExpressionNode(node ExpressionNode)
	VisitStatementNode(node StatementNode)
	VisitRootNode(node RootNode) // Entry point for visiting the entire program

	// Literal value visitors - handle primitive data types
	VisitIntegerLiteralExpressionNode(node IntegerLiteralExpressionNode) // Integer literals: 42, -15, 0
	VisitBooleanLiteralExpressionNode(node BooleanLiteralExpressionNode) // Boolean literals: true, false
	VisitFloatLiteralExpressionNode(node FloatLiteralExpressionNode)     // Float literals: 3.14, -2.5
	VisitStringLiteralExpressionNode(node StringLiteralExpressionNode)   // String literals: "hello", 'world'
	VisitNilLiteralExpressionNode(node NilLiteralExpressionNode)         // Nil/null literal

	// Expression visitors - handle operations and computations
	VisitBinaryExpressionNode(node BinaryExpressionNode)               // Binary operations: +, -, *, /, %
	VisitUnaryExpressionNode(node UnaryExpressionNode)                 // Unary operations: -, !, +
	VisitBooleanExpressionNode(node BooleanExpressionNode)             // Boolean operations: &&, ||, ==, !=, <, >
	VisitParenthesizedExpressionNode(node ParenthesizedExpressionNode) // Parenthesized expressions: (expr)

	// Statement and identifier visitors
	VisitDeclarativeStatementNode(node DeclarativeStatementNode) // Variable declarations: var x = 10, let y = 5
	VisitIdentifierExpressionNode(node IdentifierExpressionNode) // Variable/function identifiers: x, myVar
	VisitReturnStatementNode(node ReturnStatementNode)           // Return statements: return expr
	VisitBlockStatementNode(node BlockStatementNode)             // Code blocks: { stmt1; stmt2; }
	VisitAssignmentExpressionNode(node AssignmentExpressionNode) // Assignments: x = 10

	// Conditional control flow visitors
	VisitIfExpressionNode(node IfExpressionNode) // If-else conditionals: if (cond) { ... } else { ... }

	// Function-related visitors
	VisitFunctionStatementNode(node FunctionStatementNode) // Function definitions: func name(params) { body }
	VisitCallExpressionNode(node CallExpressionNode)       // Function calls: funcName(arg1, arg2)

	// Loop control flow visitors
	VisitForLoopStatementNode(node ForLoopStatementNode)     // For loops: for(init; cond; update) { ... }
	VisitWhileLoopStatementNode(node WhileLoopStatementNode) // While loops: while(cond) { ... }

	// Data structure visitors - handle collections
	VisitArrayExpressionNode(node ArrayExpressionNode) // Array literals: [1, 2, 3]
	VisitMapExpressionNode(node MapExpressionNode)     // Map literals: map{key: value}
	VisitSetExpressionNode(node SetExpressionNode)     // Set literals: set{1, 2, 3}
	VisitIndexExpressionNode(node IndexExpressionNode) // Array indexing: arr[0], arr[-1]
	VisitSliceExpressionNode(node SliceExpressionNode) // Array slicing: arr[1:3], arr[:5], arr[2:]

	// Range and foreach visitors
	VisitRangeExpressionNode(node RangeExpressionNode)           // Range expressions: 2...5
	VisitForeachLoopStatementNode(node ForeachLoopStatementNode) // Foreach loops: foreach i in range { ... }
}

// Node: base interface for all nodes of the AST
// Literal(): returns the string representation of the node
// Accept(): accepts a visitor
type Node interface {
	Literal() string
	Accept(visitor NodeVisitor)
}

// StatementNode: base interface for all statement nodes
// Node: every statement node is a node
// Statement(): returns the string representation of the node
type StatementNode interface {
	Node
	Statement()
}

// ExpressionNode: base interface for all expression nodes
// Node: every expression node is a node
// StatementNode: every expression is also a statement
// Expression(): returns the string representation of the node
type ExpressionNode interface {
	Node
	StatementNode
	Expression()
}

// RootNode: represents the root of the AST (the program node)
// Statements: list of statements in the program
// Value: value of the program (for expressions, or value of final expression below root)
type RootNode struct {
	Statements []StatementNode     // every line of code is a statement
	Value      objects.GoMixObject // (e.g. 2 + 3 * 4 + 2 => 16)
}

// RootNode.Literal(): string represenation of the root node's statement/expression
func (root *RootNode) Literal() string {
	res := ""
	for _, stmt := range root.Statements {
		res += stmt.Literal()
		res += ";"
	}
	return res
}

// RootNode.Accept(): accepts a visitor (eg PrintVisitor)
func (root *RootNode) Accept(visitor NodeVisitor) {
	visitor.VisitRootNode(*root)
}

// There can be many types of ExpressionNodes
// IntegerLiteralExpressionNode: represents an integer number literal
// Example: 42, 0, -15
type IntegerLiteralExpressionNode struct {
	Token lexer.Token         // The integer token with its literal value
	Value objects.GoMixObject // The integer object value
}

// NumberLiteralExpressionNode.Literal(): string represenation of the node
func (node *IntegerLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// NumberLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *IntegerLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIntegerLiteralExpressionNode(*node)
}

// NumberLiteralExpressionNode.Statement(): every expression is also a statement
func (node *IntegerLiteralExpressionNode) Statement() {

}

// NumberLiteralExpressionNode.Expression(): every expression node is a node
func (node *IntegerLiteralExpressionNode) Expression() {

}

// FloatLiteralExpressionNode: represents a floating-point number literal
// Example: 3.14, 0.5, -2.718
type FloatLiteralExpressionNode struct {
	Token lexer.Token         // The float token with its literal value
	Value objects.GoMixObject // The float object value
}

// FloatLiteralExpressionNode.Literal(): string represenation of the node
func (node *FloatLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// FloatLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *FloatLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitFloatLiteralExpressionNode(*node)
}

// FloatLiteralExpressionNode.Statement(): every expression is also a statement
func (node *FloatLiteralExpressionNode) Statement() {

}

// FloatLiteralExpressionNode.Expression(): every expression node is a node
func (node *FloatLiteralExpressionNode) Expression() {

}

// BooleanLiteralExpressionNode: represents a boolean literal value
// Example: true or false
type BooleanLiteralExpressionNode struct {
	Token lexer.Token         // The boolean token (true/false)
	Value objects.GoMixObject // The boolean object value
}

// BooleanLiteralExpressionNode.Literal(): string represenation of the node
func (node *BooleanLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// BooleanLiteralExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BooleanLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitBooleanLiteralExpressionNode(*node)
}

// BooleanLiteralExpressionNode.Statement(): every expression is also a statement
func (node *BooleanLiteralExpressionNode) Statement() {

}

// BooleanLiteralExpressionNode.Expression(): every expression node is a node
func (node *BooleanLiteralExpressionNode) Expression() {

}

// BinaryExpressionNode: represents a binary operation expression with two operands
// Example: 2 + 3, x * y, a - b
type BinaryExpressionNode struct {
	Operation lexer.Token         // The binary operator token (+, -, *, /, %, etc.)
	Left      ExpressionNode      // Left operand expression
	Right     ExpressionNode      // Right operand expression
	Value     objects.GoMixObject // Evaluated result of the operation
}

// BinaryExpressionNode.Literal(): string represenation of the node
func (node *BinaryExpressionNode) Literal() string {
	return node.Left.Literal() + node.Operation.Literal + node.Right.Literal()
}

// BinaryExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BinaryExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitBinaryExpressionNode(*node)
}

// BinaryExpressionNode.Statement(): every expression is also a statement
func (node *BinaryExpressionNode) Statement() {

}

// BinaryExpressionNode.Expression(): every expression node is a node
func (node *BinaryExpressionNode) Expression() {

}

// UnaryExpressionNode: represents a unary operation expression with one operand
// Example: -x, !flag, +5
type UnaryExpressionNode struct {
	Operation lexer.Token         // The unary operator token (-, !, +)
	Right     ExpressionNode      // The operand expression
	Value     objects.GoMixObject // Evaluated result of the operation
}

// UnaryExpressionNode.Literal(): string represenation of the node
func (node *UnaryExpressionNode) Literal() string {
	return node.Operation.Literal + node.Right.Literal()
}

// UnaryExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *UnaryExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitUnaryExpressionNode(*node)
}

// UnaryExpressionNode.Statement(): every expression is also a statement
func (node *UnaryExpressionNode) Statement() {

}

// UnaryExpressionNode.Expression(): every expression node is a node
func (node *UnaryExpressionNode) Expression() {

}

// BooleanExpressionNode: represents an expression with a boolean operator (&&, ||, ==, !=, <, >, <=, >=)
// Used for logical and comparison operations between two expressions
type BooleanExpressionNode struct {
	Operation lexer.Token         // The boolean operator token
	Left      ExpressionNode      // Left operand expression
	Right     ExpressionNode      // Right operand expression
	Value     objects.GoMixObject // Evaluated boolean result
}

// ParenthesizedExpressionNode: represents an expression wrapped in parentheses for precedence control
// Example: (2 + 3) * 4
type ParenthesizedExpressionNode struct {
	Expr  ExpressionNode      // The inner expression
	Value objects.GoMixObject // Evaluated value of the expression
}

// ParenthesizedExpressionNode.Literal(): string represenation of the node
func (node *ParenthesizedExpressionNode) Literal() string {
	return "(" + node.Expr.Literal() + ")"
}

// ParenthesizedExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *ParenthesizedExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitParenthesizedExpressionNode(*node)
}

// ParenthesizedExpressionNode.Statement(): every expression is also a statement
func (node *ParenthesizedExpressionNode) Statement() {

}

// ParenthesizedExpressionNode.Expression(): every expression node is a node
func (node *ParenthesizedExpressionNode) Expression() {

}

// DeclarativeStatementNode: represents a variable declaration statement
// Example: var x = 10 or let name = "John"
type DeclarativeStatementNode struct {
	VarToken   lexer.Token              // The declaration keyword token (var/let)
	Identifier IdentifierExpressionNode // The variable identifier being declared
	Expr       ExpressionNode           // The initialization expression
	Value      objects.GoMixObject      // The assigned value
}

// DeclarativeStatementNode.Literal(): string represenation of the node
func (node *DeclarativeStatementNode) Literal() string {
	return node.VarToken.Literal + " " + node.Identifier.Name + " = " + node.Expr.Literal()
}

// DeclarativeStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *DeclarativeStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitDeclarativeStatementNode(*node)
}

// DeclarativeStatementNode.Statement(): every expression is also a statement
func (node *DeclarativeStatementNode) Statement() {

}

// DeclarativeStatementNode.Expression(): we should not have it
// why? because declarative statement is not an expression
// func (node *DeclarativeStatementNode) Expression() {
// }

// IdentifierExpressionNode: represents a variable or function identifier
// Example: x, myVar, functionName
type IdentifierExpressionNode struct {
	Name  string              // The identifier name
	Value objects.GoMixObject // The value associated with this identifier
	Type  string              // The type of the identifier (if applicable)
	IsLet bool                // Whether this was declared with 'let' (immutable)
}

// IdentifierExpressionNode.Literal(): string represenation of the node
func (node *IdentifierExpressionNode) Literal() string {
	return node.Name
}

// IdentifierExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *IdentifierExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIdentifierExpressionNode(*node)
}

// IdentifierExpressionNode.Statement(): every expression is also a statement
func (node *IdentifierExpressionNode) Statement() {

}

// IdentifierExpressionNode.Expression(): every expression node is a node
func (node *IdentifierExpressionNode) Expression() {

}

// ReturnStatementNode: represents a return statement in a function
// Example: return x + 5 or return "result"
type ReturnStatementNode struct {
	ReturnToken lexer.Token         // The 'return' keyword token
	Expr        ExpressionNode      // The expression to return
	Value       objects.GoMixObject // The evaluated return value
}

// ReturnStatementNode.Literal(): string represenation of the node
func (node *ReturnStatementNode) Literal() string {
	return node.ReturnToken.Literal + " " + node.Expr.Literal()
}

// ReturnStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *ReturnStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitReturnStatementNode(*node)
}

// ReturnStatementNode.Statement(): every expression is also a statement
func (node *ReturnStatementNode) Statement() {

}

// ReturnStatementNode.Expression(): every expression node is a node
func (node *ReturnStatementNode) Expression() {

}

// BlockStatementNode: represents a block of statements enclosed in braces
// Example: { stmt1; stmt2; stmt3; }
type BlockStatementNode struct {
	Statements []StatementNode     // List of statements in the block
	Value      objects.GoMixObject // Value of the last expression in the block
}

// BlockStatementNode.Literal(): string represenation of the node
func (node *BlockStatementNode) Literal() string {
	str := "{"
	for _, stmt := range node.Statements {
		str += stmt.Literal()
		str += ";"
	}
	str += "}"
	return str
}

// BlockStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BlockStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitBlockStatementNode(*node)
}

// BlockStatementNode.Statement(): every expression is also a statement
func (node *BlockStatementNode) Statement() {

}

// BlockStatementNode.Expression(): every expression node is a node
func (node *BlockStatementNode) Expression() {

}

// BooleanExpressionNode.Literal(): string represenation of the node
func (node *BooleanExpressionNode) Literal() string {
	return node.Left.Literal() + node.Operation.Literal + node.Right.Literal()
}

// BooleanExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *BooleanExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitBooleanExpressionNode(*node)
}

// BooleanExpressionNode.Statement(): every expression is also a statement
func (node *BooleanExpressionNode) Statement() {

}

// BooleanExpressionNode.Expression(): every expression node is a node
func (node *BooleanExpressionNode) Expression() {

}

// AssignmentExpressionNode: represents a variable assignment expression
// Example: x = 10, count = count + 1, a[0] = 11, map["key"] = value
type AssignmentExpressionNode struct {
	Operation lexer.Token         // The assignment operator token (=)
	Left      ExpressionNode      // The target being assigned to (identifier or index expression)
	Right     ExpressionNode      // The expression being assigned
	Value     objects.GoMixObject // The assigned value
}

// AssignmentExpressionNode.Literal(): string represenation of the node
func (node *AssignmentExpressionNode) Literal() string {
	return node.Left.Literal() + " " + node.Operation.Literal + " " + node.Right.Literal()
}

// AssignmentExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *AssignmentExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitAssignmentExpressionNode(*node)
}

// AssignmentExpressionNode.Statement(): every expression is also a statement
func (node *AssignmentExpressionNode) Statement() {

}

// AssignmentExpressionNode.Expression(): every expression node is a node
func (node *AssignmentExpressionNode) Expression() {

}

// IfExpressionNode: represents an if-else conditional expression
// Example: if (x > 0) { ... } else { ... }
type IfExpressionNode struct {
	IfToken        lexer.Token         // The 'if' keyword token
	Condition      ExpressionNode      // The condition expression to evaluate
	ConditionValue objects.GoMixObject // Evaluated condition result
	ThenBlock      BlockStatementNode  // Block to execute if condition is true
	ElseBlock      BlockStatementNode  // Block to execute if condition is false (optional)
}

// IfExpressionNode.Literal(): string represenation of the node
func (node *IfExpressionNode) Literal() string {
	res := node.IfToken.Literal + " " + node.Condition.Literal() + " " + node.ThenBlock.Literal()
	if len(node.ElseBlock.Statements) > 0 {
		if len(node.ElseBlock.Statements) == 1 {
			if nestedIf, ok := node.ElseBlock.Statements[0].(*IfExpressionNode); ok {
				return res + " else " + nestedIf.Literal()
			}
		}
		return res + " else " + node.ElseBlock.Literal()
	}
	return res
}

// IfExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *IfExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIfExpressionNode(*node)
}

// IfExpressionNode.Statement(): every expression is also a statement
func (node *IfExpressionNode) Statement() {

}

// IfExpressionNode.Expression(): every expression node is a node
func (node *IfExpressionNode) Expression() {

}

// EMPTY_BLOCK: a reusable empty block statement node
// Used as a default value for optional blocks (e.g., else blocks without statements)
var EMPTY_BLOCK = &BlockStatementNode{
	Statements: []StatementNode{},
}

// NewIfStatement: creates a new if statement node with default empty values
// Returns an initialized IfExpressionNode ready for parsing
func NewIfStatement() *IfExpressionNode {
	return &IfExpressionNode{
		Condition:      &BinaryExpressionNode{},
		ThenBlock:      *EMPTY_BLOCK,
		ElseBlock:      *EMPTY_BLOCK,
		ConditionValue: &objects.Nil{},
		IfToken:        lexer.Token{},
	}
}

// StringLiteralExpressionNode: represents a string literal in the source code
// Example: "hello world" or 'test string'
type StringLiteralExpressionNode struct {
	Token lexer.Token         // The string token with its literal value
	Value objects.GoMixObject // The string object value
}

// StringLiteral.Literal(): string represenation of the node
func (node *StringLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// StringLiteral.Accept(): accepts a visitor (eg PrintVisitor)
func (node *StringLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitStringLiteralExpressionNode(*node)
}

// StringLiteral.Statement(): every expression is also a statement
func (node *StringLiteralExpressionNode) Statement() {

}

// StringLiteral.Expression(): every expression node is a node
func (node *StringLiteralExpressionNode) Expression() {

}

// NilLiteralExpressionNode: represents a nil/null literal value
// Used to represent the absence of a value or uninitialized state
type NilLiteralExpressionNode struct {
	Token lexer.Token         // The nil token
	Value objects.GoMixObject // The nil object value
}

// NullLiteral.Literal(): string represenation of the node
func (node *NilLiteralExpressionNode) Literal() string {
	return node.Token.Literal
}

// NullLiteral.Accept(): accepts a visitor (eg PrintVisitor)
func (node *NilLiteralExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitNilLiteralExpressionNode(*node)
}

// NullLiteral.Statement(): every expression is also a statement
func (node *NilLiteralExpressionNode) Statement() {

}

// NullLiteral.Expression(): every expression node is a node
func (node *NilLiteralExpressionNode) Expression() {

}

// FunctionStatementNode: represents a function definition statement
// Example: func add(x, y) { return x + y; }
type FunctionStatementNode struct {
	FuncToken  lexer.Token                 // The 'func' keyword token
	FuncName   IdentifierExpressionNode    // The function name identifier
	FuncParams []*IdentifierExpressionNode // List of parameter identifiers
	FuncBody   BlockStatementNode          // The function body block
	Value      objects.GoMixObject         // The function object value
}

// FunctionStatementNode.Literal(): string represenation of the node
func (node *FunctionStatementNode) Literal() string {

	funcParams := ""
	for _, param := range node.FuncParams {
		funcParams += param.Literal() + ","
	}
	if len(funcParams) > 0 {
		funcParams = funcParams[:len(funcParams)-1]
	}
	return node.FuncToken.Literal + " " + node.FuncName.Literal() + " (" + funcParams + ") " + node.FuncBody.Literal()
}

// FunctionStatementNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *FunctionStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitFunctionStatementNode(*node)
}

// FunctionStatementNode.Statement(): every expression is also a statement
func (node *FunctionStatementNode) Statement() {

}

// FunctionStatementNode.Expression(): every expression node is a node
func (node *FunctionStatementNode) Expression() {

}

// NewFunctionStatementNode: creates a new function statement node with default empty values
// Returns an initialized FunctionStatementNode ready for parsing function definitions
func NewFunctionStatementNode() *FunctionStatementNode {
	return &FunctionStatementNode{
		FuncToken:  lexer.Token{Type: lexer.FUNC_KEY, Literal: "func"},
		FuncName:   IdentifierExpressionNode{Name: "", Value: &objects.Nil{}},
		FuncParams: make([]*IdentifierExpressionNode, 0),
		FuncBody:   *EMPTY_BLOCK,
	}
}

// CallExpressionNode: represents a function call expression
// Example: myFunc(arg1, arg2) or print("hello")
type CallExpressionNode struct {
	FunctionIdentifier IdentifierExpressionNode // The function name being called
	Arguments          []ExpressionNode         // List of argument expressions
	Value              objects.GoMixObject      // Return value from the function
}

// CallExpressionNode.Literal(): string represenation of the node
func (node *CallExpressionNode) Literal() string {
	args := ""
	for _, arg := range node.Arguments {
		args += arg.Literal() + ","
	}
	if len(args) > 0 {
		args = args[:len(args)-1]
	}
	return node.FunctionIdentifier.Literal() + "(" + args + ")"
}

// CallExpressionNode.Accept(): accepts a visitor (eg PrintVisitor)
func (node *CallExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitCallExpressionNode(*node)
}

// CallExpressionNode.Statement(): every expression is also a statement
func (node *CallExpressionNode) Statement() {

}

// CallExpressionNode.Expression(): every expression node is a node
func (node *CallExpressionNode) Expression() {

}

// ForLoopStatementNode: represents a for loop statement with C-style syntax
// Example: for(var i=0; i<10; i=i+1) { ... }
type ForLoopStatementNode struct {
	ForToken     lexer.Token        // The 'for' keyword token
	Initializers []StatementNode    // Multiple initializers like i=0, j=0 or var i=0, j=0
	Condition    ExpressionNode     // Loop condition like i <= 10 && j <= 100
	Updates      []ExpressionNode   // Multiple updates like i=i+1, j=j+1
	Body         BlockStatementNode // The loop body containing statements
	Value        objects.GoMixObject
}

// ForLoopNode.Literal(): string representation of the node
func (node *ForLoopStatementNode) Literal() string {
	res := "for("
	// Add initializers
	for i, init := range node.Initializers {
		if i > 0 {
			res += ","
		}
		res += init.Literal()
	}
	res += ";"
	if node.Condition != nil {
		res += node.Condition.Literal()
	}
	res += ";"
	// Add updates
	for i, update := range node.Updates {
		if i > 0 {
			res += ","
		}
		res += update.Literal()
	}
	res += ")" + node.Body.Literal()
	return res
}

// ForLoopNode.Accept(): accepts a visitor
func (node *ForLoopStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitForLoopStatementNode(*node)
}

// ForLoopNode.Statement(): every for loop is a statement
func (node *ForLoopStatementNode) Statement() {

}

// WhileLoopStatementNode: represents a while loop statement with condition-based iteration
// Example: while(x > 0 && y < 100) { ... }
type WhileLoopStatementNode struct {
	WhileToken lexer.Token        // The 'while' keyword token
	Conditions []ExpressionNode   // Multiple conditions combined with logical operators
	Body       BlockStatementNode // The loop body containing statements
	Value      objects.GoMixObject
}

// WhileLoopNode.Literal(): string representation of the node
func (node *WhileLoopStatementNode) Literal() string {
	conds := ""
	for i, cond := range node.Conditions {
		if i > 0 {
			conds += " && "
		}
		conds += cond.Literal()
	}
	return "while(" + conds + ")" + node.Body.Literal()
}

// WhileLoopNode.Accept(): accepts a visitor
func (node *WhileLoopStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitWhileLoopStatementNode(*node)
}

// WhileLoopNode.Statement(): every while loop is a statement
func (node *WhileLoopStatementNode) Statement() {

}

// ArrayExpressionNode: represents an array literal expression
// Example: [1, 2, 3] or ["a", "b", "c"]
type ArrayExpressionNode struct {
	Name     IdentifierExpressionNode // Optional array identifier
	Elements []ExpressionNode         // List of element expressions
	Value    objects.GoMixObject      // The array object value
}

// ArrayExpressionNode.Literal()
func (node *ArrayExpressionNode) Literal() string {
	res := "["
	for i, elem := range node.Elements {
		if i > 0 {
			res += ","
		}
		res += elem.Literal()
	}
	res += "]"
	return res
}

// ArrayExpressionNode.Accept()
func (node *ArrayExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitArrayExpressionNode(*node)
}

// ArrayExpressionNode.Statement()
func (node *ArrayExpressionNode) Statement() {

}

// ArrayExpressionNode.Expression()
func (node *ArrayExpressionNode) Expression() {

}

// IndexExpressionNode: represents array indexing operation
// Example: arr[0], myArray[i], list[-1] (negative indexing supported)
type IndexExpressionNode struct {
	Left  ExpressionNode      // The array or indexable expression
	Index ExpressionNode      // The index expression (can be negative)
	Value objects.GoMixObject // The element value at the index
}

// IndexExpressionNode.Literal()
func (node *IndexExpressionNode) Literal() string {
	return node.Left.Literal() + "[" + node.Index.Literal() + "]"
}

// IndexExpressionNode.Accept()
func (node *IndexExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitIndexExpressionNode(*node)
}

// IndexExpressionNode.Statement()
func (node *IndexExpressionNode) Statement() {

}

// IndexExpressionNode.Expression()
func (node *IndexExpressionNode) Expression() {

}

// SliceExpressionNode: represents array slicing operation
// Example: arr[1:3], arr[:5], arr[2:] (Python-style slicing)
type SliceExpressionNode struct {
	Left  ExpressionNode      // The array or indexable expression
	Start ExpressionNode      // The start index (can be nil for arr[:end])
	End   ExpressionNode      // The end index (can be nil for arr[start:])
	Value objects.GoMixObject // The sliced array value
}

// SliceExpressionNode.Literal()
func (node *SliceExpressionNode) Literal() string {
	result := node.Left.Literal() + "["
	if node.Start != nil {
		result += node.Start.Literal()
	}
	result += ":"
	if node.End != nil {
		result += node.End.Literal()
	}
	result += "]"
	return result
}

// SliceExpressionNode.Accept()
func (node *SliceExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitSliceExpressionNode(*node)
}

// SliceExpressionNode.Statement()
func (node *SliceExpressionNode) Statement() {

}

// SliceExpressionNode.Expression()
func (node *SliceExpressionNode) Expression() {

}

// RangeExpressionNode: represents a range expression with inclusive bounds
// Example: 2...5 creates a range from 2 to 5 (inclusive)
type RangeExpressionNode struct {
	Start ExpressionNode      // The start expression of the range
	End   ExpressionNode      // The end expression of the range (inclusive)
	Value objects.GoMixObject // The Range object value
}

// RangeExpressionNode.Literal()
func (node *RangeExpressionNode) Literal() string {
	return node.Start.Literal() + "..." + node.End.Literal()
}

// RangeExpressionNode.Accept()
func (node *RangeExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitRangeExpressionNode(*node)
}

// RangeExpressionNode.Statement()
func (node *RangeExpressionNode) Statement() {

}

// RangeExpressionNode.Expression()
func (node *RangeExpressionNode) Expression() {

}

// ForeachLoopStatementNode: represents a foreach loop statement
// Example: foreach i in 2...10 { body } or foreach item in array { body }
type ForeachLoopStatementNode struct {
	ForeachToken lexer.Token              // The 'foreach' keyword token
	Iterator     IdentifierExpressionNode // The loop variable (e.g., 'i' or 'item')
	Iterable     ExpressionNode           // The range or array to iterate over
	Body         BlockStatementNode       // The loop body
	Value        objects.GoMixObject      // The result value
}

// ForeachLoopStatementNode.Literal()
func (node *ForeachLoopStatementNode) Literal() string {
	return "foreach " + node.Iterator.Name + " in " + node.Iterable.Literal() + " " + node.Body.Literal()
}

// ForeachLoopStatementNode.Accept()
func (node *ForeachLoopStatementNode) Accept(visitor NodeVisitor) {
	visitor.VisitForeachLoopStatementNode(*node)
}

// ForeachLoopStatementNode.Statement()
func (node *ForeachLoopStatementNode) Statement() {

}

// MapExpressionNode: represents a map literal expression
// Example: map{10: 20, 20: 30} or map{"name": "John", "age": 25}
type MapExpressionNode struct {
	Keys   []ExpressionNode    // List of key expressions
	Values []ExpressionNode    // List of value expressions (parallel to Keys)
	Value  objects.GoMixObject // The map object value
}

// MapExpressionNode.Literal()
func (node *MapExpressionNode) Literal() string {
	res := "map{"
	for i := range node.Keys {
		if i > 0 {
			res += ", "
		}
		res += node.Keys[i].Literal() + ": " + node.Values[i].Literal()
	}
	res += "}"
	return res
}

// MapExpressionNode.Accept()
func (node *MapExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitMapExpressionNode(*node)
}

// MapExpressionNode.Statement()
func (node *MapExpressionNode) Statement() {

}

// MapExpressionNode.Expression()
func (node *MapExpressionNode) Expression() {

}

// SetExpressionNode: represents a set literal expression
// Example: set{1, 2, 3} or set{"a", "b", "c"}
type SetExpressionNode struct {
	Elements []ExpressionNode    // List of element expressions (duplicates will be removed)
	Value    objects.GoMixObject // The set object value
}

// SetExpressionNode.Literal()
func (node *SetExpressionNode) Literal() string {
	res := "set{"
	for i, elem := range node.Elements {
		if i > 0 {
			res += ", "
		}
		res += elem.Literal()
	}
	res += "}"
	return res
}

// SetExpressionNode.Accept()
func (node *SetExpressionNode) Accept(visitor NodeVisitor) {
	visitor.VisitSetExpressionNode(*node)
}

// SetExpressionNode.Statement()
func (node *SetExpressionNode) Statement() {

}

// SetExpressionNode.Expression()
func (node *SetExpressionNode) Expression() {

}
