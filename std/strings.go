/*
File    : go-mix/std/strings.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - strings.go
// This file defines the string and character builtin functions for the Go-Mix language.
// It provides functions for case conversion, searching, splitting, joining,
// and character code manipulation.
package std

import (
	"io"
	"strings"
	"unicode"
)

var stringMethods = []*Builtin{

	{Name: "upper", Callback: upperString}, // Converts string to uppercase
	{Name: "lower", Callback: lowerString}, // Converts string to lowercase
	{Name: "trim", Callback: trimString},   // Trims whitespace from both ends
	{Name: "ltrim", Callback: ltrimString}, // Trims whitespace from the left
	{Name: "rtrim", Callback: rtrimString}, // Trims whitespace from the right
	{Name: "split", Callback: splitString}, // Splits string into an array by separator
	{Name: "join", Callback: joinString},   // Joins an array into a string with separator

	{Name: "contains_string", Callback: containsString}, // Checks if string contains a substring
	{Name: "reverse_string", Callback: reverseString},   // Reverses a string
	{Name: "replace_string", Callback: replaceString},   // Replaces occurrences of a substring
	{Name: "index_string", Callback: indexString},       // Returns index of first occurrence of substring

	{Name: "ord", Callback: ordString},                // Returns integer value of a character
	{Name: "chr", Callback: chrString},                // Returns character from integer value
	{Name: "starts_with", Callback: startsWithString}, // Checks if string starts with prefix
	{Name: "ends_with", Callback: endsWithString},     // Checks if string ends with suffix
	{Name: "strcmp", Callback: strcmpString},          // Compares two strings (-1, 0, 1)

	{Name: "substring", Callback: substringString},   // Extracts a part of a string
	{Name: "capitalize", Callback: capitalizeString}, // Capitalizes the first letter
	{Name: "count", Callback: countString},           // Counts occurrences of a substring
	{Name: "is_digit", Callback: isDigitFuncString},  // Checks if string contains only digits
	{Name: "is_alpha", Callback: isAlphaFuncString},  // Checks if string contains only letters

	{Name: "size_string", Callback: stringLengthString},   // Returns the length of a string
	{Name: "length_string", Callback: stringLengthString}, // Returns the length of a string (alias)
}

// init registers the string methods as global builtins and as a package for import.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, stringMethods...)

	// Register as a package (for import functionality)
	stringsPackage := &Package{
		Name:      "strings",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range stringMethods {
		stringsPackage.Functions[method.Name] = method
	}
	RegisterPackage(stringsPackage)
}

// stringLengthString returns the length of a string.
//
// Syntax: length(str)
//
// Usage:
//
//	Returns the number of characters in the string.
//
// Example:
//
//	length("hello"); // Returns 5
func stringLengthString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: length expects 1 argument, got %d", len(args))
	}
	return &Integer{Value: int64(len([]rune(args[0].ToString())))}
}

// upperString converts a string to uppercase.
//
// Syntax: upperString(str)
//
// Usage:
//
//	Returns a copy of the string with all Unicode letters mapped to their upperString case.
//
// Example:
//
//	upperString("hello"); // Returns "HELLO"
func upperString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: upper expects 1 argument, got %d", len(args))
	}
	return &String{Value: strings.ToUpper(args[0].ToString())}
}

// lowerString converts a string to lowercase.
//
// Syntax: lowerString(str)
//
// Usage:
//
//	Returns a copy of the string with all Unicode letters mapped to their lowerString case.
//
// Example:
//
//	lowerString("HELLO"); // Returns "hello"
func lowerString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: lower expects 1 argument, got %d", len(args))
	}
	return &String{Value: strings.ToLower(args[0].ToString())}
}

// trimString removes whitespace from both ends of a string.
//
// Syntax: trimString(str)
//
// Usage:
//
//	Returns a slice of the string with all leading and trailing white space removed, as defined by Unicode.
//
// Example:
//
//	trimString("  hello  "); // Returns "hello"
func trimString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: trim expects 1 argument, got %d", len(args))
	}
	return &String{Value: strings.TrimSpace(args[0].ToString())}
}

// ltrimString removes whitespace from the left side of a string.
//
// Syntax: ltrimString(str)
//
// Usage:
//
//	Returns a copy of the string with all leading white space removed.
//
// Example:
//
//	ltrimString("  hello  "); // Returns "hello  "
func ltrimString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: ltrim expects 1 argument, got %d", len(args))
	}
	return &String{Value: strings.TrimLeftFunc(args[0].ToString(), unicode.IsSpace)}
}

// rtrimString removes whitespace from the right side of a string.
//
// Syntax: rtrimString(str)
//
// Usage:
//
//	Returns a copy of the string with all trailing white space removed.
//
// Example:
//
//	rtrimString("  hello  "); // Returns "  hello"
func rtrimString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: rtrim expects 1 argument, got %d", len(args))
	}
	return &String{Value: strings.TrimRightFunc(args[0].ToString(), unicode.IsSpace)}
}

// splitString divides a string into an array of substrings based on a separator.
//
// Syntax: splitString(str, separator)
//
// Usage:
//
//	Slices the string into all substrings separated by the separator and returns an array of the substrings.
//
// Example:
//
//	splitString("a,b,c", ","); // Returns ["a", "b", "c"]
func splitString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: split expects 2 arguments (str, sep), got %d", len(args))
	}
	s := args[0].ToString()
	sep := args[1].ToString()
	parts := strings.Split(s, sep)
	elements := make([]GoMixObject, len(parts))
	for i, p := range parts {
		elements[i] = &String{Value: p}
	}
	return &Array{Elements: elements}
}

// joinString concatenates elements of an array into a single string using a separator.
//
// Syntax: joinString(array, separator)
//
// Usage:
//
//	Concatenates the elements of its first argument to create a single string. The separator
//	string is placed between elements in the resulting string.
//
// Example:
//
//	joinString(["Go", "Mix"], "-"); // Returns "Go-Mix"
func joinString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: join expects 2 arguments (array, sep), got %d", len(args))
	}
	if args[0].GetType() != ArrayType {
		return createError("ERROR: first argument to `join` must be an array, got %s", args[0].GetType())
	}
	arr := args[0].(*Array)
	sep := args[1].ToString()
	parts := make([]string, len(arr.Elements))
	for i, el := range arr.Elements {
		parts[i] = el.ToString()
	}
	return &String{Value: strings.Join(parts, sep)}
}

// replaceString returns a copy of the string with all occurrences of 'old' replaced by 'new'.
//
// Syntax: replaceString(str, old, new)
//
// Usage:
//
//	Returns a copy of the string with all non-overlapping instances of 'old' replaced by 'new'.
//
// Example:
//
//	replaceString("banana", "a", "o"); // Returns "bonono"
func replaceString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: replace expects 3 arguments (str, old, new), got %d", len(args))
	}
	s := args[0].ToString()
	old := args[1].ToString()
	newSub := args[2].ToString()
	return &String{Value: strings.ReplaceAll(s, old, newSub)}
}

// containsString reports whether substring is within str.
//
// Syntax: containsString(str, substring)
//
// Usage:
//
//	Returns true if the substring is within the string, false otherwise.
//
// Example:
//
//	containsString("hello", "ell"); // Returns true
func containsString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: contains expects 2 arguments, got %d", len(args))
	}
	return &Boolean{Value: strings.Contains(args[0].ToString(), args[1].ToString())}
}

// indexString returns the indexString of the first instance of substring in str, or -1 if not present.
//
// Syntax: indexString(str, substring)
//
// Usage:
//
//	Returns the indexString of the first instance of the substring in the string, or -1 if the substring is not present.
//
// Example:
//
//	indexString("hello", "e"); // Returns 1
func indexString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: index expects 2 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(strings.Index(args[0].ToString(), args[1].ToString()))}
}

// ordString returns the integer Unicode code point of a character.
// If a string is provided, it returns the code point of the first character.
//
// Syntax: ordString(char_or_string)
//
// Usage:
//
//	Returns the integer value representing the Unicode code point of the character.
//
// Example:
//
//	ordString('A');   // Returns 65
//	ordString("ABC"); // Returns 65
func ordString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: ord expects 1 argument, got %d", len(args))
	}
	if args[0].GetType() == CharType {
		return &Integer{Value: int64(args[0].(*Char).Value)}
	}
	s := args[0].ToString()
	if len(s) == 0 {
		return createError("ERROR: ord expects non-empty string")
	}
	return &Integer{Value: int64([]rune(s)[0])}
}

// chrString returns a character object representing the given Unicode code point.
//
// Syntax: chrString(integer)
//
// Usage:
//
//	Returns a character object corresponding to the provided integer Unicode code point.
//
// Example:
//
//	chrString(65); // Returns 'A'
func chrString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: chr expects 1 argument, got %d", len(args))
	}
	if args[0].GetType() != IntegerType {
		return createError("ERROR: chr expects integer, got %s", args[0].GetType())
	}
	return &Char{Value: rune(args[0].(*Integer).Value)}
}

// startsWithString tests whether the string starts with prefix.
//
// Syntax: starts_with(str, prefix)
//
// Usage:
//
//	Returns true if the string starts with the specified prefix, false otherwise.
//
// Example:
//
//	starts_with("hello", "he"); // Returns true
func startsWithString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: starts_with expects 2 arguments, got %d", len(args))
	}
	return &Boolean{Value: strings.HasPrefix(args[0].ToString(), args[1].ToString())}
}

// endsWithString tests whether the string ends with suffix.
//
// Syntax: ends_with(str, suffix)
//
// Usage:
//
//	Returns true if the string ends with the specified suffix, false otherwise.
//
// Example:
//
//	ends_with("hello", "lo"); // Returns true
func endsWithString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: ends_with expects 2 arguments, got %d", len(args))
	}
	return &Boolean{Value: strings.HasSuffix(args[0].ToString(), args[1].ToString())}
}

// strcmpString compares two strings lexicographically.
// Returns 0 if s1 == s2, -1 if s1 < s2, and 1 if s1 > s2.
//
// Syntax: strcmpString(s1, s2)
//
// Usage:
//
//	Performs a lexicographical comparison of two strings.
//
// Example:
//
//	strcmpString("apple", "banana"); // Returns -1
//	strcmpString("hello", "hello");  // Returns 0
//	strcmpString("zoo", "zebra");    // Returns 1
func strcmpString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: strcmp expects 2 arguments, got %d", len(args))
	}
	s1 := args[0].ToString()
	s2 := args[1].ToString()
	if s1 < s2 {
		return &Integer{Value: -1}
	} else if s1 > s2 {
		return &Integer{Value: 1}
	}
	return &Integer{Value: 0}
}

// reverseString returns a new string with the characters in reverseString order.
//
// Syntax: reverseString(str)
//
// Usage:
//
//	Returns a new string with the characters of the input string in reverseString order.
//	Handles Unicode characters correctly.
//
// Example:
//
//	reverseString("abc"); // Returns "cba"
func reverseString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: reverse expects 1 argument, got %d", len(args))
	}
	s := args[0].ToString()
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return &String{Value: string(runes)}
}

// substringString extracts a part of a string.
//
// Syntax: substringString(str, start, [length])
//
// Usage:
//
//	Returns a substringString starting at the 'start' index. If 'length' is provided,
//	it returns that many characters. Otherwise, it returns until the end of the string.
//
// Example:
//
//	substringString("hello", 1, 2); // Returns "el"
//	substringString("hello", 2);    // Returns "llo"
func substringString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 2 || len(args) > 3 {
		return createError("ERROR: substring expects 2 or 3 arguments, got %d", len(args))
	}
	s := args[0].ToString()
	runes := []rune(s)
	strLen := int64(len(runes))

	if args[1].GetType() != IntegerType {
		return createError("ERROR: substring start index must be an integer")
	}
	start := args[1].(*Integer).Value

	if start < 0 || start > strLen {
		return createError("ERROR: substring start index out of bounds")
	}

	length := strLen - start
	if len(args) == 3 {
		if args[2].GetType() != IntegerType {
			return createError("ERROR: substring length must be an integer")
		}
		length = args[2].(*Integer).Value
	}

	if length < 0 || start+length > strLen {
		return createError("ERROR: substring length out of bounds")
	}

	return &String{Value: string(runes[start : start+length])}
}

// capitalizeString converts the first character of a string to uppercase and the rest to lowercase.
//
// Syntax: capitalizeString(str)
//
// Example:
//
//	capitalizeString("gOMIX"); // Returns "Gomix"
func capitalizeString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: capitalize expects 1 argument, got %d", len(args))
	}
	s := args[0].ToString()
	if len(s) == 0 {
		return &String{Value: ""}
	}
	runes := []rune(s)
	return &String{Value: strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))}
}

// countString returns the number of non-overlapping instances of a substring in a string.
//
// Syntax: countString(str, substring)
//
// Example:
//
//	countString("banana", "a"); // Returns 3
func countString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: count expects 2 arguments, got %d", len(args))
	}
	return &Integer{Value: int64(strings.Count(args[0].ToString(), args[1].ToString()))}
}

// isDigitFuncString checks if the string consists entirely of decimal digits.
//
// Syntax: is_digit(str)
//
// Example:
//
//	is_digit("123"); // Returns true
//	is_digit("12a"); // Returns false
func isDigitFuncString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: is_digit expects 1 argument, got %d", len(args))
	}
	s := args[0].ToString()
	if len(s) == 0 {
		return &Boolean{Value: false}
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return &Boolean{Value: false}
		}
	}
	return &Boolean{Value: true}
}

// isAlphaFuncString checks if the string consists entirely of alphabetic characters.
//
// Syntax: is_alpha(str)
//
// Example:
//
//	is_alpha("abc"); // Returns true
//	is_alpha("a12"); // Returns false
func isAlphaFuncString(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: is_alpha expects 1 argument, got %d", len(args))
	}
	s := args[0].ToString()
	if len(s) == 0 {
		return &Boolean{Value: false}
	}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return &Boolean{Value: false}
		}
	}
	return &Boolean{Value: true}
}
