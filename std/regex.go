/*
File    : go-mix/std/regex.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - regex.go
// This file defines regular expression builtin functions.
package std

import (
	"io"
	"regexp"
)

var regexMethods = []*Builtin{
	{Name: "match_regex", Callback: regexMatch},     // Checks if a pattern matches a string
	{Name: "find_regex", Callback: regexFind},       // Finds the first match in a string
	{Name: "findall_regex", Callback: regexFindAll}, // Finds all matches in a string
	{Name: "replace_regex", Callback: regexReplace}, // Replaces matches in a string
	{Name: "split_regex", Callback: regexSplit},     // Splits string by pattern
}

func init() {
	Builtins = append(Builtins, regexMethods...)

	regexPackage := &Package{
		Name:      "regex",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range regexMethods {
		regexPackage.Functions[method.Name] = method
	}
	RegisterPackage(regexPackage)
}

// regexMatch checks if the string matches the pattern.
// Syntax: match_regex(pattern, str)
func regexMatch(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: match_regex expects 2 arguments (pattern, str)")
	}
	pattern := args[0].ToString()
	str := args[1].ToString()

	matched, err := regexp.MatchString(pattern, str)
	if err != nil {
		return createError("ERROR: invalid regex pattern: %v", err)
	}

	return &Boolean{Value: matched}
}

// regexFind returns the first substring matching the pattern.
// Syntax: find_regex(pattern, str)
func regexFind(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: find_regex expects 2 arguments (pattern, str)")
	}
	pattern := args[0].ToString()
	str := args[1].ToString()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return createError("ERROR: invalid regex pattern: %v", err)
	}

	return &String{Value: re.FindString(str)}
}

// regexFindAll returns all substrings matching the pattern.
// Syntax: findall_regex(pattern, str, [n])
func regexFindAll(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 2 || len(args) > 3 {
		return createError("ERROR: findall_regex expects 2 or 3 arguments (pattern, str, [n])")
	}
	pattern := args[0].ToString()
	str := args[1].ToString()
	n := -1

	if len(args) == 3 {
		if args[2].GetType() != IntegerType {
			return createError("ERROR: third argument to findall_regex must be an integer")
		}
		n = int(args[2].(*Integer).Value)
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return createError("ERROR: invalid regex pattern: %v", err)
	}

	matches := re.FindAllString(str, n)
	elements := make([]GoMixObject, len(matches))
	for i, match := range matches {
		elements[i] = &String{Value: match}
	}

	return &Array{Elements: elements}
}

// regexReplace replaces occurrences of the pattern with the replacement string.
// Syntax: replace_regex(pattern, str, repl)
func regexReplace(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: replace_regex expects 3 arguments (pattern, str, repl)")
	}
	pattern := args[0].ToString()
	str := args[1].ToString()
	repl := args[2].ToString()

	re, err := regexp.Compile(pattern)
	if err != nil {
		return createError("ERROR: invalid regex pattern: %v", err)
	}

	return &String{Value: re.ReplaceAllString(str, repl)}
}

// regexSplit splits the string into substrings separated by the pattern.
// Syntax: split_regex(pattern, str, [n])
func regexSplit(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 2 || len(args) > 3 {
		return createError("ERROR: split_regex expects 2 or 3 arguments (pattern, str, [n])")
	}
	pattern := args[0].ToString()
	str := args[1].ToString()
	n := -1

	if len(args) == 3 {
		if args[2].GetType() != IntegerType {
			return createError("ERROR: third argument to split_regex must be an integer")
		}
		n = int(args[2].(*Integer).Value)
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return createError("ERROR: invalid regex pattern: %v", err)
	}

	parts := re.Split(str, n)
	elements := make([]GoMixObject, len(parts))
	for i, part := range parts {
		elements[i] = &String{Value: part}
	}

	return &Array{Elements: elements}
}
