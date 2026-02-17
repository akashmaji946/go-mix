/*
File    : go-mix/std/crypto.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

package std

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
)

var cryptoMethods = []*Builtin{
	{Name: "md5", Callback: md5Func},
	{Name: "sha1", Callback: sha1Func},
	{Name: "sha256", Callback: sha256Func},
	{Name: "base64_encode", Callback: base64Encode},
	{Name: "base64_decode", Callback: base64Decode},
	{Name: "hex_encode", Callback: hexEncode},
	{Name: "hex_decode", Callback: hexDecode},
	{Name: "uuid", Callback: uuidFunc},
	{Name: "random", Callback: randomFunc},
}

func init() {
	cryptoPackage := &Package{
		Name:      "crypto",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range cryptoMethods {
		cryptoPackage.Functions[method.Name] = method
	}
	RegisterPackage(cryptoPackage)
}

func md5Func(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: md5 expects 1 argument (string)")
	}
	data := args[0].ToString()
	hash := md5.Sum([]byte(data))
	return &String{Value: fmt.Sprintf("%x", hash)}
}

func sha1Func(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: sha1 expects 1 argument (string)")
	}
	data := args[0].ToString()
	hash := sha1.Sum([]byte(data))
	return &String{Value: fmt.Sprintf("%x", hash)}
}

func sha256Func(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: sha256 expects 1 argument (string)")
	}
	data := args[0].ToString()
	hash := sha256.Sum256([]byte(data))
	return &String{Value: fmt.Sprintf("%x", hash)}
}

func base64Encode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: base64_encode expects 1 argument (string)")
	}
	data := args[0].ToString()
	encoded := base64.StdEncoding.EncodeToString([]byte(data))
	return &String{Value: encoded}
}

func base64Decode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: base64_decode expects 1 argument (string)")
	}
	data := args[0].ToString()
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return createError("ERROR: failed to decode base64: %v", err)
	}
	return &String{Value: string(decoded)}
}

func hexEncode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: hex_encode expects 1 argument (string)")
	}
	data := args[0].ToString()
	encoded := hex.EncodeToString([]byte(data))
	return &String{Value: encoded}
}

func hexDecode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: hex_decode expects 1 argument (string)")
	}
	data := args[0].ToString()
	decoded, err := hex.DecodeString(data)
	if err != nil {
		return createError("ERROR: failed to decode hex: %v", err)
	}
	return &String{Value: string(decoded)}
}

func uuidFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: uuid expects 0 arguments")
	}
	u := make([]byte, 16)
	_, err := rand.Read(u)
	if err != nil {
		return createError("ERROR: failed to generate UUID: %v", err)
	}
	// Set version (4) and variant (RFC 4122)
	u[6] = (u[6] & 0x0f) | 0x40
	u[8] = (u[8] & 0x3f) | 0x80

	return &String{Value: fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])}
}

func randomFunc(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: random expects 1 argument (number of bytes)")
	}
	n, ok := args[0].(*Integer)
	if !ok {
		return createError("ERROR: argument to `random` must be a number, got '%s'", args[0].GetType())
	}
	if n.Value < 0 {
		return createError("ERROR: number of bytes must be non-negative")
	}
	bytes := make([]byte, int(n.Value))
	_, err := rand.Read(bytes)
	if err != nil {
		return createError("ERROR: failed to generate random bytes: %v", err)
	}
	return &String{Value: string(bytes)}
}
