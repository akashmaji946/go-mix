package main

import (
	"bytes"
	"fmt"

	"github.com/akashmaji946/go-mix/parser"
)

const INDENT_SIZE = 4

func (p *PrintingVisitor) indent() {
	for i := 0; i < p.Indent; i++ {
		p.Buf.WriteString(" ")
	}
}

type PrintingVisitor struct {
	Indent int
	Buf    bytes.Buffer
}

func (p *PrintingVisitor) VisitRootNode(node parser.RootNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Root Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	for _, stmt := range node.Statements {
		stmt.Accept(p)
	}
	p.Indent -= INDENT_SIZE
}

func (p *PrintingVisitor) VisitExpressionNode(node parser.ExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Expression Node (%s)\n", node.Literal()))

}

func (p *PrintingVisitor) VisitStatementNode(node parser.StatementNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Statement Node (%s)\n", node.Literal()))

}

func (p *PrintingVisitor) VisitNumberLiteralExpressionNode(node parser.NumberLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Number Node [%s] (%s => %d)\n", node.Literal(), node.Literal(), node.Value))

}

func (p *PrintingVisitor) VisitBooleanLiteralExpressionNode(node parser.BooleanLiteralExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Boolean Node [%s] (%s => %t)\n", node.Literal(), node.Literal(), node.Value))
}

func (p *PrintingVisitor) VisitBinaryExpressionNode(node parser.BinaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Binary Node [%s] (%s => %d)\n", node.Operation.Literal, node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Left.Accept(p)
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

func (p *PrintingVisitor) VisitUnaryExpressionNode(node parser.UnaryExpressionNode) {
	p.indent()
	p.Buf.WriteString(fmt.Sprintf("Visiting Unary Node [%s] (%s => %d)\n", node.Operation.Literal, node.Literal(), node.Value))
	p.Indent += INDENT_SIZE
	node.Right.Accept(p)
	p.Indent -= INDENT_SIZE
}

func main() {

	fmt.Println("Hello, go-mix!")

	src1 := `1 + 2 * 3`
	root1 := parser.NewParser(src1).Parse()
	visitor1 := &PrintingVisitor{}
	root1.Accept(visitor1)
	fmt.Println(visitor1.Buf.String())

	src2 := `!!true`
	root2 := parser.NewParser(src2).Parse()
	visitor2 := &PrintingVisitor{}
	root2.Accept(visitor2)
	fmt.Println(visitor2.Buf.String())
}
