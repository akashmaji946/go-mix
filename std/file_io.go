/*
File    : go-mix/std/file_io.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - file_io.go
// This file defines the file system builtin functions for the GoMix language.
// It provides functions for reading, writing, and manipulating files and directories.
package std

import (
	"fmt"
	"io"
	"os"
	"time"
)

var fileIOMethods = []*Builtin{
	{Name: "read_file", Callback: readFile},         // Reads entire file content as string
	{Name: "write_file", Callback: writeFile},       // Writes string to a file
	{Name: "append_file", Callback: appendFile},     // Appends string to a file
	{Name: "file_exists", Callback: fileExists},     // Checks if a file or directory exists
	{Name: "is_dir", Callback: isDir},               // Checks if path is a directory
	{Name: "is_file", Callback: isFile},             // Checks if path is a regular file
	{Name: "mkdir", Callback: mkdir},                // Creates a directory (and parents)
	{Name: "remove_file", Callback: removeFile},     // Removes a file or directory
	{Name: "touch", Callback: touch},                // Creates an empty file or updates timestamps
	{Name: "list_dir", Callback: listDir},           // Returns an array of names in a directory
	{Name: "pwd", Callback: pwd},                    // Returns the current working directory
	{Name: "home", Callback: home},                  // Returns the user's home directory
	{Name: "truncate_file", Callback: truncateFile}, // Changes the size of a file
	{Name: "remove_all", Callback: removeAll},       // Removes path and any children (rm -rf)
	{Name: "rename_file", Callback: renameFile},     // Renames or moves a file/directory
	{Name: "chmod", Callback: chmod},                // Changes file permissions
	{Name: "cat", Callback: cat},                    // Prints file content to output
}

// init registers the file I/O methods as global builtins.
func init() {
	Builtins = append(Builtins, fileIOMethods...)
}

// readFile reads the entire contents of a file into a string.
//
// Syntax: read_file(path)
func readFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: read_file expects 1 argument (path)")
	}
	path := args[0].ToString()
	content, err := os.ReadFile(path)
	if err != nil {
		return createError("ERROR: could not read file '%s': %v", path, err)
	}
	return &String{Value: string(content)}
}

// writeFile writes a string to a file, creating it if it doesn't exist.
//
// Syntax: write_file(path, content)
func writeFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: write_file expects 2 arguments (path, content)")
	}
	path := args[0].ToString()
	data := args[1].ToString()
	err := os.WriteFile(path, []byte(data), 0644)
	if err != nil {
		return createError("ERROR: could not write to file '%s': %v", path, err)
	}
	return &Nil{}
}

// cat reads the contents of one or more files and prints them to the output writer.
//
// Syntax: cat(path1, [path2, ...])
//
// Example:
//
//	cat("build.sh");
func cat(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) == 0 {
		return createError("ERROR: cat expects at least 1 argument (path)")
	}
	for _, arg := range args {
		path := arg.ToString()
		content, err := os.ReadFile(path)
		if err != nil {
			return createError("ERROR: could not read file '%s': %v", path, err)
		}
		fmt.Fprint(writer, string(content))
	}
	if flusher, ok := writer.(interface{ Sync() error }); ok {
		flusher.Sync()
	}
	return &Nil{}
}

// touch creates an empty file if it doesn't exist, or updates the access and
// modification times to the current time if it does.
//
// Syntax: touch(path)
func touch(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: touch expects 1 argument (path)")
	}
	path := args[0].ToString()

	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		f, err := os.Create(path)
		if err != nil {
			return createError("ERROR: could not create file '%s': %v", path, err)
		}
		f.Close()
	} else {
		currentTime := time.Now()
		if err := os.Chtimes(path, currentTime, currentTime); err != nil {
			return createError("ERROR: could not update timestamps for '%s': %v", path, err)
		}
	}
	return &Nil{}
}

// listDir returns an array containing the names of the entries in the directory.
//
// Syntax: list_dir(path)
func listDir(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: list_dir expects 1 argument (path)")
	}
	path := args[0].ToString()
	entries, err := os.ReadDir(path)
	if err != nil {
		return createError("ERROR: could not read directory '%s': %v", path, err)
	}

	elements := make([]GoMixObject, len(entries))
	for i, entry := range entries {
		elements[i] = &String{Value: entry.Name()}
	}
	return &Array{Elements: elements}
}

// pwd returns the current working directory path.
//
// Syntax: pwd()
//
// Example:
//
//	var current = pwd();
func pwd(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: pwd expects 0 arguments")
	}

	dir, err := os.Getwd()
	if err != nil {
		return createError("ERROR: could not get current working directory: %v", err)
	}

	return &String{Value: dir}
}

// truncateFile changes the size of the named file.
//
// Syntax: truncate_file(path, size)
//
// Example:
//
//	truncate_file("data.log", 0); // Clears the file content
func truncateFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: truncate_file expects 2 arguments (path, size)")
	}
	if args[1].GetType() != IntegerType {
		return createError("ERROR: size must be an integer, got %s", args[1].GetType())
	}

	path := args[0].ToString()
	size := args[1].(*Integer).Value

	err := os.Truncate(path, size)
	if err != nil {
		return createError("ERROR: could not truncate file '%s': %v", path, err)
	}
	return &Nil{}
}

// removeAll removes path and any children it contains.
// It is the GoMix equivalent of 'rm -rf'.
//
// Syntax: remove_all(path)
func removeAll(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: remove_all expects 1 argument (path)")
	}
	path := args[0].ToString()
	err := os.RemoveAll(path)
	if err != nil {
		return createError("ERROR: could not remove '%s': %v", path, err)
	}
	return &Nil{}
}

// renameFile renames (moves) oldpath to newpath.
// If newpath already exists and is not a directory, it is replaced.
//
// Syntax: rename_file(oldpath, newpath)
func renameFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: rename_file expects 2 arguments (old, new)")
	}
	oldPath := args[0].ToString()
	newPath := args[1].ToString()

	err := os.Rename(oldPath, newPath)
	if err != nil {
		return createError("ERROR: could not rename '%s' to '%s': %v", oldPath, newPath, err)
	}
	return &Nil{}
}

// chmod changes the mode of the named file to mode.
//
// Syntax: chmod(path, mode_int)
//
// Example:
//
//	chmod("script.gm", 0755);
func chmod(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: chmod expects 2 arguments (path, mode)")
	}
	if args[1].GetType() != IntegerType {
		return createError("ERROR: mode must be an integer, got %s", args[1].GetType())
	}

	path := args[0].ToString()
	mode := args[1].(*Integer).Value

	err := os.Chmod(path, os.FileMode(mode))
	if err != nil {
		return createError("ERROR: could not change mode for '%s': %v", path, err)
	}
	return &Nil{}
}

// home returns the current user's home directory path.
//
// Syntax: home()
//
// Example:
//
//	var userHome = home();
func home(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: home expects 0 arguments")
	}

	dir, err := os.UserHomeDir()
	if err != nil {
		return createError("ERROR: could not get home directory: %v", err)
	}

	return &String{Value: dir}
}

// appendFile appends a string to a file, creating it if it doesn't exist.
//
// Syntax: append_file(path, content)
func appendFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: append_file expects 2 arguments (path, content)")
	}
	path := args[0].ToString()
	data := args[1].ToString()
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return createError("ERROR: could not open file '%s' for appending: %v", path, err)
	}
	defer f.Close()
	if _, err := f.WriteString(data); err != nil {
		return createError("ERROR: could not write to file '%s': %v", path, err)
	}
	return &Nil{}
}

// fileExists checks if a file or directory exists at the given path.
//
// Syntax: file_exists(path)
func fileExists(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: file_exists expects 1 argument")
	}
	path := args[0].ToString()
	_, err := os.Stat(path)
	return &Boolean{Value: !os.IsNotExist(err)}
}

// isDir checks if the given path is a directory.
//
// Syntax: is_dir(path)
func isDir(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: is_dir expects 1 argument")
	}
	path := args[0].ToString()
	info, err := os.Stat(path)
	if err != nil {
		return &Boolean{Value: false}
	}
	return &Boolean{Value: info.IsDir()}
}

// isFile checks if the given path is a regular file.
//
// Syntax: is_file(path)
func isFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: is_file expects 1 argument")
	}
	path := args[0].ToString()
	info, err := os.Stat(path)
	if err != nil {
		return &Boolean{Value: false}
	}
	return &Boolean{Value: !info.IsDir()}
}

// mkdir creates a new directory and any necessary parent directories.
//
// Syntax: mkdir(path)
func mkdir(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: mkdir expects 1 argument")
	}
	path := args[0].ToString()
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return createError("ERROR: could not create directory '%s': %v", path, err)
	}
	return &Nil{}
}

// removeFile removes a file or a directory.
//
// Syntax: remove_file(path, [force])
//
// Usage:
//
//	Removes the file or directory at the specified path.
//	If force is true, it removes the directory and all its contents recursively.
//	If force is false (default), it only removes the file or an empty directory.
//
// Example:
//
//	remove_file("temp.txt");
//	remove_file("old_dir", true);
func removeFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 1 || len(args) > 2 {
		return createError("ERROR: remove_file expects 1 or 2 arguments")
	}
	path := args[0].ToString()
	force := false
	if len(args) == 2 {
		if args[1].GetType() != BooleanType {
			return createError("ERROR: second argument to `remove_file` must be a boolean (force)")
		}
		force = args[1].(*Boolean).Value
	}

	var err error
	if force {
		err = os.RemoveAll(path)
	} else {
		err = os.Remove(path)
	}

	if err != nil {
		return createError("ERROR: could not remove '%s': %v", path, err)
	}
	return &Nil{}
}
