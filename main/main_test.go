package main

import (
	"fmt"
	"testing"

	"github.com/akashmaji946/go-mix/parser"
)

func TestMain(t *testing.T) {

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

	// works but not ok. (why? a is reassigned)(will fix later)
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

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src26 := `7 > 2 + 1`
	root26 := parser.NewParser(src26).Parse()
	visitor26 := &PrintingVisitor{}
	root26.Accept(visitor26)
	fmt.Println(visitor26)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src27 := `return 7-1>2+1`
	root27 := parser.NewParser(src27).Parse()
	visitor27 := &PrintingVisitor{}
	root27.Accept(visitor27)
	fmt.Println(visitor27)

	// BOOLEAN EXPRESSION WITH RELATIONAL OPERATOR
	src28 := `var a = 7; var b = 1; var c = 2; var d = 1; return ((a-b)>(c+d));`
	root28 := parser.NewParser(src28).Parse()
	visitor28 := &PrintingVisitor{}
	root28.Accept(visitor28)
	fmt.Println(visitor28)

	// block statement
	src29 := `{var a = 10; var b = a + 100;}`
	root29 := parser.NewParser(src29).Parse()
	visitor29 := &PrintingVisitor{}
	root29.Accept(visitor29)
	fmt.Println(visitor29)

	// block statement with return statement
	src30 := `
	{
	var a = 10;
	var b = a + 10;
	var c = b * 100;
	var d =  c + 1000;
	{
	var e = d + 1000;
	return e;
	}
	return 0;
	}`
	root30 := parser.NewParser(src30).Parse()
	visitor30 := &PrintingVisitor{}
	root30.Accept(visitor30)
	fmt.Println(visitor30)

	// block statement with return statement and assignment
	src31 := `
	var X = 1234;
	{
	var a = 10;
	var b = a + 10;
	var c = b * 100;
	var d =  c + 1000;
	{
	var e = d + 1000;
	return e;
	}
	X = 111111 * 2;
	return X;
	}
	`
	root31 := parser.NewParser(src31).Parse()
	visitor31 := &PrintingVisitor{}
	root31.Accept(visitor31)
	fmt.Println(visitor31)

	// block statement with assignment
	src32 := `	
	var X = 1234;
	{
	X = X + 1;
	}
	X = X * 100 + 2;
	return X;`
	root32 := parser.NewParser(src32).Parse()
	visitor32 := &PrintingVisitor{}
	root32.Accept(visitor32)
	fmt.Println(visitor32)

	// currently block scope is overwriting the outer scope
	//  ok for now, will fix later
	src33 := `	
	var X = 1234;
	{
	var X = 6789;
	X = X + 1;
	}
	X = X + 1;
	return X;
	`
	root33 := parser.NewParser(src33).Parse()
	visitor33 := &PrintingVisitor{}
	root33.Accept(visitor33)
	fmt.Println(visitor33)

	// if expression
	src34 := `if (1 + 1 == 2) { 
	 2 + 3;
	}`
	root34 := parser.NewParser(src34).Parse()
	visitor34 := &PrintingVisitor{}
	root34.Accept(visitor34)
	fmt.Println(visitor34)

	// IF ELSE
	src35 := `if (1 + 1 == 2) { 
	 2 + 3;
	} else {
	 2 + 4;
	}`
	root35 := parser.NewParser(src35).Parse()
	visitor35 := &PrintingVisitor{}
	root35.Accept(visitor35)
	fmt.Println(visitor35)

	// IF ELSE IF ELSE
	src36 := `if (1 + 1 == 21) { 
	 2 + 3;
	} else if (2 + 2 == 4) {
	 2 + 4;
	} else {	
	 2 + 6;
	}`
	root36 := parser.NewParser(src36).Parse()
	visitor36 := &PrintingVisitor{}
	root36.Accept(visitor36)
	fmt.Println(visitor36)

	// IF ELSE IF ELSE
	src37 := `if (1 + 1 == 3) { 
	 2 + 3;
	} else if (2 + 2 == 5) {
	 2 + 4;
	} else if (3 + 3 == 7){	
	 2 + 6;
	}`
	root37 := parser.NewParser(src37).Parse()
	visitor37 := &PrintingVisitor{}
	root37.Accept(visitor37)
	fmt.Println(visitor37)

	// IF ELSE IF ELSE
	src38 := `
	var a = 100;
	var b = 0;
	if (2 * a == 200) {
		var c = 1; 
		b = 1;
	} else if (2 * a != 200) {
	 	var d = (2 * a + 2 + 2 - 2 + 3 * 5 / 2 + 1);
		b = 2;
	} else {
		b = 311111;
	}
	return b;`
	root38 := parser.NewParser(src38).Parse()
	visitor38 := &PrintingVisitor{}
	root38.Accept(visitor38)
	fmt.Println(visitor38)

	// if else if else with empty blocks
	src39 := `{
	var x = 1;
	if(x==9){}else{}
	if(x==10){}else{}
	if(x==11){}else{}
	if(x==12){}else{}
	}`
	root39 := parser.NewParser(src39).Parse()
	visitor39 := &PrintingVisitor{}
	root39.Accept(visitor39)
	fmt.Println(visitor39)

	// IF ELSE IF ELSE with empty blocks and nested blocks
	src40 := `{{{var x=1;{{{{if(x==9){}else{x+1;}}}}}}}}`
	root40 := parser.NewParser(src40).Parse()
	visitor40 := &PrintingVisitor{}
	root40.Accept(visitor40)
	fmt.Println(visitor40)

	// string literal
	src41 := `"hello world" "C++" Pascal 2026 1234567890 "123"`
	root41 := parser.NewParser(src41).Parse()
	visitor41 := &PrintingVisitor{}
	root41.Accept(visitor41)
	fmt.Println(visitor41)

	// function statement
	src42 := `func foo() { var a = 1; var b = 2; return a + b; }`
	root42 := parser.NewParser(src42).Parse()
	visitor42 := &PrintingVisitor{}
	root42.Accept(visitor42)
	fmt.Println(visitor42)

	// function statement
	src43 := `func foo(a, b, c, d) { var a = 1; var b = 2; return a * b; }`
	root43 := parser.NewParser(src43).Parse()
	visitor43 := &PrintingVisitor{}
	root43.Accept(visitor43)
	fmt.Println(visitor43)

	// function statement
	src44 := `func foo(a) { if(a==1){a=a+1;}else if(a==2){a=a+2;}else{a=a+3;} return a;}`
	root44 := parser.NewParser(src44).Parse()
	visitor44 := &PrintingVisitor{}
	root44.Accept(visitor44)
	fmt.Println(visitor44)

	// function call expression
	src45 := `foo(1, 2, 3, 4)`
	root45 := parser.NewParser(src45).Parse()
	visitor45 := &PrintingVisitor{}
	root45.Accept(visitor45)
	fmt.Println(visitor45)

	// function call expression
	src46 := `foo(1 + 2 * 3 - 8, true, (2==3), !!!!!true)`
	root46 := parser.NewParser(src46).Parse()
	visitor46 := &PrintingVisitor{}
	root46.Accept(visitor46)
	fmt.Println(visitor46)

	// function call expression with return value
	src47 := `var b = 1; var a = foo(b, 2, 3); b = a + 1;`
	root47 := parser.NewParser(src47).Parse()
	visitor47 := &PrintingVisitor{}
	root47.Accept(visitor47)
	fmt.Println(visitor47)

	// redeclaration
	src48 := `var a = 1; var a = 2; var c = a;`
	root48 := parser.NewParser(src48).Parse()
	visitor48 := &PrintingVisitor{}
	root48.Accept(visitor48)
	fmt.Println(visitor48)

	// function declaration
	src49 := `func foo(a) { if(a) {return 1;} else {return 2;}}`
	root49 := parser.NewParser(src49).Parse()
	visitor49 := &PrintingVisitor{}
	root49.Accept(visitor49)
	fmt.Println(visitor49)

	// function call
	src50 := `func fib(n) { if(n==0){return 0;} else if(n == 1) {return 1;} else {return fib(n-1) + fib(n-2);}} fib(10);`
	root50 := parser.NewParser(src50).Parse()
	visitor50 := &PrintingVisitor{}
	root50.Accept(visitor50)
	fmt.Println(visitor50)

	// function expression
	src51 := `var foo = func(a) { if(a) {return 1;} else {return 2;}};;;;foo(true);`
	root51 := parser.NewParser(src51).Parse()
	visitor51 := &PrintingVisitor{}
	root51.Accept(visitor51)
	fmt.Println(visitor51)

	// constant
	src52 := `var a = 1; const b = 2; a = 14;`
	root52 := parser.NewParser(src52).Parse()
	visitor52 := &PrintingVisitor{}
	root52.Accept(visitor52)
	fmt.Println(visitor51)

}
