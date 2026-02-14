/*
File    : go-mix/std/time.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - time.go
// This file defines the date and time builtin functions for the Go-Mix language.
// It provides functions for retrieving current time, formatting, parsing,
// and handling timezones.
package std

import (
	"io"
	"time"
)

var timeMethods = []*Builtin{
	{Name: "now", Callback: now},                // Returns current Unix timestamp (seconds)
	{Name: "now_ms", Callback: nowMs},           // Returns current Unix timestamp (milliseconds)
	{Name: "utc_now", Callback: utcNow},         // Returns current UTC Unix timestamp
	{Name: "format_time", Callback: formatTime}, // Formats a Unix timestamp
	{Name: "parse_time", Callback: parseTime},   // Parses a time string to Unix timestamp
	{Name: "timezone", Callback: timezone},      // Returns current timezone name
}

// init registers the time methods as global builtins and as a package for import.
func init() {
	// Register as global builtins (for backward compatibility)
	Builtins = append(Builtins, timeMethods...)

	// Register as a package (for import functionality)
	timePackage := &Package{
		Name:      "time",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range timeMethods {
		timePackage.Functions[method.Name] = method
	}
	RegisterPackage(timePackage)
}

// now returns the current local Unix timestamp in seconds.
//
// Syntax: now()
//
// Example:
//
//	var t = now();
//	println("Current timestamp: " + t);
func now(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: now expects 0 arguments")
	}
	return &Integer{Value: time.Now().Unix()}
}

// nowMs returns the current local Unix timestamp in milliseconds.
//
// Syntax: now_ms()
func nowMs(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: now_ms expects 0 arguments")
	}
	return &Integer{Value: time.Now().UnixMilli()}
}

// utcNow returns the current UTC Unix timestamp in seconds.
//
// Syntax: utc_now()
func utcNow(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: utc_now expects 0 arguments")
	}
	return &Integer{Value: time.Now().UTC().Unix()}
}

// formatTime converts a Unix timestamp to a formatted string.
// It uses Go's reference time layout: Mon Jan 2 15:04:05 MST 2006.
//
// Syntax: format_time(timestamp, layout)
//
// Example:
//
//	var s = format_time(now(), "2006-01-02 15:04:05");
//	println("Formatted: " + s);
func formatTime(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: format_time expects 2 arguments (timestamp, layout)")
	}
	if args[0].GetType() != IntegerType {
		return createError("ERROR: first argument to `format_time` must be an integer (timestamp)")
	}
	if args[1].GetType() != StringType {
		return createError("ERROR: second argument to `format_time` must be a string (layout)")
	}

	ts := args[0].(*Integer).Value
	layout := args[1].ToString()
	t := time.Unix(ts, 0)
	return &String{Value: t.Format(layout)}
}

// parseTime converts a formatted string to a Unix timestamp.
//
// Syntax: parse_time(value, layout)
//
// Example:
//
//	var ts = parse_time("2023-10-27", "2006-01-02");
//	println("Parsed timestamp: " + ts);
func parseTime(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: parse_time expects 2 arguments (value, layout)")
	}
	if args[0].GetType() != StringType {
		return createError("ERROR: first argument to `parse_time` must be a string (value)")
	}
	if args[1].GetType() != StringType {
		return createError("ERROR: second argument to `parse_time` must be a string (layout)")
	}

	val := args[0].ToString()
	layout := args[1].ToString()

	// Use local time for parsing to match 'now()' behavior
	t, err := time.ParseInLocation(layout, val, time.Local)
	if err != nil {
		return createError("ERROR: failed to parse time: %v", err)
	}
	return &Integer{Value: t.Unix()}
}

// timezone returns the name of the current local timezone.
//
// Syntax: timezone()
//
// Example:
//
//	var tz = timezone();
//	println("Current zone: " + tz);
func timezone(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: timezone expects 0 arguments")
	}
	name, _ := time.Now().Zone()
	return &String{Value: name}
}
