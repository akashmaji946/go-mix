/*
File: go-mix/lexer/lexer_utils.go
Author: Akash Maji
Contact: akashmaji(@iisc.ac.in)
*/
package lexer

import (
	"fmt"
	"strings"
	"unicode"
)

// isDigitASCII reports whether c is an ASCII decimal digit ('0'..'9').
// This is used in the hot path for number scanning.
func isDigitASCII(c byte) bool {
	return c >= '0' && c <= '9'
}

// isHexDigitASCII reports whether c is an ASCII hexadecimal digit.
// This is used when scanning 0x/0X integer literals.
func isHexDigitASCII(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'f') || (c >= 'A' && c <= 'F')
}

// isWhitespace checks if the given byte is a whitespace character.
// Uses Unicode's definition of whitespace, which includes:
//   - Space, tab, newline, carriage return, form feed, vertical tab
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is whitespace, false otherwise
func isWhitespace(curr byte) bool {
	return unicode.IsSpace(rune(curr))
}

// isAlphanumeric checks if the given byte is an alphanumeric character.
// This includes both letters (a-z, A-Z) and digits (0-9).
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is a letter or digit, false otherwise
func isAlphanumeric(curr byte) bool {
	return unicode.IsLetter(rune(curr)) || unicode.IsDigit(rune(curr))
}

// isNumeric checks if the given byte is a numeric digit (0-9).
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is a digit, false otherwise
func isNumeric(curr byte) bool {
	return unicode.IsDigit(rune(curr))
}

// isAlpha checks if the given byte is an alphabetic character (a-z, A-Z).
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is a letter, false otherwise
func isAlpha(curr byte) bool {
	return unicode.IsLetter(rune(curr))
}

// isEscape checks if the given byte is an escape character.
// This function is kept for backward compatibility but is no longer used
// in readStringLiteral as escape sequences are now handled properly
// through the escapeChar function.
//
// Parameters:
//   - curr: The byte to check
//
// Returns:
//   - bool: true if curr is an escape character, false otherwise
func isEscape(curr byte) bool {
	return curr == '\\' || curr == '\'' || curr == 0 ||
		curr == '\n' || curr == '\t' || curr == '\r' || curr == '\f' || curr == '\v'
}

// isSpecial checks if a character is a special symbol that is not valid in GoMix.
// This includes characters that are not part of the defined token set and are not alphanumeric or whitespace.
func isSpecial(c byte) bool {
	return !isAlphanumeric(c) && !isWhitespace(c) && !strings.ContainsRune("=+-*/%&|^~!<>.,;:(){}[]\"", rune(c))
}

// readStringLiteral reads and tokenizes a string literal from the source.
// It handles escape sequences like \n, \t, \\, \", etc.
// String literals must be enclosed in double quotes (").
//
// Supported escape sequences:
//   - \n: newline
//   - \t: tab
//   - \r: carriage return
//   - \f: form feed
//   - \v: vertical tab
//   - \\: backslash
//   - \": double quote
//   - \': single quote
//   - \0: null character
//
// Parameters:
//   - lex: Pointer to the lexer instance
//
// Returns:
//   - Token: A STRING_LIT token with the string content, or INVALID_TYPE on error
//
// Example:
//
//	Source: "hello\nworld"
//	Returns: Token{Type: STRING_LIT, Literal: "hello\nworld"}
func readStringLiteral(lex *Lexer) Token {
	if lex.Current != '"' {
		// Error: expected opening quote
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed string literal", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}
	lex.Advance() // Consume opening quote

	var builder strings.Builder

	// Read characters until closing quote
	for lex.Current != '"' {
		// Check for unterminated string (EOF or actual newline in source)
		if lex.Current == 0 {
			fmt.Errorf("[%d:%d] LEXER ERROR: string literal not terminated - unexpected EOF", lex.Line, lex.Column)
			return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
		}

		// Handle escape sequences
		if lex.Current == '\\' {
			lex.Advance() // Consume the backslash
			escapedChar, valid := escapeChar(lex.Current)
			if !valid {
				fmt.Errorf("[%d:%d] LEXER ERROR: invalid escape sequence: \\%c", lex.Line, lex.Column, lex.Current)
				return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
			}
			builder.WriteByte(escapedChar)
			lex.Advance()
			continue
		}

		// Regular character - add to string
		builder.WriteByte(lex.Current)
		lex.Advance()
	}

	lex.Advance() // Consume closing quote
	return NewTokenWithMetadata(STRING_LIT, builder.String(), lex.Line, lex.Column)
}

// escapeChar converts an escape sequence character to its actual byte value.
// This is used when processing escape sequences in string literals.
//
// Parameters:
//   - c: The character following the backslash in an escape sequence
//
// Returns:
//   - byte: The actual byte value of the escape sequence
//   - bool: true if the escape sequence is valid, false otherwise
//
// Example:
//
//	escapeChar('n') -> ('\n', true)
//	escapeChar('x') -> (0, false)
func escapeChar(c byte) (byte, bool) {
	switch c {
	case 'n':
		return '\n', true // Newline
	case 't':
		return '\t', true // Tab
	case 'r':
		return '\r', true // Carriage return
	case 'f':
		return '\f', true // Form feed
	case 'v':
		return '\v', true // Vertical tab
	case '\\':
		return '\\', true // Backslash
	case '"':
		return '"', true // Double quote
	case '\'':
		return '\'', true // Single quote
	case '0':
		return 0, true // Null character
	default:
		return 0, false // Invalid escape sequence
	}
}

// readNumber reads and tokenizes a numeric literal from the source.
// It supports both integer and floating-point numbers.
//
// Currently supported formats:
//   - Integers: 0, 10, 123, etc.
//   - Floats: 10.5, 0.123, 123.456, etc.
//   - Scientific notation: 1e9, 1.4e9, 12E-2
//   - Hexadecimal integers: 0x16
//   - Octal integers: 0777
//
// Parameters:
//   - lex: Pointer to the lexer instance
//
// Returns:
//   - Token: An INT_LIT or FLOAT_LIT token, or INVALID_TYPE on error
//
// Example:
//
//	Source: "123.45"
//	Returns: Token{Type: FLOAT_LIT, Literal: "123.45"}
func readNumber(lex *Lexer) Token {
	start := lex.Position

	// Ensure we start with a digit
	if lex.Current < '0' || lex.Current > '9' {
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed number literal", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}

	src := lex.Src
	n := lex.SrcLength

	// Fast-path: hexadecimal integer literal (0x...)
	if lex.Current == '0' && start+2 < n {
		prefix := src[start+1]
		if (prefix == 'x' || prefix == 'X') && isHexDigitASCII(src[start+2]) {
			i := start + 3
			for i < n && isHexDigitASCII(src[i]) {
				i++
			}
			lex.Column += i - start
			lex.Position = i
			if i >= n {
				lex.Current = 0
				lex.Position = n
			} else {
				lex.Current = src[i]
			}
			return NewTokenWithMetadata(INT_LIT, src[start:i], lex.Line, lex.Column)
		}
	}

	i := start + 1 // already know src[start] is a digit
	hasDot := false
	hasExp := false

	for i < n {
		c := src[i]
		if isDigitASCII(c) {
			i++
			continue
		}

		if c == '.' {
			// Stop before range operator (...)
			if i+1 < n && src[i+1] == '.' {
				break
			}
			if hasDot || hasExp {
				break
			}
			hasDot = true
			i++
			continue
		}

		if c == 'e' || c == 'E' {
			if hasExp {
				break
			}
			j := i + 1
			if j < n && (src[j] == '+' || src[j] == '-') {
				j++
			}
			if j < n && isDigitASCII(src[j]) {
				hasExp = true
				i = j + 1
				for i < n && isDigitASCII(src[i]) {
					i++
				}
				continue
			}
			break
		}

		break
	}

	lex.Column += i - start
	lex.Position = i
	if i >= n {
		lex.Current = 0
		lex.Position = n
	} else {
		lex.Current = src[i]
	}

	tokenType := INT_LIT
	if hasDot || hasExp {
		tokenType = FLOAT_LIT
	}
	return NewTokenWithMetadata(tokenType, src[start:i], lex.Line, lex.Column)
}

// readIdentifier reads and tokenizes an identifier or keyword from the source.
// Identifiers can be variable names, function names, or language keywords.
//
// Rules:
//   - Must start with a letter (a-z, A-Z) or underscore (_)
//   - Can contain letters, digits, or underscores
//   - Keywords are identified using the lookupIdent function
//
// Parameters:
//   - lex: Pointer to the lexer instance
//
// Returns:
//   - Token: An IDENTIFIER_ID token or a keyword token type, or INVALID_TYPE on error
//
// Example:
//
//	Source: "myVariable"
//	Returns: Token{Type: IDENTIFIER_ID, Literal: "myVariable"}
//
//	Source: "if"
//	Returns: Token{Type: IF_KEY, Literal: "if"}
func readIdentifier(lex *Lexer) Token {
	position := lex.Position

	// Ensure we start with a letter or underscore
	if isAlpha(lex.Current) || lex.Current == '_' {
		lex.Advance()
	} else {
		fmt.Errorf("[%d:%d] LEXER ERROR: malformed identifier", lex.Line, lex.Column)
		return NewTokenWithMetadata(INVALID_TYPE, "", lex.Line, lex.Column)
	}

	// Continue reading alphanumeric characters and underscores
	for isAlphanumeric(lex.Current) || lex.Current == '_' {
		lex.Advance()
	}

	literal := lex.Src[position:lex.Position]

	// Check if this identifier is actually a keyword
	return NewTokenWithMetadata(lookupIdent(literal), literal, lex.Line, lex.Column)
}
