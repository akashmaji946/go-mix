package main

import (
	"fmt"
	"testing"

	"github.com/akashmaji946/go-mix/parser"
)

func runParseTest(t *testing.T, src string) {
	t.Helper()
	root := parser.NewParser(src).Parse()
	visitor := &PrintingVisitor{}
	root.Accept(visitor)
	fmt.Println(visitor)
}

func TestSrc72(t *testing.T) {
	src := `var a = [1, 2, func(){2+3;}]; var b = a[2]; b();`
	runParseTest(t, src)
}

// binary expression with operator precedence

func TestSrc1(t *testing.T) {
	src := `1 + 2 * 3`
	runParseTest(t, src)
}

// unary expression with double negation

func TestSrc2(t *testing.T) {
	src := `!!true`
	runParseTest(t, src)
}

// parenthesised expression with mixed operators

func TestSrc3(t *testing.T) {
	src := `4-(1+2)+2+3*4/2`
	runParseTest(t, src)
}

// parenthesised expression with grouped operations

func TestSrc4(t *testing.T) {
	src := `4-(1+2)+(2+3)*4/2`
	runParseTest(t, src)
}

// declarative statement with simple assignment

func TestSrc5(t *testing.T) {
	src := `var a = 1`
	runParseTest(t, src)
}

// declarative statement with parenthesised expression

func TestSrc6(t *testing.T) {
	src := `var a = (1 + 2) * 3`
	runParseTest(t, src)
}

// declarative statement using identifiers

func TestSrc7(t *testing.T) {
	src := `var a = 11
	var b = a + 10`
	runParseTest(t, src)
}

// declarative statement with multiple dependent variables

func TestSrc8(t *testing.T) {
	src := `var a = (1 + 2) * 3
	var b = (a + 10 * 2)
	var c = (b + 10 * 4)
	`
	runParseTest(t, src)
}

// declarative statement with semicolon separators

func TestSrc9(t *testing.T) {
	src := `var a = (1 + 2) * 3;
	var b = (a + 10 * 2);
	var c = (b + 10 * 4);
	var d = (c + 10 * 5);
	`
	runParseTest(t, src)
}

// return statement with literal value

func TestSrc10(t *testing.T) {
	src := `return 1`
	runParseTest(t, src)
}

// return statement with complex expression

func TestSrc11(t *testing.T) {
	src := `return (1 + 2) * 100 + (2 * 3) * 100 + 100`
	runParseTest(t, src)
}

// return statement with variable and expression

func TestSrc12(t *testing.T) {
	src := `var a = (100 + 900); return ((a * 2) + 100 + 100 / 100);`
	runParseTest(t, src)
}

// variable redeclaration test case

func TestSrc13(t *testing.T) {
	src := `var a = 1; var a = a + 10; return a;`
	runParseTest(t, src)
}

// boolean expression with AND operator

func TestSrc14(t *testing.T) {
	src := `true && false`
	runParseTest(t, src)
}

// boolean expression with OR operator

func TestSrc15(t *testing.T) {
	src := `true || false`
	runParseTest(t, src)
}

// boolean expression with parentheses

func TestSrc16(t *testing.T) {
	src := `true && (false || true)`
	runParseTest(t, src)
}

// boolean expression in return statement

func TestSrc17(t *testing.T) {
	src := `return true && (false || true);`
	runParseTest(t, src)
}

// boolean expression with variables

func TestSrc18(t *testing.T) {
	src := `var a = true; var b = a && false; return b || true;`
	runParseTest(t, src)
}

// boolean expression with relational operator

func TestSrc19(t *testing.T) {
	src := `false || 1 < 2`
	runParseTest(t, src)
}

// relational operator in return statement

func TestSrc20(t *testing.T) {
	src := `return false || 1 < 2`
	runParseTest(t, src)
}

// relational operator with boolean combination

func TestSrc21(t *testing.T) {
	src := `return false || (1 <= 2 && true)`
	runParseTest(t, src)
}

// relational operator with variable

func TestSrc22(t *testing.T) {
	src := `var a = false; return a || (10 <= 20 && true);`
	runParseTest(t, src)
}

// multiple relational and equality operators

func TestSrc23(t *testing.T) {
	src := `(10 <= 20 && (10 != 20) && (true != false) && (true == true))`
	runParseTest(t, src)
}

// bitwise expression with boolean operators

func TestSrc24(t *testing.T) {
	src := `(3 & 7) == 3 && true || false && true`
	runParseTest(t, src)
}

// complex bitwise and boolean expression

func TestSrc25(t *testing.T) {
	src := `return ((3&7)!=3&&true||false&&true)||true`
	runParseTest(t, src)
}

// relational operator with arithmetic

func TestSrc26(t *testing.T) {
	src := `7 > 2 + 1`
	runParseTest(t, src)
}

// relational operator comparing expressions

func TestSrc27(t *testing.T) {
	src := `return 7-1>2+1`
	runParseTest(t, src)
}

// relational operator with variable expressions

func TestSrc28(t *testing.T) {
	src := `var a = 7; var b = 1; var c = 2; var d = 1; return ((a-b)>(c+d));`
	runParseTest(t, src)
}

// block statement with local variables

func TestSrc29(t *testing.T) {
	src := `{var a = 10; var b = a + 100;}`
	runParseTest(t, src)
}

// nested block statements with return

func TestSrc30(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// block statement with outer variable assignment

func TestSrc31(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// block statement modifying outer variable

func TestSrc32(t *testing.T) {
	src := `	
	var X = 1234;
	{
	X = X + 1;
	}
	X = X * 100 + 2;
	return X;`
	runParseTest(t, src)
}

// block statement with variable shadowing

func TestSrc33(t *testing.T) {
	src := `	
	var X = 1234;
	{
	var X = 6789;
	X = X + 1;
	}
	X = X + 1;
	return X;
	`
	runParseTest(t, src)
}

// if statement with expression condition

func TestSrc34(t *testing.T) {
	src := `if (1 + 1 == 2) { 
	 2 + 3;
	}`
	runParseTest(t, src)
}

// if-else statement

func TestSrc35(t *testing.T) {
	src := `if (1 + 1 == 2) { 
	 2 + 3;
	} else {
	 2 + 4;
	}`
	runParseTest(t, src)
}

// if-else-if chain

func TestSrc36(t *testing.T) {
	src := `if (1 + 1 == 21) { 
	 2 + 3;
	} else if (2 + 2 == 4) {
	 2 + 4;
	} else {	
	 2 + 6;
	}`
	runParseTest(t, src)
}

// multiple else-if conditions

func TestSrc37(t *testing.T) {
	src := `if (1 + 1 == 3) { 
	 2 + 3;
	} else if (2 + 2 == 5) {
	 2 + 4;
	} else if (3 + 3 == 7){	
	 2 + 6;
	}`
	runParseTest(t, src)
}

// if-else-if with variable assignments

func TestSrc38(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// if-else with empty blocks

func TestSrc39(t *testing.T) {
	src := `{
	var x = 1;
	if(x==9){}else{}
	if(x==10){}else{}
	if(x==11){}else{}
	if(x==12){}else{}
	}`
	runParseTest(t, src)
}

// deeply nested blocks with if-else

func TestSrc40(t *testing.T) {
	src := `{{{var x=1;{{{{if(x==9){}else{x+1;}}}}}}}}`
	runParseTest(t, src)
}

// string literals and identifiers

func TestSrc41(t *testing.T) {
	src := `"hello world" "C++" Pascal 2026 1234567890 "123"`
	runParseTest(t, src)
}

// function declaration with return

func TestSrc42(t *testing.T) {
	src := `func foo() { var a = 1; var b = 2; return a + b; }`
	runParseTest(t, src)
}

// function with parameters

func TestSrc43(t *testing.T) {
	src := `func foo(a, b, c, d) { var a = 1; var b = 2; return a * b; }`
	runParseTest(t, src)
}

// function with conditional logic

func TestSrc44(t *testing.T) {
	src := `func foo(a) { if(a==1){a=a+1;}else if(a==2){a=a+2;}else{a=a+3;} return a;}`
	runParseTest(t, src)
}

// function call with arguments

func TestSrc45(t *testing.T) {
	src := `foo(1, 2, 3, 4)`
	runParseTest(t, src)
}

// function call with complex arguments

func TestSrc46(t *testing.T) {
	src := `foo(1 + 2 * 3 - 8, true, (2==3), !!!!!true)`
	runParseTest(t, src)
}

// function call assigned to variable

func TestSrc47(t *testing.T) {
	src := `var b = 1; var a = foo(b, 2, 3); b = a + 1;`
	runParseTest(t, src)
}

// variable redeclaration

func TestSrc48(t *testing.T) {
	src := `var a = 1; var a = 2; var c = a;`
	runParseTest(t, src)
}

// function with boolean parameter

func TestSrc49(t *testing.T) {
	src := `func foo(a) { if(a) {return 1;} else {return 2;}}`
	runParseTest(t, src)
}

// recursive fibonacci function

func TestSrc50(t *testing.T) {
	src := `func fib(n) { if(n==0){return 0;} else if(n == 1) {return 1;} else {return fib(n-1) + fib(n-2);}} fib(10);`
	runParseTest(t, src)
}

// anonymous function assigned to variable

func TestSrc51(t *testing.T) {
	src := `var foo = func(a) { if(a) {return 1;} else {return 2;}};;;;foo(true);`
	runParseTest(t, src)
}

// const and var declarations

func TestSrc52(t *testing.T) {
	src := `var a = 1; const b = 2; a = 14;`
	runParseTest(t, src)
}

// while loop with single condition

func TestSrc53(t *testing.T) {
	src := `var i = 0; while(i < 5){ i = i + 1; }`
	runParseTest(t, src)
}

// while loop with multiple conditions

func TestSrc54(t *testing.T) {
	src := `var i = 0; var j = 10; while(i < 5, j > 5){ i = i + 1; j = j - 1; }`
	runParseTest(t, src)
}

// while loop with three conditions

func TestSrc55(t *testing.T) {
	src := `var a = 0; var b = 20; var c = 10; while(a < 10, b > 10, c > 5){ a = a + 1; b = b - 1; c = c - 1; }`
	runParseTest(t, src)
}

// for loop with multiple initializers and updates

func TestSrc56(t *testing.T) {
	src := `for(i = 0, j = 10; i < 5 && j > 5; i = i + 1, j = j - 1){ }`
	runParseTest(t, src)
}

// comprehensive while loop test suite

func TestSrc57(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// fibonacci with main function

func TestSrc58(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// let keyword with reassignment

func TestSrc59(t *testing.T) {
	src := `let a = 1; a = 2; a = 2.9;`
	runParseTest(t, src)
}

// const keyword with reassignment attempt

func TestSrc60(t *testing.T) {
	src := `const a = 1; a = 2; a = 2.9;`
	runParseTest(t, src)
}

// single-line comments

func TestSrc61(t *testing.T) {
	src := `// this is a comment line
	// again a new comment
	// again
	// yet again
	// `
	runParseTest(t, src)
}

// mixed single and multi-line comments

func TestSrc62(t *testing.T) {
	src := `// this is a comment line
	// again a new comment
	/* haha I a comment too */
	// now a line
	(2 + 3);
	(3 + 9);`
	runParseTest(t, src)
}

// arithmetic compound assignments

func TestSrc64(t *testing.T) {
	src := `var a = 10; a += 5; a -= 3; a *= 2; a /= 4; a %= 3`
	runParseTest(t, src)
}

// bitwise compound assignments

func TestSrc65(t *testing.T) {
	src := `var a = 12; a &= 10; a |= 3; a ^= 5`
	runParseTest(t, src)
}

// shift compound assignments

func TestSrc66(t *testing.T) {
	src := `var a = 4; a <<= 2; a >>= 1`
	runParseTest(t, src)
}

// compound assignment in for loop

func TestSrc67(t *testing.T) {
	src := `for(i = 0; i < 10; i += 2){ var x = i * 2; }`
	runParseTest(t, src)
}

// compound assignment in while loop

func TestSrc68(t *testing.T) {
	src := `var sum = 0; var i = 1; while(i <= 5){ sum += i; i += 1; }`
	runParseTest(t, src)
}

// multiple compound assignments in for loop

func TestSrc69(t *testing.T) {
	src := `for(i = 0, j = 10; i < 5; i += 1, j -= 2){ var diff = j - i; }`
	runParseTest(t, src)
}

// compound assignment with complex expressions

func TestSrc70(t *testing.T) {
	src := `var a = 100; a += 2 * 3 + 4; a -= (10 - 5); a *= 2`
	runParseTest(t, src)
}

// complex nested loop patterns

func TestSrc71(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// array with function element
// src72 := `var a = [1, 2, func(){2+3;}]; var b = a[2]; b();`
// root72 := parser.NewParser(src72).Parse()
// visitor72 := &PrintingVisitor{}
// root72.Accept(visitor72)
// fmt.Println(visitor72)

// chained function and array access

func TestSrc73(t *testing.T) {
	src := `a()[b()+1]`
	runParseTest(t, src)
}

// array with mixed types and negative index

func TestSrc74(t *testing.T) {
	src := `var arr = [1, 2, true, 2.5, func () { }];  var foo = arr[-1]; foo();`
	runParseTest(t, src)
}

// array with builtin functions (push, pop, length)

func TestSrc75(t *testing.T) {
	src := `var x = 1; 
	var y = 1 + x; 
	var arr = [x + y, x, y * 2.0, 1, 2, true, 2.5, func () {  while(true){} }];  
	var foo = arr[-1]; 
	foo(); 
	arr = push(arr, 4); 
	pop(arr); 
	length(arr);`
	runParseTest(t, src)
}

// Test 76: Simple range expression

func TestSrc76(t *testing.T) {
	src := `var r = 2...5; r`
	runParseTest(t, src)
}

// Test 77: Foreach loop with range

func TestSrc77(t *testing.T) {
	src := `var sum = 0; foreach i in 1...5 { sum = sum + i; } sum`
	runParseTest(t, src)
}

// Test 78: Foreach loop with array

func TestSrc78(t *testing.T) {
	src := `var total = 0; foreach num in [10, 20, 30, 40] { total = total + num; } total`
	runParseTest(t, src)
}

// Test 79: Nested foreach loops with ranges

func TestSrc79(t *testing.T) {
	src := `
	var result = 0;
	foreach i in 1...3 {
		foreach j in 1...3 {
			result = result + i * 10 + j;
		}
	}
	result
	`
	runParseTest(t, src)
}

// Test 80: Range builtin function and foreach with complex operations

func TestSrc80(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// Test 81: Map creation and access

func TestSrc81(t *testing.T) {
	src := `var m = map{"name": "John", "age": 25}; m["name"]`
	runParseTest(t, src)
}

// Test 82: Map with keys_map builtin

func TestSrc82(t *testing.T) {
	src := `var person = map{"name": "Alice", "age": 30, "city": "NYC"}; keys_map(person)`
	runParseTest(t, src)
}

// Test 83: Map with insert_map and remove_map

func TestSrc83(t *testing.T) {
	src := `
	var config = map{"debug": false, "port": 8080};
	insert_map(config, "host", "localhost");
	remove_map(config, "debug");
	config
	`
	runParseTest(t, src)
}

// Test 84: Map with contain_map

func TestSrc84(t *testing.T) {
	src := `
	var settings = map{"theme": "dark", "lang": "en"};
	var has_theme = contain_map(settings, "theme");
	var has_font = contain_map(settings, "font");
	has_theme
	`
	runParseTest(t, src)
}

// Test 85: Map with enumerate_map (NEW)

func TestSrc85(t *testing.T) {
	src := `
	var data = map{"x": 10, "y": 20, "z": 30};
	var pairs = enumerate_map(data);
	pairs
	`
	runParseTest(t, src)
}

// Test 86: Set creation and operations

func TestSrc86(t *testing.T) {
	src := `var s = set{1, 2, 3, 4, 5}; s`
	runParseTest(t, src)
}

// Test 87: Set with insert_set and remove_set

func TestSrc87(t *testing.T) {
	src := `
	var numbers = set{10, 20, 30};
	insert_set(numbers, 40);
	remove_set(numbers, 20);
	numbers
	`
	runParseTest(t, src)
}

// Test 88: Set with contains_set

func TestSrc88(t *testing.T) {
	src := `
	var tags = set{"golang", "programming", "tutorial"};
	var has_golang = contains_set(tags, "golang");
	var has_python = contains_set(tags, "python");
	has_golang
	`
	runParseTest(t, src)
}

// Test 89: Set with values_set (NEW)

func TestSrc89(t *testing.T) {
	src := `
	var unique_ids = set{101, 102, 103, 104};
	var id_array = values_set(unique_ids);
	id_array
	`
	runParseTest(t, src)
}

// Test 90: Complex map and set operations

func TestSrc90(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// Test 91: List creation and basic operations

func TestSrc91(t *testing.T) {
	src := `var l = list(1, 2, 3, 4, 5); l`
	runParseTest(t, src)
}

// Test 92: List with push and pop operations

func TestSrc92(t *testing.T) {
	src := `
	var nums = list(10, 20, 30);
	pushback_list(nums, 40);
	pushfront_list(nums, 5);
	var last = popback_list(nums);
	var first = popfront_list(nums);
	nums
	`
	runParseTest(t, src)
}

// Test 93: List indexing and slicing

func TestSrc93(t *testing.T) {
	src := `
	var data = list(0, 10, 20, 30, 40, 50);
	var elem = data[2];
	var slice = data[1:4];
	slice
	`
	runParseTest(t, src)
}

// Test 94: List with foreach loop

func TestSrc94(t *testing.T) {
	src := `
	var sum = 0;
	foreach num in list(1, 2, 3, 4, 5) {
		sum = sum + num;
	}
	sum
	`
	runParseTest(t, src)
}

// Test 95: Nested lists

func TestSrc95(t *testing.T) {
	src := `
	var matrix = list(list(1, 2, 3), list(4, 5, 6), list(7, 8, 9));
	var row = matrix[1];
	var elem = matrix[0][2];
	elem
	`
	runParseTest(t, src)
}

// Test 96: Tuple creation and basic operations

func TestSrc96(t *testing.T) {
	src := `var t = tuple(10, 20, 30); t`
	runParseTest(t, src)
}

// Test 97: Tuple indexing and slicing

func TestSrc97(t *testing.T) {
	src := `
	var coords = tuple(100, 200, 300, 400);
	var x = coords[0];
	var slice = coords[1:3];
	slice
	`
	runParseTest(t, src)
}

// Test 98: Tuple with foreach loop

func TestSrc98(t *testing.T) {
	src := `
	var product = 1;
	foreach val in tuple(2, 3, 4) {
		product = product * val;
	}
	product
	`
	runParseTest(t, src)
}

// Test 99: Tuple immutability and peek operations

func TestSrc99(t *testing.T) {
	src := `
	var person = tuple("Alice", 25, true);
	var name = peekfront_tuple(person);
	var active = peekback_tuple(person);
	var size = size_tuple(person);
	size
	`
	runParseTest(t, src)
}

// Test 100: Mixed list and tuple operations

func TestSrc100(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// Test 101: List with length and size functions

func TestSrc101(t *testing.T) {
	src := `
	var items = list("a", "b", "c", "d", "e");
	var len1 = length(items);
	var len2 = size_list(items);
	len1
	`
	runParseTest(t, src)
}

// Test 102: Tuple with length function

func TestSrc102(t *testing.T) {
	src := `
	var rgb = tuple(255, 128, 64);
	var len = length(rgb);
	len
	`
	runParseTest(t, src)
}

// Test 103: List index assignment

func TestSrc103(t *testing.T) {
	src := `
	var arr = list(1, 2, 3, 4, 5);
	arr[0] = 100;
	arr[2] = 300;
	arr[-1] = 999;
	arr
	`
	runParseTest(t, src)
}

// Test 104: Complex list manipulation

func TestSrc104(t *testing.T) {
	src := `
	var stack = list();
	pushback_list(stack, 10);
	pushback_list(stack, 20);
	pushback_list(stack, 30);
	var top = peekback_list(stack);
	var popped = popback_list(stack);
	var new_size = size_list(stack);
	new_size
	`
	runParseTest(t, src)
}

// Test 105: Tuple as function return value simulation

func TestSrc105(t *testing.T) {
	src := `
	var result = tuple(42, "success", true);
	var code = result[0];
	var message = result[1];
	var status = result[2];
	status
	`
	runParseTest(t, src)
}

// Test 106: List and tuple with nested structures

func TestSrc106(t *testing.T) {
	src := `
	var complex_list = list(tuple(1, 2), tuple(3, 4), tuple(5, 6));
	var first_tuple = complex_list[0];
	var second_elem = first_tuple[1];
	second_elem
	`
	runParseTest(t, src)
}

// Test 107: Empty list and tuple

func TestSrc107(t *testing.T) {
	src := `var empty_list = list(); var empty_tuple = tuple(); empty_list; empty_tuple;`
	runParseTest(t, src)
}

// Test 108: List and tuple with different data types

func TestSrc108(t *testing.T) {
	src := `
	var mixed_list = list(1, "two", 3.0, true, func() { return "hello"; });
	var mixed_tuple = tuple("start", 42, false, func() { return "world"; });
	mixed_list; mixed_tuple;
	`
	runParseTest(t, src)
}

// Test 109: struct creation

func TestSrc109(t *testing.T) {
	src := `var Point = struct Data {}`
	runParseTest(t, src)
}

// Test 110: struct with init

func TestSrc110(t *testing.T) {
	src := `var Point = struct Data { func init(){} }`
	runParseTest(t, src)
}

// Test 111: struct with other methods

func TestSrc111(t *testing.T) {
	src := `struct Data { func init(){} func move(x, y){} }`
	runParseTest(t, src)
}

// Test 112: new keyword with struct

func TestSrc112(t *testing.T) {
	src := `struct Data {}; var p = new Point(10, 20);`
	runParseTest(t, src)
}

// Test 113: struct with fields and methods

func TestSrc113(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

func TestSrc114(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// Test 115: Struct with multiple methods and field access

func TestSrc115(t *testing.T) {
	src := `struct A{ func init(){} func hello(){ return 12 } } var a=new A(); var res=a.hello(); println(res);`
	runParseTest(t, src)
}

// Test 116: Struct with nested struct

func TestSrc116(t *testing.T) {
	src := `
	struct C { func add(x, y) { return x + y } }
	var c = new C();
	var d = c.add(3, 4);
	println(d);
	`
	runParseTest(t, src)
}

// Test 117: struct with new and calling method

func TestSrc117(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

// Test 118: struct with new and field access

func TestSrc118(t *testing.T) {
	src := `
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
	runParseTest(t, src)
}

func TestSrc119(t *testing.T) {
	src := `
	struct Logic { func run(x) { if (x > 0) { while(x > 0) { x = x - 1; } } return x; } }
	`
	runParseTest(t, src)
}

func TestSrc120(t *testing.T) {
	src := `
	struct S { func init(v) { this.v = v; } func get() { return this.v; } } var s = new S(5); s.get()
	`
	runParseTest(t, src)
}

// Test 121: import statement basic

func TestSrc121(t *testing.T) {
	src := `
	import math;
	`
	runParseTest(t, src)
}

// Test 122: import with package function call

func TestSrc122(t *testing.T) {
	src := `
	import math;
	math.abs(-5);
	`
	runParseTest(t, src)
}

// Test 123: import with variable assignment from package function

func TestSrc123(t *testing.T) {
	src := `
	import math;
	var result = math.abs(-10);
	println(result);
	`
	runParseTest(t, src)
}

// Test 124: import with multiple package function calls

func TestSrc124(t *testing.T) {
	src := `
	import math;
	var x = math.pow(2, 3);
	var y = math.sqrt(16);
	var z = math.min(x, y);
	println(z);
	`
	runParseTest(t, src)
}

// Test 125: import with nested package function calls

func TestSrc125(t *testing.T) {
	src := `
	import math;
	var a = math.pow(2, 3);
	var b = math.sqrt(a);
	var c = math.abs(-b);
	println(c);
	`
	runParseTest(t, src)
}

// Test 126: import with package function calls in expressions

func TestSrc126(t *testing.T) {
	src := `
	import math;
	var result = math.pow(2, 3) + math.sqrt(16) - math.abs(-5);
	println(result);
	`
	runParseTest(t, src)
}

// Test 127: import with package function calls in loops

func TestSrc127(t *testing.T) {
	src := `
	import math;
	var sum = 0;
	for(i = 1; i <= 5; i = i + 1) {
		sum = sum + math.pow(i, 2);
	}
	println(sum);
	`
	runParseTest(t, src)
}

// Test 128: import with package function calls in conditionals

func TestSrc128(t *testing.T) {
	src := `
	import math;
	var x = -10;
	if (math.abs(x) > 5) {
		println("Large");
	} else {
		println("Small");
	}
	`
	runParseTest(t, src)
}

// Test 129: import with package function calls in functions

func TestSrc129(t *testing.T) {
	src := `
	import math;
	func calculate(a, b) {
		return math.pow(a, 2) + math.pow(b, 2);
	}
	var result = calculate(3, 4);
	println(result);
	`
	runParseTest(t, src)
}

// Test 130: import with package function calls in struct methods

func TestSrc130(t *testing.T) {
	src := `
	import math;
	struct Circle {
		func init(radius) {
			this.radius = radius;
		}
		func area() {
			return math.PI * math.pow(this.radius, 2);
		}
	}
	var c = new Circle(5);
	println(c.area());
	`
	runParseTest(t, src)
}

// Test131: enums

func TestSrc131(t *testing.T) {
	src := `enum Color { RED, GREEN, BLUE }
			var c = Color.RED;
			c;
			`
	runParseTest(t, src)
}

// Test132: enums

func TestSrc132(t *testing.T) {
	src := `enum Direction { NORTH, EAST = 1, SOUTH = 100, WEST }
			var d = Direction.SOUTH;
			var e = Direction.EAST;
			var x = d + e;
			x;
			`
	runParseTest(t, src)
}

// Test133: switch

func TestSrc133(t *testing.T) {
	src := `
	var x = 2;
	switch(x) {
		case 1:
			println("One");
			break;
		case 2:
			println("Two");
			break;
		case 3:
			println("Three");
			break;
		default:
			println("Other");
	}
	`
	runParseTest(t, src)
}

// Test134: switch with multiple cases

func TestSrc134(t *testing.T) {
	src := `
	var x = 3;
	switch(x) {
		case 1:
		case 2:
			println("One or Two");
			break;
		}`
	runParseTest(t, src)
}
