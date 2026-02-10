/*
File    : go-mix/file/file.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package file implements stateful file I/O for the GoMix language.
// It provides a File object that wraps a native OS file handle and
// exposes methods for reading, writing, seeking, and closing files.
package file

import (
	"fmt"
	"io"
	"os"

	"github.com/akashmaji946/go-mix/std"
)

// FileObject represents an open file handle in GoMix.
type FileObject struct {
	Handle *os.File
	Path   string
}

// GetType returns the GoMixType of the object ("file").
func (f *FileObject) GetType() std.GoMixType { return std.FileType }

// ToString returns a string representation of the file handle.
func (f *FileObject) ToString() string { return fmt.Sprintf("<file: %s>", f.Path) }

// ToObject returns a detailed representation of the file handle.
func (f *FileObject) ToObject() string { return f.ToString() }

var fileMethods = []*std.Builtin{
	{Name: "fopen", Callback: fopen},   // Opens a file and returns a handle
	{Name: "fclose", Callback: fclose}, // Closes an open file handle
	{Name: "fread", Callback: fread},   // Reads N bytes from a file handle
	{Name: "fwrite", Callback: fwrite}, // Writes a string to a file handle
	{Name: "fseek", Callback: fseek},   // Moves the file cursor
	{Name: "ftell", Callback: ftell},   // Returns the current cursor position
}

// init registers the file methods as global builtins.
func init() {
	std.Builtins = append(std.Builtins, fileMethods...)
}

// createError is a local helper to create GoMix error objects.
func createError(format string, a ...interface{}) *std.Error {
	return &std.Error{Message: fmt.Sprintf(format, a...)}
}

// fopen opens a file with the specified mode.
//
// Syntax: fopen(path, mode)
// Modes: "r" (read), "w" (write/truncate), "a" (append), "r+" (read/write)
//
// Example:
//
//	var f = fopen("test.txt", "r");
func fopen(rt std.Runtime, writer io.Writer, args ...std.GoMixObject) std.GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: fopen expects 2 arguments (path, mode)")
	}

	path := args[0].ToString()
	mode := args[1].ToString()

	var flag int
	switch mode {
	case "r":
		flag = os.O_RDONLY
	case "w":
		flag = os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	case "a":
		flag = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	case "r+":
		flag = os.O_RDWR
	case "w+":
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	default:
		return createError("ERROR: invalid file mode '%s'", mode)
	}

	handle, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return createError("ERROR: could not open file '%s': %v", path, err)
	}

	return &FileObject{Handle: handle, Path: path}
}

// fclose closes an open file handle.
//
// Syntax: fclose(file_handle)
func fclose(rt std.Runtime, writer io.Writer, args ...std.GoMixObject) std.GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: fclose expects 1 argument")
	}
	if args[0].GetType() != std.FileType {
		return createError("ERROR: argument to `fclose` must be a file handle")
	}

	f := args[0].(*FileObject)
	if err := f.Handle.Close(); err != nil {
		return createError("ERROR: failed to close file: %v", err)
	}
	return &std.Nil{}
}

// fread reads a specified number of bytes from the file handle.
//
// Syntax: fread(file_handle, num_bytes)
func fread(rt std.Runtime, writer io.Writer, args ...std.GoMixObject) std.GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: fread expects 2 arguments (handle, size)")
	}
	if args[0].GetType() != std.FileType {
		return createError("ERROR: first argument to `fread` must be a file handle")
	}
	if args[1].GetType() != std.IntegerType {
		return createError("ERROR: second argument to `fread` must be an integer")
	}

	f := args[0].(*FileObject)
	size := args[1].(*std.Integer).Value
	buf := make([]byte, size)

	n, err := f.Handle.Read(buf)
	if err != nil && err != io.EOF {
		return createError("ERROR: read failed: %v", err)
	}

	return &std.String{Value: string(buf[:n])}
}

// fwrite writes a string to the file handle.
//
// Syntax: fwrite(file_handle, content_string)
func fwrite(rt std.Runtime, writer io.Writer, args ...std.GoMixObject) std.GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: fwrite expects 2 arguments (handle, content)")
	}
	if args[0].GetType() != std.FileType {
		return createError("ERROR: first argument to `fwrite` must be a file handle")
	}

	f := args[0].(*FileObject)
	content := args[1].ToString()

	n, err := f.Handle.WriteString(content)
	if err != nil {
		return createError("ERROR: write failed: %v", err)
	}

	return &std.Integer{Value: int64(n)}
}

// fseek sets the offset for the next Read or Write on the file.
//
// Syntax: fseek(file_handle, offset, whence)
// Whence: 0 (start), 1 (current), 2 (end)
func fseek(rt std.Runtime, writer io.Writer, args ...std.GoMixObject) std.GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: fseek expects 3 arguments (handle, offset, whence)")
	}
	if args[0].GetType() != std.FileType {
		return createError("ERROR: first argument to `fseek` must be a file handle")
	}

	f := args[0].(*FileObject)
	if args[1].GetType() != std.IntegerType {
		return createError("ERROR: second argument to `fseek` must be an integer (offset)")
	}
	offset := args[1].(*std.Integer).Value

	if args[2].GetType() != std.IntegerType {
		return createError("ERROR: third argument to `fseek` must be an integer (whence)")
	}
	whence := int(args[2].(*std.Integer).Value)

	newPos, err := f.Handle.Seek(offset, whence)
	if err != nil {
		return createError("ERROR: seek failed: %v", err)
	}

	return &std.Integer{Value: newPos}
}

// ftell returns the current offset of the file cursor.
//
// Syntax: ftell(file_handle)
func ftell(rt std.Runtime, writer io.Writer, args ...std.GoMixObject) std.GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: ftell expects 1 argument")
	}
	if args[0].GetType() != std.FileType {
		return createError("ERROR: argument to `ftell` must be a file handle")
	}

	f := args[0].(*FileObject)
	pos, err := f.Handle.Seek(0, io.SeekCurrent)
	if err != nil {
		return createError("ERROR: ftell failed: %v", err)
	}

	return &std.Integer{Value: pos}
}
