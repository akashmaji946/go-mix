/*
File    : go-mix/main/main_test.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/
package main

import (
	"fmt"
	"testing"

	"github.com/akashmaji946/go-mix/parser"
)

// TestMain_Main exercises the parser with comprehensive code samples covering all language features
func TestMain_Main(t *testing.T) {

	fmt.Println("Hello, go-mix!")

	// binary expression with operator precedence
	src1 := `1 + 2 * 3`
	root1 := parser.NewParser(src1).Parse()
	visitor1 := &PrintingVisitor{}
	root1.Accept(visitor1)
	fmt.Println(visitor1)

	// unary expression with double negation
	src2 := `!!true`
	root2 := parser.NewParser(src2).Parse()
	visitor2 := &PrintingVisitor{}
	root2.Accept(visitor2)
	fmt.Println(visitor2)

	// parenthesised expression with mixed operators
	src3 := `4-(1+2)+2+3*4/2`
	root3 := parser.NewParser(src3).Parse()
	visitor3 := &PrintingVisitor{}
	root3.Accept(visitor3)
	fmt.Println(visitor3)

	// parenthesised expression with grouped operations
	src4 := `4-(1+2)+(2+3)*4/2`
	root4 := parser.NewParser(src4).Parse()
	visitor4 := &PrintingVisitor{}
	root4.Accept(visitor4)
	fmt.Println(visitor4)

	// declarative statement with simple assignment
	src5 := `var a = 1`
	root5 := parser.NewParser(src5).Parse()
	visitor5 := &PrintingVisitor{}
	root5.Accept(visitor5)
	fmt.Println(visitor5)

	// declarative statement with parenthesised expression
	src6 := `var a = (1 + 2) * 3`
	root6 := parser.NewParser(src6).Parse()
	visitor6 := &PrintingVisitor{}
	root6.Accept(visitor6)
	fmt.Println(visitor6)

	// declarative statement using identifiers
	src7 := `var a = 11
	var b = a + 10`
	root7 := parser.NewParser(src7).Parse()
	visitor7 := &PrintingVisitor{}
	root7.Accept(visitor7)
	fmt.Println(visitor7)

	// declarative statement with multiple dependent variables
	src8 := `var a = (1 + 2) * 3
	var b = (a + 10 * 2)
	var c = (b + 10 * 4)
	`
	root8 := parser.NewParser(src8).Parse()
	visitor8 := &PrintingVisitor{}
	root8.Accept(visitor8)
	fmt.Println(visitor8)

	// declarative statement with semicolon separators
	src9 := `var a = (1 + 2) * 3;
	var b = (a + 10 * 2);
	var c = (b + 10 * 4);
	var d = (c + 10 * 5);
	`
	root9 := parser.NewParser(src9).Parse()
	visitor9 := &PrintingVisitor{}
	root9.Accept(visitor9)
	fmt.Println(visitor9)

	// return statement with literal value
	src10 := `return 1`
	root10 := parser.NewParser(src10).Parse()
	visitor10 := &PrintingVisitor{}
	root10.Accept(visitor10)
	fmt.Println(visitor10)

	// return statement with complex expression
	src11 := `return (1 + 2) * 100 + (2 * 3) * 100 + 100`
	root11 := parser.NewParser(src11).Parse()
	visitor11 := &PrintingVisitor{}
	root11.Accept(visitor11)
	fmt.Println(visitor11)

	// return statement with variable and expression
	src12 := `var a = (100 + 900); return ((a * 2) + 100 + 100 / 100);`
	root12 := parser.NewParser(src12).Parse()
	visitor12 := &PrintingVisitor{}
	root12.Accept(visitor12)
	fmt.Println(visitor12)

	// variable redeclaration test case
	src13 := `var a = 1; var a = a + 10; return a;`
	root13 := parser.NewParser(src13).Parse()
	visitor13 := &PrintingVisitor{}
	root13.Accept(visitor13)
	fmt.Println(visitor13)

	// boolean expression with AND operator
	src14 := `true && false`
	root14 := parser.NewParser(src14).Parse()
	visitor14 := &PrintingVisitor{}
	root14.Accept(visitor14)
	fmt.Println(visitor14)

	// boolean expression with OR operator
	src15 := `true || false`
	root15 := parser.NewParser(src15).Parse()
	visitor15 := &PrintingVisitor{}
	root15.Accept(visitor15)
	fmt.Println(visitor15)

	// boolean expression with parentheses
	src16 := `true && (false || true)`
	root16 := parser.NewParser(src16).Parse()
	visitor16 := &PrintingVisitor{}
	root16.Accept(visitor16)
	fmt.Println(visitor16)

	// boolean expression in return statement
	src17 := `return true && (false || true);`
	root17 := parser.NewParser(src17).Parse()
	visitor17 := &PrintingVisitor{}
	root17.Accept(visitor17)
	fmt.Println(visitor17)

	// boolean expression with variables
	src18 := `var a = true; var b = a && false; return b || true;`
	root18 := parser.NewParser(src18).Parse()
	visitor18 := &PrintingVisitor{}
	root18.Accept(visitor18)
	fmt.Println(visitor18)

	// boolean expression with relational operator
	src19 := `false || 1 < 2`
	root19 := parser.NewParser(src19).Parse()
	visitor19 := &PrintingVisitor{}
	root19.Accept(visitor19)
	fmt.Println(visitor19)

	// relational operator in return statement
	src20 := `return false || 1 < 2`
	root20 := parser.NewParser(src20).Parse()
	visitor20 := &PrintingVisitor{}
	root20.Accept(visitor20)
	fmt.Println(visitor20)

	// relational operator with boolean combination
	src21 := `return false || (1 <= 2 && true)`
	root21 := parser.NewParser(src21).Parse()
	visitor21 := &PrintingVisitor{}
	root21.Accept(visitor21)
	fmt.Println(visitor21)

	// relational operator with variable
	src22 := `var a = false; return a || (10 <= 20 && true);`
	root22 := parser.NewParser(src22).Parse()
	visitor22 := &PrintingVisitor{}
	root22.Accept(visitor22)
	fmt.Println(visitor22)

	// multiple relational and equality operators
	src23 := `(10 <= 20 && (10 != 20) && (true != false) && (true == true))`
	root23 := parser.NewParser(src23).Parse()
	visitor23 := &PrintingVisitor{}
	root23.Accept(visitor23)
	fmt.Println(visitor23)

	// bitwise expression with boolean operators
	src24 := `(3 & 7) == 3 && true || false && true`
	root24 := parser.NewParser(src24).Parse()
	visitor24 := &PrintingVisitor{}
	root24.Accept(visitor24)
	fmt.Println(visitor24)

	// complex bitwise and boolean expression
	src25 := `return ((3&7)!=3&&true||false&&true)||true`
	root25 := parser.NewParser(src25).Parse()
	visitor25 := &PrintingVisitor{}
	root25.Accept(visitor25)
	fmt.Println(visitor25)

	// relational operator with arithmetic
	src26 := `7 > 2 + 1`
	root26 := parser.NewParser(src26).Parse()
	visitor26 := &PrintingVisitor{}
	root26.Accept(visitor26)
	fmt.Println(visitor26)

	// relational operator comparing expressions
	src27 := `return 7-1>2+1`
	root27 := parser.NewParser(src27).Parse()
	visitor27 := &PrintingVisitor{}
	root27.Accept(visitor27)
	fmt.Println(visitor27)

	// relational operator with variable expressions
	src28 := `var a = 7; var b = 1; var c = 2; var d = 1; return ((a-b)>(c+d));`
	root28 := parser.NewParser(src28).Parse()
	visitor28 := &PrintingVisitor{}
	root28.Accept(visitor28)
	fmt.Println(visitor28)

	// block statement with local variables
	src29 := `{var a = 10; var b = a + 100;}`
	root29 := parser.NewParser(src29).Parse()
	visitor29 := &PrintingVisitor{}
	root29.Accept(visitor29)
	fmt.Println(visitor29)

	// nested block statements with return
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

	// block statement with outer variable assignment
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

	// block statement modifying outer variable
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

	// block statement with variable shadowing
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

	// if statement with expression condition
	src34 := `if (1 + 1 == 2) { 
	 2 + 3;
	}`
	root34 := parser.NewParser(src34).Parse()
	visitor34 := &PrintingVisitor{}
	root34.Accept(visitor34)
	fmt.Println(visitor34)

	// if-else statement
	src35 := `if (1 + 1 == 2) { 
	 2 + 3;
	} else {
	 2 + 4;
	}`
	root35 := parser.NewParser(src35).Parse()
	visitor35 := &PrintingVisitor{}
	root35.Accept(visitor35)
	fmt.Println(visitor35)

	// if-else-if chain
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

	// multiple else-if conditions
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

	// if-else-if with variable assignments
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

	// if-else with empty blocks
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

	// deeply nested blocks with if-else
	src40 := `{{{var x=1;{{{{if(x==9){}else{x+1;}}}}}}}}`
	root40 := parser.NewParser(src40).Parse()
	visitor40 := &PrintingVisitor{}
	root40.Accept(visitor40)
	fmt.Println(visitor40)

	// string literals and identifiers
	src41 := `"hello world" "C++" Pascal 2026 1234567890 "123"`
	root41 := parser.NewParser(src41).Parse()
	visitor41 := &PrintingVisitor{}
	root41.Accept(visitor41)
	fmt.Println(visitor41)

	// function declaration with return
	src42 := `func foo() { var a = 1; var b = 2; return a + b; }`
	root42 := parser.NewParser(src42).Parse()
	visitor42 := &PrintingVisitor{}
	root42.Accept(visitor42)
	fmt.Println(visitor42)

	// function with parameters
	src43 := `func foo(a, b, c, d) { var a = 1; var b = 2; return a * b; }`
	root43 := parser.NewParser(src43).Parse()
	visitor43 := &PrintingVisitor{}
	root43.Accept(visitor43)
	fmt.Println(visitor43)

	// function with conditional logic
	src44 := `func foo(a) { if(a==1){a=a+1;}else if(a==2){a=a+2;}else{a=a+3;} return a;}`
	root44 := parser.NewParser(src44).Parse()
	visitor44 := &PrintingVisitor{}
	root44.Accept(visitor44)
	fmt.Println(visitor44)

	// function call with arguments
	src45 := `foo(1, 2, 3, 4)`
	root45 := parser.NewParser(src45).Parse()
	visitor45 := &PrintingVisitor{}
	root45.Accept(visitor45)
	fmt.Println(visitor45)

	// function call with complex arguments
	src46 := `foo(1 + 2 * 3 - 8, true, (2==3), !!!!!true)`
	root46 := parser.NewParser(src46).Parse()
	visitor46 := &PrintingVisitor{}
	root46.Accept(visitor46)
	fmt.Println(visitor46)

	// function call assigned to variable
	src47 := `var b = 1; var a = foo(b, 2, 3); b = a + 1;`
	root47 := parser.NewParser(src47).Parse()
	visitor47 := &PrintingVisitor{}
	root47.Accept(visitor47)
	fmt.Println(visitor47)

	// variable redeclaration
	src48 := `var a = 1; var a = 2; var c = a;`
	root48 := parser.NewParser(src48).Parse()
	visitor48 := &PrintingVisitor{}
	root48.Accept(visitor48)
	fmt.Println(visitor48)

	// function with boolean parameter
	src49 := `func foo(a) { if(a) {return 1;} else {return 2;}}`
	root49 := parser.NewParser(src49).Parse()
	visitor49 := &PrintingVisitor{}
	root49.Accept(visitor49)
	fmt.Println(visitor49)

	// recursive fibonacci function
	src50 := `func fib(n) { if(n==0){return 0;} else if(n == 1) {return 1;} else {return fib(n-1) + fib(n-2);}} fib(10);`
	root50 := parser.NewParser(src50).Parse()
	visitor50 := &PrintingVisitor{}
	root50.Accept(visitor50)
	fmt.Println(visitor50)

	// anonymous function assigned to variable
	src51 := `var foo = func(a) { if(a) {return 1;} else {return 2;}};;;;foo(true);`
	root51 := parser.NewParser(src51).Parse()
	visitor51 := &PrintingVisitor{}
	root51.Accept(visitor51)
	fmt.Println(visitor51)

	// const and var declarations
	src52 := `var a = 1; const b = 2; a = 14;`
	root52 := parser.NewParser(src52).Parse()
	visitor52 := &PrintingVisitor{}
	root52.Accept(visitor52)
	fmt.Println(visitor52)

	// while loop with single condition
	src53 := `var i = 0; while(i < 5){ i = i + 1; }`
	root53 := parser.NewParser(src53).Parse()
	visitor53 := &PrintingVisitor{}
	root53.Accept(visitor53)
	fmt.Println(visitor53)

	// while loop with multiple conditions
	src54 := `var i = 0; var j = 10; while(i < 5, j > 5){ i = i + 1; j = j - 1; }`
	root54 := parser.NewParser(src54).Parse()
	visitor54 := &PrintingVisitor{}
	root54.Accept(visitor54)
	fmt.Println(visitor54)

	// while loop with three conditions
	src55 := `var a = 0; var b = 20; var c = 10; while(a < 10, b > 10, c > 5){ a = a + 1; b = b - 1; c = c - 1; }`
	root55 := parser.NewParser(src55).Parse()
	visitor55 := &PrintingVisitor{}
	root55.Accept(visitor55)
	fmt.Println(visitor55)

	// for loop with multiple initializers and updates
	src56 := `for(i = 0, j = 10; i < 5 && j > 5; i = i + 1, j = j - 1){ }`
	root56 := parser.NewParser(src56).Parse()
	visitor56 := &PrintingVisitor{}
	root56.Accept(visitor56)
	fmt.Println(visitor56)

	// comprehensive while loop test suite
	src57 := `
	// Test multiple conditions in while loops

	print("=== Test 1: While with 2 conditions ===")
	var i = 0
	var j = 10
	while(i < 5, j > 5){
		print("i:", i, "j:", j)
		i = i + 1
		j = j - 1
	}

	print("\n=== Test 2: While with 3 conditions ===")
	var a = 0
	var b = 20
	var c = 10
	while(a < 10, b > 10, c > 5){
		print("a:", a, "b:", b, "c:", c)
		a = a + 1
		b = b - 1
		c = c - 1
	}

	print("\n=== Test 3: While with complex conditions ===")
	var x = 0
	var y = 0
	while(x < 5, y < 10, x + y < 12){
		print("x:", x, "y:", y, "sum:", x + y)
		x = x + 1
		y = y + 2
	}

	print("\nAll while loop tests completed!")

	`
	root57 := parser.NewParser(src57).Parse()
	visitor57 := &PrintingVisitor{}
	root57.Accept(visitor57)
	fmt.Println(visitor57)

	// fibonacci with main function
	src58 := `
		var ans = 0;
		const n = 10;

		func fibb(n){
			if(n < 2){
				return n
			}
			return fibb(n-1) + fibb(n-2)
		}

		func main(){
			ans = fibb(n)
			print(ans)
		}

		main()
	`
	root58 := parser.NewParser(src58).Parse()
	visitor58 := &PrintingVisitor{}
	root58.Accept(visitor58)
	fmt.Println(visitor58)

	// let keyword with reassignment
	src59 := `let a = 1; a = 2; a = 2.9;`
	root59 := parser.NewParser(src59).Parse()
	visitor59 := &PrintingVisitor{}
	root59.Accept(visitor59)
	fmt.Println(visitor59)

	// const keyword with reassignment attempt
	src60 := `const a = 1; a = 2; a = 2.9;`
	root60 := parser.NewParser(src60).Parse()
	visitor60 := &PrintingVisitor{}
	root60.Accept(visitor60)
	fmt.Println(visitor60)

	// single-line comments
	src61 := `// this is a comment line
	// again a new comment
	// again
	// yet again
	// `
	root61 := parser.NewParser(src61).Parse()
	visitor61 := &PrintingVisitor{}
	root61.Accept(visitor61)
	fmt.Println(visitor61)

	// mixed single and multi-line comments
	src62 := `// this is a comment line
	// again a new comment
	/* haha I a comment too */
	// now a line
	(2 + 3);
	(3 + 9);`
	root62 := parser.NewParser(src62).Parse()
	visitor62 := &PrintingVisitor{}
	root62.Accept(visitor62)
	fmt.Println(visitor62)

	// arithmetic compound assignments
	src64 := `var a = 10; a += 5; a -= 3; a *= 2; a /= 4; a %= 3`
	root64 := parser.NewParser(src64).Parse()
	visitor64 := &PrintingVisitor{}
	root64.Accept(visitor64)
	fmt.Println(visitor64)

	// bitwise compound assignments
	src65 := `var a = 12; a &= 10; a |= 3; a ^= 5`
	root65 := parser.NewParser(src65).Parse()
	visitor65 := &PrintingVisitor{}
	root65.Accept(visitor65)
	fmt.Println(visitor65)

	// shift compound assignments
	src66 := `var a = 4; a <<= 2; a >>= 1`
	root66 := parser.NewParser(src66).Parse()
	visitor66 := &PrintingVisitor{}
	root66.Accept(visitor66)
	fmt.Println(visitor66)

	// compound assignment in for loop
	src67 := `for(i = 0; i < 10; i += 2){ var x = i * 2; }`
	root67 := parser.NewParser(src67).Parse()
	visitor67 := &PrintingVisitor{}
	root67.Accept(visitor67)
	fmt.Println(visitor67)

	// compound assignment in while loop
	src68 := `var sum = 0; var i = 1; while(i <= 5){ sum += i; i += 1; }`
	root68 := parser.NewParser(src68).Parse()
	visitor68 := &PrintingVisitor{}
	root68.Accept(visitor68)
	fmt.Println(visitor68)

	// multiple compound assignments in for loop
	src69 := `for(i = 0, j = 10; i < 5; i += 1, j -= 2){ var diff = j - i; }`
	root69 := parser.NewParser(src69).Parse()
	visitor69 := &PrintingVisitor{}
	root69.Accept(visitor69)
	fmt.Println(visitor69)

	// compound assignment with complex expressions
	src70 := `var a = 100; a += 2 * 3 + 4; a -= (10 - 5); a *= 2`
	root70 := parser.NewParser(src70).Parse()
	visitor70 := &PrintingVisitor{}
	root70.Accept(visitor70)
	fmt.Println(visitor70)

	// complex nested loop patterns
	src71 := `
	// Complex loop patterns with variable declarations

	// Test 1: Matrix-like operations (2D array simulation)
	print("=== Test 1: Matrix Operations ===")
	var rows = 3
	var cols = 3
	var matrix_sum = 0

	var row = 0
	for(row = 0; row < rows; row = row + 1) {
		var row_product = 1
		var col = 0
		for(col = 0; col < cols; col = col + 1) {
			var cell = row * cols + col + 1
			var cell_squared = cell * cell
			matrix_sum = matrix_sum + cell
			row_product = row_product * cell
			print("Cell[", row, ",", col, "] =", cell, "squared =", cell_squared)
		}
		print("Row", row, "product:", row_product)
	}
	print("Matrix sum:", matrix_sum)

	// Test 2: Fibonacci-like pattern using loops
	print("\n=== Test 2: Fibonacci Pattern ===")
	var fib_count = 10
	var prev = 0
	var curr = 1

	var fib_i = 0
	for(fib_i = 0; fib_i < fib_count; fib_i = fib_i + 1) {
		print("Fib[", fib_i, "] =", prev)
		var next = prev + curr
		prev = curr
		curr = next
	}

	// Test 3: Countdown loops with nested operations
	print("\n=== Test 3: Countdown Loops ===")
	var outer_count = 5
	while(outer_count > 0) {
		var countdown_value = outer_count * 10
		print("Countdown:", outer_count, "value:", countdown_value)
		
		var inner_count = outer_count
		while(inner_count > 0) {
			var inner_value = inner_count * 2
			print("  Inner countdown:", inner_count, "value:", inner_value)
			inner_count = inner_count - 1
		}
		outer_count = outer_count - 1
	}

	`
	root71 := parser.NewParser(src71).Parse()
	visitor71 := &PrintingVisitor{}
	root71.Accept(visitor71)
	fmt.Println(visitor71)

	// array with function element
	src72 := `var a = [1, 2, func(){2+3;}]; var b = a[2]; b();`
	root72 := parser.NewParser(src72).Parse()
	visitor72 := &PrintingVisitor{}
	root72.Accept(visitor72)
	fmt.Println(visitor72)

	// chained function and array access
	src73 := `a()[b()+1]`
	root73 := parser.NewParser(src73).Parse()
	visitor73 := &PrintingVisitor{}
	root73.Accept(visitor73)
	fmt.Println(visitor73)

	// array with mixed types and negative index
	src74 := `var arr = [1, 2, true, 2.5, func () { }];  var foo = arr[-1]; foo();`
	root74 := parser.NewParser(src74).Parse()
	visitor74 := &PrintingVisitor{}
	root74.Accept(visitor74)
	fmt.Println(visitor74)

	// array with builtin functions (push, pop, length)
	src75 := `var x = 1; 
	var y = 1 + x; 
	var arr = [x + y, x, y * 2.0, 1, 2, true, 2.5, func () {  while(true){} }];  
	var foo = arr[-1]; 
	foo(); 
	arr = push(arr, 4); 
	pop(arr); 
	length(arr);`
	root75 := parser.NewParser(src75).Parse()
	visitor75 := &PrintingVisitor{}
	root75.Accept(visitor75)
	fmt.Println(visitor75)

	// Test 76: Simple range expression
	src76 := `var r = 2...5; r`
	root76 := parser.NewParser(src76).Parse()
	visitor76 := &PrintingVisitor{}
	root76.Accept(visitor76)
	fmt.Println(visitor76)

	// Test 77: Foreach loop with range
	src77 := `var sum = 0; foreach i in 1...5 { sum = sum + i; } sum`
	root77 := parser.NewParser(src77).Parse()
	visitor77 := &PrintingVisitor{}
	root77.Accept(visitor77)
	fmt.Println(visitor77)

	// Test 78: Foreach loop with array
	src78 := `var total = 0; foreach num in [10, 20, 30, 40] { total = total + num; } total`
	root78 := parser.NewParser(src78).Parse()
	visitor78 := &PrintingVisitor{}
	root78.Accept(visitor78)
	fmt.Println(visitor78)

	// Test 79: Nested foreach loops with ranges
	src79 := `
	var result = 0;
	foreach i in 1...3 {
		foreach j in 1...3 {
			result = result + i * 10 + j;
		}
	}
	result
	`
	root79 := parser.NewParser(src79).Parse()
	visitor79 := &PrintingVisitor{}
	root79.Accept(visitor79)
	fmt.Println(visitor79)

	// Test 80: Range builtin function and foreach with complex operations
	src80 := `
	// Test range() builtin and foreach with various operations
	
	// Create range using builtin
	var r1 = range(1, 5);
	
	// Use foreach to calculate factorial-like sum
	var factorial_sum = 0;
	var multiplier = 1;
	foreach n in r1 {
		factorial_sum = factorial_sum + n * multiplier;
		multiplier = multiplier + 1;
	}
	
	// Nested foreach with range and array
	var matrix_result = 0;
	foreach row in 1...2 {
		var row_values = [10, 20, 30];
		foreach col_val in row_values {
			matrix_result = matrix_result + row * col_val;
		}
	}
	
	// Foreach with range variable
	var my_range = 5...8;
	var range_sum = 0;
	foreach val in my_range {
		range_sum = range_sum + val;
	}
	
	// Return combined result
	factorial_sum + matrix_result + range_sum
	`
	root80 := parser.NewParser(src80).Parse()
	visitor80 := &PrintingVisitor{}
	root80.Accept(visitor80)
	fmt.Println(visitor80)

	// Test 81: Map creation and access
	src81 := `var m = map{"name": "John", "age": 25}; m["name"]`
	root81 := parser.NewParser(src81).Parse()
	visitor81 := &PrintingVisitor{}
	root81.Accept(visitor81)
	fmt.Println(visitor81)

	// Test 82: Map with keys_map builtin
	src82 := `var person = map{"name": "Alice", "age": 30, "city": "NYC"}; keys_map(person)`
	root82 := parser.NewParser(src82).Parse()
	visitor82 := &PrintingVisitor{}
	root82.Accept(visitor82)
	fmt.Println(visitor82)

	// Test 83: Map with insert_map and remove_map
	src83 := `
	var config = map{"debug": false, "port": 8080};
	insert_map(config, "host", "localhost");
	remove_map(config, "debug");
	config
	`
	root83 := parser.NewParser(src83).Parse()
	visitor83 := &PrintingVisitor{}
	root83.Accept(visitor83)
	fmt.Println(visitor83)

	// Test 84: Map with contain_map
	src84 := `
	var settings = map{"theme": "dark", "lang": "en"};
	var has_theme = contain_map(settings, "theme");
	var has_font = contain_map(settings, "font");
	has_theme
	`
	root84 := parser.NewParser(src84).Parse()
	visitor84 := &PrintingVisitor{}
	root84.Accept(visitor84)
	fmt.Println(visitor84)

	// Test 85: Map with enumerate_map (NEW)
	src85 := `
	var data = map{"x": 10, "y": 20, "z": 30};
	var pairs = enumerate_map(data);
	pairs
	`
	root85 := parser.NewParser(src85).Parse()
	visitor85 := &PrintingVisitor{}
	root85.Accept(visitor85)
	fmt.Println(visitor85)

	// Test 86: Set creation and operations
	src86 := `var s = set{1, 2, 3, 4, 5}; s`
	root86 := parser.NewParser(src86).Parse()
	visitor86 := &PrintingVisitor{}
	root86.Accept(visitor86)
	fmt.Println(visitor86)

	// Test 87: Set with insert_set and remove_set
	src87 := `
	var numbers = set{10, 20, 30};
	insert_set(numbers, 40);
	remove_set(numbers, 20);
	numbers
	`
	root87 := parser.NewParser(src87).Parse()
	visitor87 := &PrintingVisitor{}
	root87.Accept(visitor87)
	fmt.Println(visitor87)

	// Test 88: Set with contains_set
	src88 := `
	var tags = set{"golang", "programming", "tutorial"};
	var has_golang = contains_set(tags, "golang");
	var has_python = contains_set(tags, "python");
	has_golang
	`
	root88 := parser.NewParser(src88).Parse()
	visitor88 := &PrintingVisitor{}
	root88.Accept(visitor88)
	fmt.Println(visitor88)

	// Test 89: Set with values_set (NEW)
	src89 := `
	var unique_ids = set{101, 102, 103, 104};
	var id_array = values_set(unique_ids);
	id_array
	`
	root89 := parser.NewParser(src89).Parse()
	visitor89 := &PrintingVisitor{}
	root89.Accept(visitor89)
	fmt.Println(visitor89)

	// Test 90: Complex map and set operations
	src90 := `
	// Build a map from set values
	var source_set = set{"a", "b", "c"};
	var set_vals = values_set(source_set);
	var result_map = map{};
	insert_map(result_map, set_vals[0], 1);
	insert_map(result_map, set_vals[1], 2);
	insert_map(result_map, set_vals[2], 3);
	
	// Enumerate the map
	var map_pairs = enumerate_map(result_map);
	
	// Build a set from map keys
	var map_keys = keys_map(result_map);
	var result_set = set{};
	insert_set(result_set, map_keys[0]);
	insert_set(result_set, map_keys[1]);
	insert_set(result_set, map_keys[2]);
	
	result_set
	`
	root90 := parser.NewParser(src90).Parse()
	visitor90 := &PrintingVisitor{}
	root90.Accept(visitor90)
	fmt.Println(visitor90)

	// Test 91: List creation and basic operations
	src91 := `var l = list(1, 2, 3, 4, 5); l`
	root91 := parser.NewParser(src91).Parse()
	visitor91 := &PrintingVisitor{}
	root91.Accept(visitor91)
	fmt.Println(visitor91)

	// Test 92: List with push and pop operations
	src92 := `
	var nums = list(10, 20, 30);
	pushback_list(nums, 40);
	pushfront_list(nums, 5);
	var last = popback_list(nums);
	var first = popfront_list(nums);
	nums
	`
	root92 := parser.NewParser(src92).Parse()
	visitor92 := &PrintingVisitor{}
	root92.Accept(visitor92)
	fmt.Println(visitor92)

	// Test 93: List indexing and slicing
	src93 := `
	var data = list(0, 10, 20, 30, 40, 50);
	var elem = data[2];
	var slice = data[1:4];
	slice
	`
	root93 := parser.NewParser(src93).Parse()
	visitor93 := &PrintingVisitor{}
	root93.Accept(visitor93)
	fmt.Println(visitor93)

	// Test 94: List with foreach loop
	src94 := `
	var sum = 0;
	foreach num in list(1, 2, 3, 4, 5) {
		sum = sum + num;
	}
	sum
	`
	root94 := parser.NewParser(src94).Parse()
	visitor94 := &PrintingVisitor{}
	root94.Accept(visitor94)
	fmt.Println(visitor94)

	// Test 95: Nested lists
	src95 := `
	var matrix = list(list(1, 2, 3), list(4, 5, 6), list(7, 8, 9));
	var row = matrix[1];
	var elem = matrix[0][2];
	elem
	`
	root95 := parser.NewParser(src95).Parse()
	visitor95 := &PrintingVisitor{}
	root95.Accept(visitor95)
	fmt.Println(visitor95)

	// Test 96: Tuple creation and basic operations
	src96 := `var t = tuple(10, 20, 30); t`
	root96 := parser.NewParser(src96).Parse()
	visitor96 := &PrintingVisitor{}
	root96.Accept(visitor96)
	fmt.Println(visitor96)

	// Test 97: Tuple indexing and slicing
	src97 := `
	var coords = tuple(100, 200, 300, 400);
	var x = coords[0];
	var slice = coords[1:3];
	slice
	`
	root97 := parser.NewParser(src97).Parse()
	visitor97 := &PrintingVisitor{}
	root97.Accept(visitor97)
	fmt.Println(visitor97)

	// Test 98: Tuple with foreach loop
	src98 := `
	var product = 1;
	foreach val in tuple(2, 3, 4) {
		product = product * val;
	}
	product
	`
	root98 := parser.NewParser(src98).Parse()
	visitor98 := &PrintingVisitor{}
	root98.Accept(visitor98)
	fmt.Println(visitor98)

	// Test 99: Tuple immutability and peek operations
	src99 := `
	var person = tuple("Alice", 25, true);
	var name = peekfront_tuple(person);
	var active = peekback_tuple(person);
	var size = size_tuple(person);
	size
	`
	root99 := parser.NewParser(src99).Parse()
	visitor99 := &PrintingVisitor{}
	root99.Accept(visitor99)
	fmt.Println(visitor99)

	// Test 100: Mixed list and tuple operations
	src100 := `
	// Create a list of tuples
	var points = list(tuple(0, 0), tuple(10, 20), tuple(30, 40));
	var second_point = points[1];
	var y_coord = second_point[1];
	
	// Create a tuple of lists
	var data = tuple(list(1, 2, 3), list(4, 5, 6));
	var first_list = data[0];
	pushback_list(first_list, 99);
	
	first_list
	`
	root100 := parser.NewParser(src100).Parse()
	visitor100 := &PrintingVisitor{}
	root100.Accept(visitor100)
	fmt.Println(visitor100)

	// Test 101: List with length and size functions
	src101 := `
	var items = list("a", "b", "c", "d", "e");
	var len1 = length(items);
	var len2 = size_list(items);
	len1
	`
	root101 := parser.NewParser(src101).Parse()
	visitor101 := &PrintingVisitor{}
	root101.Accept(visitor101)
	fmt.Println(visitor101)

	// Test 102: Tuple with length function
	src102 := `
	var rgb = tuple(255, 128, 64);
	var len = length(rgb);
	len
	`
	root102 := parser.NewParser(src102).Parse()
	visitor102 := &PrintingVisitor{}
	root102.Accept(visitor102)
	fmt.Println(visitor102)

	// Test 103: List index assignment
	src103 := `
	var arr = list(1, 2, 3, 4, 5);
	arr[0] = 100;
	arr[2] = 300;
	arr[-1] = 999;
	arr
	`
	root103 := parser.NewParser(src103).Parse()
	visitor103 := &PrintingVisitor{}
	root103.Accept(visitor103)
	fmt.Println(visitor103)

	// Test 104: Complex list manipulation
	src104 := `
	var stack = list();
	pushback_list(stack, 10);
	pushback_list(stack, 20);
	pushback_list(stack, 30);
	var top = peekback_list(stack);
	var popped = popback_list(stack);
	var new_size = size_list(stack);
	new_size
	`
	root104 := parser.NewParser(src104).Parse()
	visitor104 := &PrintingVisitor{}
	root104.Accept(visitor104)
	fmt.Println(visitor104)

	// Test 105: Tuple as function return value simulation
	src105 := `
	var result = tuple(42, "success", true);
	var code = result[0];
	var message = result[1];
	var status = result[2];
	status
	`
	root105 := parser.NewParser(src105).Parse()
	visitor105 := &PrintingVisitor{}
	root105.Accept(visitor105)
	fmt.Println(visitor105)

	// Test 106: List and tuple with nested structures
	src106 := `
	var complex_list = list(tuple(1, 2), tuple(3, 4), tuple(5, 6));
	var first_tuple = complex_list[0];
	var second_elem = first_tuple[1];
	second_elem
	`
	root106 := parser.NewParser(src106).Parse()
	visitor106 := &PrintingVisitor{}
	root106.Accept(visitor106)
	fmt.Println(visitor106)

	// Test 107: Empty list and tuple
	src107 := `var empty_list = list(); var empty_tuple = tuple(); empty_list; empty_tuple;`
	root107 := parser.NewParser(src107).Parse()
	visitor107 := &PrintingVisitor{}
	root107.Accept(visitor107)
	fmt.Println(visitor107)

	// Test 108: List and tuple with different data types
	src108 := `
	var mixed_list = list(1, "two", 3.0, true, func() { return "hello"; });
	var mixed_tuple = tuple("start", 42, false, func() { return "world"; });
	mixed_list; mixed_tuple;
	`
	root108 := parser.NewParser(src108).Parse()
	visitor108 := &PrintingVisitor{}
	root108.Accept(visitor108)
	fmt.Println(visitor108)

	// Test 109: struct creation
	src109 := `var Point = struct Data {}`
	root109 := parser.NewParser(src109).Parse()
	visitor109 := &PrintingVisitor{}
	root109.Accept(visitor109)
	fmt.Println(visitor109)

	// Test 110: struct with init
	src110 := `var Point = struct Data { func init(){} }`
	root110 := parser.NewParser(src110).Parse()
	visitor110 := &PrintingVisitor{}
	root110.Accept(visitor110)
	fmt.Println(visitor110)

	// Test 111: struct with other methods
	src111 := `struct Data { func init(){} func move(x, y){} }`
	root111 := parser.NewParser(src111).Parse()
	visitor111 := &PrintingVisitor{}
	root111.Accept(visitor111)
	fmt.Println(visitor111)

	// Test 112: new keyword with struct
	src112 := `struct Data {}; var p = new Point(10, 20);`
	root112 := parser.NewParser(src112).Parse()
	visitor112 := &PrintingVisitor{}
	root112.Accept(visitor112)
	fmt.Println(visitor112)

	// Test 113: struct with fields and methods
	src113 := `
	struct Point {
		func init(x, y) {
			this.x = x;
			this.y = y;
		}
		func move(dx, dy) {
			this.x = this.x + dx;
			this.y = this.y + dy;
		}
	}
	var p = new Point(10, 20);
	p.move(5, -5);
	p
	`
	root113 := parser.NewParser(src113).Parse()
	visitor113 := &PrintingVisitor{}
	root113.Accept(visitor113)
	fmt.Println(visitor113)

	src114 := `
	// Struct with new constructor
	var aa = 10;
	struct A { 
		func init() { 
			var dd = 12;
			println(aa);
			println("In the constructor"); 
		} 
	}
	var a = new A();
	println(a);
	`
	root114 := parser.NewParser(src114).Parse()
	visitor114 := &PrintingVisitor{}
	root114.Accept(visitor114)
	fmt.Println(visitor114)

	// Test 115: Struct with multiple methods and field access
	src115 := `struct A{ func init(){} func hello(){ return 12 } } var a=new A(); var res=a.hello(); println(res);`
	root115 := parser.NewParser(src115).Parse()
	visitor115 := &PrintingVisitor{}
	root115.Accept(visitor115)
	fmt.Println(visitor115)

	// Test 116: Struct with nested struct
	src116 := `
	struct C { func add(x, y) { return x + y } }
	var c = new C();
	var d = c.add(3, 4);
	println(d);
	`
	root116 := parser.NewParser(src116).Parse()
	visitor116 := &PrintingVisitor{}
	root116.Accept(visitor116)
	fmt.Println(visitor116)

	// Test 117: struct with new and calling method
	src117 := `
	struct Point {
		func init(x, y) {
			this.x = x;
			this.y = y;
		}
		func sum() {
			return this.x + this.y;
		}
	}
	var p = new Point(10, 20);
	p.sum();
	`
	root117 := parser.NewParser(src117).Parse()
	visitor117 := &PrintingVisitor{}
	root117.Accept(visitor117)
	fmt.Println(visitor117)

	// Test 118: struct with new and field access
	src118 := `
	func foo(){
		var m = [[0,0], [1, 1], [2, 2]]
		return m;
	}

	struct Test{
		func init() {
			this.m = foo()
		}
		func getter(x){
			return this.m[x]
		}
		func setter(x, y){
			this.m[x] = y;
		
		}
	}

	var t = new Test();
	println(t.getter(2))
	t.setter(2, 20);
	println(t.getter(2))
	t.setter(1, 10);
	println(t.getter(1))
	t.setter(0, map{1:11, 2:22})
	println(t.getter(0)[1])
	`
	root118 := parser.NewParser(src118).Parse()
	visitor118 := &PrintingVisitor{}
	root118.Accept(visitor118)
	fmt.Println(visitor118)

	src119 := `
	struct Logic { func run(x) { if (x > 0) { while(x > 0) { x = x - 1; } } return x; } }
	`
	root119 := parser.NewParser(src119).Parse()
	visitor119 := &PrintingVisitor{}
	root119.Accept(visitor119)
	fmt.Println(visitor119)

	src120 := `
	struct S { func init(v) { this.v = v; } func get() { return this.v; } } var s = new S(5); s.get()
	`
	root120 := parser.NewParser(src120).Parse()
	visitor120 := &PrintingVisitor{}
	root120.Accept(visitor120)
	fmt.Println(visitor120)

}
