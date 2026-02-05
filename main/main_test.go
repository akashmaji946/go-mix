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

}
