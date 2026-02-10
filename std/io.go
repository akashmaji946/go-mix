/*
File    : go-mix/std/io.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - io.go
// This file defines the I/O builtin functions for the GoMix language.
// It provides functions for reading input from the standard input stream,
// including line-based reading and formatted scanning.
package std

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// stdinReader is a persistent buffered reader to ensure sequential I/O calls
// do not lose data between buffer swaps.
var stdinReader = bufio.NewReader(os.Stdin)

var ioMethods = []*Builtin{
	{Name: "scanln", Callback: scanln},     // Reads a line of text from stdin
	{Name: "scanf", Callback: scanf},       // Reads formatted input from stdin
	{Name: "input", Callback: input},       // Prints a prompt and reads a line
	{Name: "scan", Callback: scan},         // Reads until a specific string delimiter
	{Name: "getchar", Callback: getchar},   // Reads a single character
	{Name: "putchar", Callback: putchar},   // Prints a single character
	{Name: "gets", Callback: scanln},       // Alias for scanln (C-style)
	{Name: "puts", Callback: println},      // Alias for println (C-style)
	{Name: "sprintf", Callback: sprintf},   // Returns a formatted string
	{Name: "flush", Callback: flush},       // Flushes the input buffer and output writer
	{Name: "eprintln", Callback: eprintln}, // Prints to standard error with newline
	{Name: "eprintf", Callback: eprintf},   // Prints formatted string to standard error
	// {Name: "exit", Callback: exitFunc},     // Terminates the program
}

// init registers the I/O methods as global builtins.
func init() {
	Builtins = append(Builtins, ioMethods...)
}

// scanln reads a single line of text from standard input.
//
// Syntax: scanln()
//
// Usage:
//
//	Reads characters from stdin until a newline character is encountered.
//	Returns the line as a string (excluding the newline).
//
// Example:
//
//	var name = scanln();
//	println("Hello, " + name);
func scanln(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: scanln expects 0 arguments, got %d", len(args))
	}

	text, err := stdinReader.ReadString('\n')
	if err != nil && err != io.EOF {
		return createError("ERROR: failed to read from stdin: %v", err)
	}

	// Trim the newline character
	return &String{Value: strings.TrimRight(text, "\r\n")}
}

// scanf reads formatted input from standard input.
//
// Syntax: scanf(format_string)
//
// Usage:
//
//	Reads from stdin according to the provided format string.
//	Since GoMix does not support pointers for arguments, this function
//	returns an Array containing the values successfully scanned.
//
// Example:
//
//	var data = scanf("%d %s");
//	var age = data[0];
//	var name = data[1];
func scanf(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: scanf expects 1 argument (format string), got %d", len(args))
	}

	if args[0].GetType() != StringType {
		return createError("ERROR: argument to `scanf` must be a string, got '%s'", args[0].GetType())
	}

	format := args[0].ToString()

	// We need to parse the format string to provide pointers to concrete types to fmt.Fscanf.
	// fmt.Scanf cannot scan into *interface{}.
	var ptrs []interface{}
	var types []string

	for i := 0; i < len(format); i++ {
		if format[i] == '%' && i+1 < len(format) {
			i++
			if format[i] == '%' {
				continue
			}
			// Skip flags, width, precision
			for i < len(format) && strings.ContainsRune("0123456789.+- ", rune(format[i])) {
				i++
			}
			if i >= len(format) {
				break
			}
			verb := format[i]
			switch verb {
			case 'd', 'o', 'x', 'X', 'b', 'c':
				var v int64
				ptrs = append(ptrs, &v)
				types = append(types, "int")
			case 'f', 'e', 'E', 'g', 'G':
				var v float64
				ptrs = append(ptrs, &v)
				types = append(types, "float")
			case 's', 'q':
				var v string
				ptrs = append(ptrs, &v)
				types = append(types, "string")
			case 't':
				var v bool
				ptrs = append(ptrs, &v)
				types = append(types, "bool")
			default:
				var v string
				ptrs = append(ptrs, &v)
				types = append(types, "string")
			}
		}
	}

	if len(ptrs) == 0 {
		return &Array{Elements: []GoMixObject{}}
	}

	n, err := fmt.Fscanf(stdinReader, format, ptrs...)
	if err != nil && err != io.EOF {
		return createError("ERROR: scanf failed: %v", err)
	}

	// Convert scanned values to GoMix objects
	elements := make([]GoMixObject, n)
	for i := 0; i < n; i++ {
		switch types[i] {
		case "int":
			elements[i] = &Integer{Value: *ptrs[i].(*int64)}
		case "float":
			elements[i] = &Float{Value: *ptrs[i].(*float64)}
		case "string":
			elements[i] = &String{Value: *ptrs[i].(*string)}
		case "bool":
			elements[i] = &Boolean{Value: *ptrs[i].(*bool)}
		}
	}

	return &Array{Elements: elements}
}

// input prints a prompt to the writer and reads a line of text from stdin.
//
// Syntax: input([prompt_string])
//
// Example:
//
//	var age = input("Enter your age: ");
func input(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) > 0 {
		print(rt, writer, args...)
		if flusher, ok := writer.(interface{ Sync() error }); ok {
			flusher.Sync()
		}
	}
	return scanln(rt, writer)
}

// scan reads from stdin until the specified string delimiter is encountered.
// The delimiter sequence is NOT consumed and remains in the buffer.
//
// Syntax: scan(delimiter_string)
//
// Example:
//
//	var part = scan(";;");
//	var first_semi = getchar(); // Consumes the first ";"
func scan(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: scan expects 1 argument (delimiter), got %d", len(args))
	}

	delim := args[0].ToString()
	if len(delim) == 0 {
		return createError("ERROR: scan delimiter cannot be empty")
	}

	var builder strings.Builder
	for {
		// Peek to see if the next sequence matches the delimiter
		peeked, err := stdinReader.Peek(len(delim))
		if err == nil && string(peeked) == delim {
			break
		}

		// If not, read one byte and continue
		b, err := stdinReader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return createError("ERROR: failed to read: %v", err)
		}
		builder.WriteByte(b)
	}

	return &String{Value: builder.String()}
}

// getchar reads a single character from standard input.
//
// Syntax: getchar()
func getchar(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: getchar expects 0 arguments, got %d", len(args))
	}

	b, err := stdinReader.ReadByte()
	if err != nil {
		if err == io.EOF {
			return &Nil{}
		}
		return createError("ERROR: getchar failed: %v", err)
	}

	return &String{Value: string(b)}
}

// putchar outputs a single character to the writer.
//
// Syntax: putchar(char_or_int)
func putchar(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: putchar expects 1 argument, got %d", len(args))
	}

	var charStr string
	arg := args[0]

	if arg.GetType() == IntegerType {
		charStr = string(rune(arg.(*Integer).Value))
	} else {
		s := arg.ToString()
		if len(s) > 0 {
			charStr = string(s[0])
		}
	}

	fmt.Fprint(writer, charStr)

	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}

	return &Nil{}
}

// eprintln outputs string representations to standard error with a trailing newline.
//
// Syntax: eprintln(arg1, arg2, ...)
func eprintln(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	res := ""
	for _, arg := range args {
		res += arg.ToString() + " "
	}
	if len(args) > 0 {
		res = res[:len(res)-1]
	}
	fmt.Fprintln(os.Stderr, res)
	return &Nil{}
}

// eprintf outputs a formatted string to standard error.
//
// Syntax: eprintf(format, arg1, ...)
func eprintf(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: eprintf expects at least 1 argument")
	}
	if args[0].GetType() != StringType {
		return createError("ERROR: first argument to `eprintf` must be a string")
	}

	format := args[0].ToString()
	arguments := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		val, err := ExtractValue(arg)
		if err != nil {
			return &Error{Message: err.Error()}
		}
		arguments[i] = val
	}
	fmt.Fprintf(os.Stderr, format, arguments...)
	return &Nil{}
}

// sprintf returns a formatted string using Go's fmt.Sprintf style formatting.
//
// Syntax: sprintf(format_string, arg1, arg2, ...)
func sprintf(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: sprintf expects at least 1 argument (format string)")
	}

	if args[0].GetType() != StringType {
		return createError("ERROR: first argument to `sprintf` must be a string, got `%s`", args[0].GetType())
	}

	format := args[0].ToString()
	arguments := make([]interface{}, len(args)-1)
	for i, arg := range args[1:] {
		val, err := ExtractValue(arg)
		if err != nil {
			return &Error{Message: err.Error()}
		}
		arguments[i] = val
	}

	return &String{Value: fmt.Sprintf(format, arguments...)}
}

// flush clears the input buffer and flushes the output writer.
//
// Syntax: flush()
//
// Usage:
//  1. Discards all currently buffered characters in the input stream. This ensures
//     that the next input operation (like scanln) will wait for new typing.
//  2. Flushes the output writer (if supported) to ensure all text is displayed.
func flush(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: flush expects 0 arguments, got %d", len(args))
	}

	// Clear the input buffer
	stdinReader.Discard(stdinReader.Buffered())

	// Flush the output writer
	if flusher, ok := writer.(interface{ Flush() error }); ok {
		flusher.Flush()
	} else if syncer, ok := writer.(interface{ Sync() error }); ok {
		syncer.Sync()
	}

	return &Nil{}
}
