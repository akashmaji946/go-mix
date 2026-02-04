package objects

import (
	"fmt"
	"io"
)

type CallbackFunc func(writer io.Writer, args ...GoMixObject) GoMixObject

type Builtin struct {
	Name     string
	Callback CallbackFunc
}

// global builtins
var Builtins = []*Builtin{
	{
		Name:     "print",
		Callback: print,
	},
	{
		Name:     "println",
		Callback: println,
	},
	{
		Name:     "printf",
		Callback: printf,
	},
	{
		Name:     "length",
		Callback: length,
	},
	{
		Name:     "tostring",
		Callback: tostring,
	},
}

func createError(format string, a ...interface{}) *Error {
	return &Error{Message: fmt.Sprintf(format, a...)}
}

func tostring(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1", len(args))
	}
	return &String{Value: fmt.Sprintf("\"%s\"", args[0].ToString())}
}

func print(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or more", len(args))
	}
	res := ""
	for _, arg := range args {
		res += arg.ToString() + " "
	}
	if len(args) > 0 {
		res = res[:len(res)-1]
	}
	fmt.Fprint(writer, res)
	// Flush if writer supports it
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	return &Nil{}
}

func println(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or more", len(args))
	}
	res := ""
	for _, arg := range args {
		res += arg.ToString() + " "
	}
	if len(args) > 0 {
		res = res[:len(res)-1]
	}
	fmt.Fprintln(writer, res)
	// Flush if writer supports it
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	return &Nil{}
}

func printf(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: wrong number of arguments. got=%d, want=1 or more", len(args))
	}
	if args[0].GetType() != StringType {
		return createError("ERROR: first argument to `printf` must be a string, got `%s`", args[0].GetType())
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
	fmt.Fprintf(writer, format, arguments...)
	// Flush if writer supports it
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	return &Nil{}
}

func length(writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return &Error{Message: fmt.Sprintf("wrong number of arguments. got=%d, want=1", len(args))}
	}
	switch args[0].GetType() {
	case StringType:
		return &Integer{Value: int64(len(args[0].ToString()))}
	case ArrayType:
		return &Integer{Value: int64(len(args[0].(*Array).Elements))}
	default:
		return &Error{Message: fmt.Sprintf("argument to `length` not supported, got '%s'", args[0].GetType())}
	}
}
