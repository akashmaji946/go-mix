/*
File    : go-mix/parser/enum_parser.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package parser

import (
	"fmt"
	"strconv"

	"github.com/akashmaji946/go-mix/lexer"
	"github.com/akashmaji946/go-mix/std"
)

// parseEnumDeclaration parses enum declaration statements.
//
// Syntax:
//
//	enum EnumName { MEMBER1, MEMBER2, MEMBER3 }
//	enum Status { PENDING = 0, ACTIVE = 1, COMPLETED = 2 }
//
// Returns:
//
//	An EnumDeclarationNode (as ExpressionNode to satisfy unaryParseFunction)
//
// Examples:
//
//	enum Color { RED, GREEN, BLUE }
//	enum Priority { LOW = 1, MEDIUM = 5, HIGH = 10 }
func (par *Parser) parseEnumDeclaration() ExpressionNode {
	enumToken := par.CurrToken

	// Expect enum name
	if !par.expectAdvance(lexer.IDENTIFIER_ID) {
		return nil
	}
	enumName := IdentifierExpressionNode{
		Token: par.CurrToken,
		Name:  par.CurrToken.Literal,
		Value: &std.Nil{},
	}

	// Expect opening brace
	if !par.expectAdvance(lexer.LEFT_BRACE) {
		return nil
	}

	// Parse enum members
	members := make([]*EnumMemberNode, 0)
	autoValue := int64(0)

	// Check for empty enum
	if par.NextToken.Type == lexer.RIGHT_BRACE {
		par.advance() // Move to }
		return &EnumDeclarationNode{
			EnumToken: enumToken,
			EnumName:  enumName,
			Members:   members,
			Value:     &std.Nil{},
		}
	}

	par.advance() // Move to first member

	for par.CurrToken.Type != lexer.RIGHT_BRACE && par.CurrToken.Type != lexer.EOF_TYPE {
		// Expect identifier for member name
		if par.CurrToken.Type != lexer.IDENTIFIER_ID {
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected identifier for enum member, got %s",
				par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Type)
			par.addError(msg)
			return nil
		}

		memberName := par.CurrToken.Literal
		memberToken := par.CurrToken
		var memberValue std.GoMixObject

		// Check for explicit value assignment
		if par.NextToken.Type == lexer.ASSIGN_OP {
			par.advance() // Move to =
			par.advance() // Move past = to value

			// Parse the value expression
			if par.CurrToken.Type == lexer.INT_LIT {
				val, err := strconv.ParseInt(par.CurrToken.Literal, 0, 64)
				if err != nil {
					msg := fmt.Sprintf("[%d:%d] PARSER ERROR: invalid integer value for enum member: %s",
						par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Literal)
					par.addError(msg)
					return nil
				}
				memberValue = &std.Integer{Value: val}
				autoValue = val + 1 // Next auto value continues from this
			} else {
				msg := fmt.Sprintf("[%d:%d] PARSER ERROR: enum member value must be an integer, got %s",
					par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Type)
				par.addError(msg)
				return nil
			}
		} else {
			// Auto-assign value
			memberValue = &std.Integer{Value: autoValue}
			autoValue++
		}

		// Create member node
		member := &EnumMemberNode{
			Name:  memberName,
			Value: memberValue,
			Token: memberToken,
		}
		members = append(members, member)

		// Check for comma or end of enum
		if par.NextToken.Type == lexer.COMMA_DELIM {
			par.advance() // Move to ,
			par.advance() // Move past , to next member
		} else if par.NextToken.Type == lexer.RIGHT_BRACE {
			par.advance() // Move to }
			break
		} else if par.NextToken.Type == lexer.EOF_TYPE {
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: unexpected end of file in enum declaration",
				par.CurrToken.Line, par.CurrToken.Column)
			par.addError(msg)
			return nil
		} else {
			msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected , or }, got %s",
				par.NextToken.Line, par.NextToken.Column, par.NextToken.Type)
			par.addError(msg)
			return nil
		}
	}

	return &EnumDeclarationNode{
		EnumToken: enumToken,
		EnumName:  enumName,
		Members:   members,
		Value:     &std.Nil{},
	}
}

// parseEnumAccessExpression parses enum member access expressions.
// This handles accessing enum members like Color.RED or Status.ACTIVE
//
// Parameters:
//
//	left - The enum name (identifier)
//
// Returns:
//
//	An EnumAccessExpressionNode
//
// Examples:
//
//	Color.RED
//	Status.ACTIVE
//	Priority.HIGH
func (par *Parser) parseEnumAccessExpression(left ExpressionNode) ExpressionNode {
	// Current token is DOT_OP
	par.advance() // Move past .

	// Expect member name identifier
	if par.CurrToken.Type != lexer.IDENTIFIER_ID {
		msg := fmt.Sprintf("[%d:%d] PARSER ERROR: expected enum member name, got %s",
			par.CurrToken.Line, par.CurrToken.Column, par.CurrToken.Type)
		par.addError(msg)
		return nil
	}

	enumName := ""
	if ident, ok := left.(*IdentifierExpressionNode); ok {
		enumName = ident.Name
	} else if bin, ok := left.(*BinaryExpressionNode); ok {
		// Handle chained access like Enum.MEMBER.VALUE
		if ident, ok := bin.Right.(*IdentifierExpressionNode); ok {
			enumName = ident.Name
		}
	}

	memberName := par.CurrToken.Literal

	return &EnumAccessExpressionNode{
		EnumName: IdentifierExpressionNode{
			Token: par.CurrToken,
			Name:  enumName,
			Value: &std.Nil{},
		},
		MemberName: IdentifierExpressionNode{
			Token: par.CurrToken,
			Name:  memberName,
			Value: &std.Nil{},
		},
		Value: &std.Nil{},
	}
}
