package main

import (
	"fmt"

	"github.com/akashmaji946/go-mix/parser"
)

func main() {

	fmt.Println("Hello, go-mix!")

	// binary expression
	src1 := `1 + 2 * 3`
	root1 := parser.NewParser(src1).Parse()
	visitor1 := &PrintingVisitor{}
	root1.Accept(visitor1)
	fmt.Println(visitor1)

	// unary expression
	src2 := `!!true`
	root2 := parser.NewParser(src2).Parse()
	visitor2 := &PrintingVisitor{}
	root2.Accept(visitor2)
	fmt.Println(visitor2)

	// parenthesised expression
	src3 := `4-(1+2)+2+3*4/2`
	root3 := parser.NewParser(src3).Parse()
	visitor3 := &PrintingVisitor{}
	root3.Accept(visitor3)
	fmt.Println(visitor3)

	// parenthesised expression
	src4 := `4-(1+2)+(2+3)*4/2`
	root4 := parser.NewParser(src4).Parse()
	visitor4 := &PrintingVisitor{}
	root4.Accept(visitor4)
	fmt.Println(visitor4)

	// declarative statement (simple)
	src5 := `var a = 1`
	root5 := parser.NewParser(src5).Parse()
	visitor5 := &PrintingVisitor{}
	root5.Accept(visitor5)
	fmt.Println(visitor5)

	// declarative statement (with parenthesised expression)
	src6 := `var a = (1 + 2) * 3`
	root6 := parser.NewParser(src6).Parse()
	visitor6 := &PrintingVisitor{}
	root6.Accept(visitor6)
	fmt.Println(visitor6)

	// declarative statement (with identifiers)
	src7 := `var a = 11
	var b = a + 10`
	root7 := parser.NewParser(src7).Parse()
	visitor7 := &PrintingVisitor{}
	root7.Accept(visitor7)
	fmt.Println(visitor7)

	// declarative statement (with multiple statements)
	src8 := `var a = (1 + 2) * 3
	var b = (a + 10 * 2)
	var c = (b + 10 * 4)
	`
	root8 := parser.NewParser(src8).Parse()
	visitor8 := &PrintingVisitor{}
	root8.Accept(visitor8)
	fmt.Println(visitor8)

	// declarative statement (with semicolons)
	src9 := `var a = (1 + 2) * 3;
	var b = (a + 10 * 2);
	var c = (b + 10 * 4);
	var d = (c + 10 * 5);
	`
	root9 := parser.NewParser(src9).Parse()
	visitor9 := &PrintingVisitor{}
	root9.Accept(visitor9)
	fmt.Println(visitor9)

	// return statement
	src10 := `return 1`
	root10 := parser.NewParser(src10).Parse()
	visitor10 := &PrintingVisitor{}
	root10.Accept(visitor10)
	fmt.Println(visitor10)

	// return statement
	src11 := `return (1 + 2) * 100 + (2 * 3) * 100 + 100`
	root11 := parser.NewParser(src11).Parse()
	visitor11 := &PrintingVisitor{}
	root11.Accept(visitor11)
	fmt.Println(visitor11)

	// return statement
	src12 := `var a = (100 + 900); return ((a * 2) + 100 + 100 / 100);`
	root12 := parser.NewParser(src12).Parse()
	visitor12 := &PrintingVisitor{}
	root12.Accept(visitor12)
	fmt.Println(visitor12)

}
