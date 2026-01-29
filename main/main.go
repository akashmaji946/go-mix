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

	src13 := `var a = 1; var a = a + 10; return a;`
	root13 := parser.NewParser(src13).Parse()
	visitor13 := &PrintingVisitor{}
	root13.Accept(visitor13)
	fmt.Println(visitor13)

	// BOOLEAN EXPRESSION
	src14 := `true && false`
	root14 := parser.NewParser(src14).Parse()
	visitor14 := &PrintingVisitor{}
	root14.Accept(visitor14)
	fmt.Println(visitor14)

	// BOOLEAN EXPRESSION
	src15 := `true || false`
	root15 := parser.NewParser(src15).Parse()
	visitor15 := &PrintingVisitor{}
	root15.Accept(visitor15)
	fmt.Println(visitor15)

	// BOOLEAN EXPRESSION
	src16 := `true && (false || true)`
	root16 := parser.NewParser(src16).Parse()
	visitor16 := &PrintingVisitor{}
	root16.Accept(visitor16)
	fmt.Println(visitor16)

	// BOOLEAN EXPRESSION
	src17 := `return true && (false || true);`
	root17 := parser.NewParser(src17).Parse()
	visitor17 := &PrintingVisitor{}
	root17.Accept(visitor17)
	fmt.Println(visitor17)

	// BOOLEAN EXPRESSION
	src18 := `var a = true; var b = a && false; return b || true;`
	root18 := parser.NewParser(src18).Parse()
	visitor18 := &PrintingVisitor{}
	root18.Accept(visitor18)
	fmt.Println(visitor18)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src19 := `false || 1 < 2`
	root19 := parser.NewParser(src19).Parse()
	visitor19 := &PrintingVisitor{}
	root19.Accept(visitor19)
	fmt.Println(visitor19)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src20 := `return false || 1 < 2`
	root20 := parser.NewParser(src20).Parse()
	visitor20 := &PrintingVisitor{}
	root20.Accept(visitor20)
	fmt.Println(visitor20)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src21 := `return false || (1 <= 2 && true)`
	root21 := parser.NewParser(src21).Parse()
	visitor21 := &PrintingVisitor{}
	root21.Accept(visitor21)
	fmt.Println(visitor21)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src22 := `var a = false; return a || (10 <= 20 && true);`
	root22 := parser.NewParser(src22).Parse()
	visitor22 := &PrintingVisitor{}
	root22.Accept(visitor22)
	fmt.Println(visitor22)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src23 := `(10 <= 20 && (10 != 20) && (true != false) && (true == true))`
	root23 := parser.NewParser(src23).Parse()
	visitor23 := &PrintingVisitor{}
	root23.Accept(visitor23)
	fmt.Println(visitor23)

	// BITWISE EXPRESSION
	src24 := `(3 & 7) == 3 && true || false && true`
	root24 := parser.NewParser(src24).Parse()
	visitor24 := &PrintingVisitor{}
	root24.Accept(visitor24)
	fmt.Println(visitor24)

	// BITWISE EXPRESSION
	src25 := `return ((3&7)!=3&&true||false&&true)||true`
	root25 := parser.NewParser(src25).Parse()
	visitor25 := &PrintingVisitor{}
	root25.Accept(visitor25)
	fmt.Println(visitor25)

}
