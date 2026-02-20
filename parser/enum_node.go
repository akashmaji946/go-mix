/*
File    : go-mix/parser/enum_node.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// EnumDeclarationNode represents an enum declaration statement
// Example: enum Color { RED, GREEN, BLUE }
// Example: enum Status { PENDING = 0, ACTIVE = 1, COMPLETED = 2 }
type EnumDeclarationNode struct {
	EnumToken lexer.Token              // The 'enum' keyword token
	EnumName  IdentifierExpressionNode // The enum name identifier
	Members   []*EnumMemberNode        // List of enum members
	Value     std.GoMixObject          // The enum type object value
}

// EnumMemberNode represents a single enum member
// Example: RED or RED = 1
type EnumMemberNode struct {
	Name  string          // The member name
	Value std.GoMixObject // The member value (auto-assigned or explicit)
	Token lexer.Token     // The token for this member
}

// EnumDeclarationNode.Literal returns string representation
func (node *EnumDeclarationNode) Literal() string {
	res := node.EnumToken.Literal + " " + node.EnumName.Name + " {"
	for i, member := range node.Members {
		if i > 0 {
			res += ", "
		}
		res += member.Name
		if member.Value != nil && member.Value.GetType() != std.NilType {
			res += " = " + member.Value.ToString()
		}
	}
	res += "}"
	return res
}

// EnumDeclarationNode.Accept accepts a visitor
func (node *EnumDeclarationNode) Accept(visitor NodeVisitor) {
	// Add this method to NodeVisitor interface
	if v, ok := visitor.(interface {
		VisitEnumDeclarationNode(node EnumDeclarationNode)
	}); ok {
		v.VisitEnumDeclarationNode(*node)
	}
}

// EnumDeclarationNode.Statement marks this as a statement
func (node *EnumDeclarationNode) Statement() {}

// EnumDeclarationNode.Expression marks this as an expression
func (node *EnumDeclarationNode) Expression() {}

// EnumAccessExpressionNode represents accessing an enum member
// Example: Color.RED or Status.ACTIVE
type EnumAccessExpressionNode struct {
	EnumName   IdentifierExpressionNode // The enum name
	MemberName IdentifierExpressionNode // The member name
	Value      std.GoMixObject          // The enum member value
}

// EnumAccessExpressionNode.Literal returns string representation
func (node *EnumAccessExpressionNode) Literal() string {
	return node.EnumName.Name + "." + node.MemberName.Name
}

// EnumAccessExpressionNode.Accept accepts a visitor
func (node *EnumAccessExpressionNode) Accept(visitor NodeVisitor) {
	// Add this method to NodeVisitor interface
	if v, ok := visitor.(interface {
		VisitEnumAccessExpressionNode(node EnumAccessExpressionNode)
	}); ok {
		v.VisitEnumAccessExpressionNode(*node)
	}
}

// EnumAccessExpressionNode.Statement marks this as a statement
func (node *EnumAccessExpressionNode) Statement() {}

// EnumAccessExpressionNode.Expression marks this as an expression
func (node *EnumAccessExpressionNode) Expression() {}
