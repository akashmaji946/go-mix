/*
File    : go-mix/std/http.go
Author  : Akash Maji
Contact : akashmaji(@iisc.ac.in)
*/

// Package std - http.go
// This file defines HTTP client builtin functions.
package std

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var httpMethods = []*Builtin{
	{Name: "get_http", Callback: httpGet},           // Performs an HTTP GET request
	{Name: "post_http", Callback: httpPost},         // Performs an HTTP POST request
	{Name: "put_http", Callback: httpPut},           // Performs an HTTP PUT request
	{Name: "delete_http", Callback: httpDelete},     // Performs an HTTP DELETE request
	{Name: "listen_http", Callback: listenHttp},     // Starts an HTTP server
	{Name: "create_server", Callback: createServer}, // Creates a new HTTP server
	{Name: "handle_server", Callback: handleServer}, // Registers a handler for a server
	{Name: "start_server", Callback: startServer},   // Starts the HTTP server
	{Name: "request_http", Callback: httpRequest},   // Performs a generic HTTP request
	{Name: "serve_static", Callback: serveStatic},   // Serves static files from a directory
	{Name: "url_encode", Callback: urlEncode},       // URL encodes a string
	{Name: "url_decode", Callback: urlDecode},       // URL decodes a string
	{Name: "download_file", Callback: downloadFile}, // Downloads a file from a URL
}

func init() {
	Builtins = append(Builtins, httpMethods...)

	httpPackage := &Package{
		Name:      "http",
		Functions: make(map[string]*Builtin),
	}
	for _, method := range httpMethods {
		httpPackage.Functions[method.Name] = method
	}
	RegisterPackage(httpPackage)
}

// Server represents an HTTP server with a multiplexer.
type Server struct {
	Mux *http.ServeMux
}

func (s *Server) GetType() GoMixType {
	return ServerType
}

func (s *Server) ToString() string {
	return "<server>"
}

func (s *Server) ToObject() string {
	return "<server>"
}

// httpGet performs a GET request to the specified URL.
// Returns the response body as a string.
// Syntax: get_http(url)
func httpGet(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: get_http expects 1 argument (url)")
	}
	url := args[0].ToString()

	resp, err := http.Get(url)
	if err != nil {
		return createError("ERROR: get_http failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return createError("ERROR: failed to read response body: %v", err)
	}

	return &String{Value: string(body)}
}

// httpRequest performs a generic HTTP request with custom headers and body.
// Returns a map containing status, headers, and body.
// Syntax: request_http(method, url, [headers], [body])
func httpRequest(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) < 2 || len(args) > 4 {
		return createError("ERROR: request_http expects 2 to 4 arguments (method, url, [headers], [body])")
	}
	method := strings.ToUpper(args[0].ToString())
	urlStr := args[1].ToString()

	var bodyReader io.Reader
	if len(args) == 4 {
		bodyReader = strings.NewReader(args[3].ToString())
	}

	req, err := http.NewRequest(method, urlStr, bodyReader)
	if err != nil {
		return createError("ERROR: failed to create request: %v", err)
	}

	if len(args) >= 3 && args[2].GetType() != NilType {
		if args[2].GetType() != MapType {
			return createError("ERROR: headers argument must be a map")
		}
		headers := args[2].(*Map)
		for _, k := range headers.Keys {
			req.Header.Set(k, headers.Pairs[k].ToString())
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return createError("ERROR: request failed: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return createError("ERROR: failed to read response body: %v", err)
	}

	// Construct response map
	respMap := &Map{
		Pairs: make(map[string]GoMixObject),
		Keys:  []string{},
	}
	addKV := func(k string, v GoMixObject) {
		respMap.Pairs[k] = v
		respMap.Keys = append(respMap.Keys, k)
	}

	addKV("status", &Integer{Value: int64(resp.StatusCode)})
	addKV("body", &String{Value: string(respBody)})

	headersMap := &Map{
		Pairs: make(map[string]GoMixObject),
		Keys:  []string{},
	}
	for k, v := range resp.Header {
		val := &String{Value: strings.Join(v, ", ")}
		headersMap.Pairs[k] = val
		headersMap.Keys = append(headersMap.Keys, k)
	}
	addKV("headers", headersMap)

	return respMap
}

// serveStatic registers a handler to serve static files from a directory.
// Syntax: serve_static(server, prefix, root_path)
// Example: serve_static(srv, "/static/", "./public") will serve files from the "public" directory at the "/static/" URL path.
func serveStatic(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: serve_static expects 3 arguments (server, prefix, root_path)")
	}
	if args[0].GetType() != ServerType {
		return createError("ERROR: first argument to serve_static must be a server")
	}
	server := args[0].(*Server)
	prefix := args[1].ToString()
	root := args[2].ToString()

	fs := http.FileServer(http.Dir(root))
	server.Mux.Handle(prefix, http.StripPrefix(prefix, fs))
	return &Nil{}
}

// urlEncode encodes a string for safe use in URLs.
// Syntax: url_encode(str)
// Example: url_encode("Hello World!") returns "Hello%20World%21"
func urlEncode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: url_encode expects 1 argument")
	}
	return &String{Value: url.QueryEscape(args[0].ToString())}
}

// urlDecode decodes a URL-encoded string.
// Syntax: url_decode(str)
// Example: url_decode("Hello%20World%21") returns "Hello World!"
func urlDecode(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: url_decode expects 1 argument")
	}
	res, err := url.QueryUnescape(args[0].ToString())
	if err != nil {
		return createError("ERROR: url_decode failed: %v", err)
	}
	return &String{Value: res}
}

// downloadFile downloads a file from a URL to a local path.
// Syntax: download_file(url, path)
// Example: download_file("https://example.com/file.txt", "local_file.txt")
func downloadFile(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: download_file expects 2 arguments (url, path)")
	}
	urlStr := args[0].ToString()
	path := args[1].ToString()

	resp, err := http.Get(urlStr)
	if err != nil {
		return createError("ERROR: download failed: %v", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(path)
	if err != nil {
		return createError("ERROR: failed to create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return createError("ERROR: failed to write file: %v", err)
	}
	return &Nil{}
}

// listenHttp starts an HTTP server on the specified address.
// It blocks execution.
// Syntax: listen_http(address, handler_function)
func listenHttp(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: listen_http expects 2 arguments (address, handler)")
	}
	address := args[0].ToString()
	handler := args[1]

	if handler.GetType() != FunctionType {
		return createError("ERROR: second argument to listen_http must be a function")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", createHttpHandler(rt, handler))

	err := http.ListenAndServe(address, mux)
	if err != nil {
		return createError("ERROR: listen_http failed: %v", err)
	}

	return &Nil{}
}

// createServer creates a new HTTP server object.
// Syntax: create_server()
func createServer(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 0 {
		return createError("ERROR: create_server expects 0 arguments")
	}
	return &Server{Mux: http.NewServeMux()}
}

// handleServer registers a handler function for a specific pattern on the server.
// Syntax: handle_server(server, pattern, handler)
func handleServer(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: handle_server expects 3 arguments (server, pattern, handler)")
	}
	if args[0].GetType() != ServerType {
		return createError("ERROR: first argument to handle_server must be a server")
	}
	server := args[0].(*Server)
	pattern := args[1].ToString()
	handler := args[2]

	if handler.GetType() != FunctionType {
		return createError("ERROR: third argument to handle_server must be a function")
	}

	server.Mux.HandleFunc(pattern, createHttpHandler(rt, handler))
	return &Nil{}
}

// startServer starts the HTTP server on the specified address.
// Syntax: start_server(server, address)
// Example:
//
// var srv = create_server();
// handle_server(srv, "/hello", func(req) { return "Hello, World!" });
// start_server(srv, ":8081");
func startServer(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 2 {
		return createError("ERROR: start_server expects 2 arguments (server, address)")
	}
	if args[0].GetType() != ServerType {
		return createError("ERROR: first argument to start_server must be a server")
	}
	server := args[0].(*Server)
	address := args[1].ToString()

	err := http.ListenAndServe(address, server.Mux)
	if err != nil {
		return createError("ERROR: start_server failed: %v", err)
	}

	return &Nil{}
}

// createHttpHandler creates a net/http HandlerFunc from a Go-Mix handler function.
func createHttpHandler(rt Runtime, handler GoMixObject) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Construct the request map
		reqMap := &Map{
			Pairs: make(map[string]GoMixObject),
			Keys:  []string{},
		}

		// Helper to add key-value to map
		addKV := func(k string, v GoMixObject) {
			reqMap.Pairs[k] = v
			reqMap.Keys = append(reqMap.Keys, k)
		}

		addKV("method", &String{Value: r.Method})
		addKV("url", &String{Value: r.URL.String()})
		addKV("path", &String{Value: r.URL.Path})
		addKV("host", &String{Value: r.Host})
		addKV("protocol", &String{Value: r.Proto})

		// Headers
		headersMap := &Map{
			Pairs: make(map[string]GoMixObject),
			Keys:  []string{},
		}
		for k, v := range r.Header {
			val := &String{Value: strings.Join(v, ", ")}
			headersMap.Pairs[k] = val
			headersMap.Keys = append(headersMap.Keys, k)
		}
		addKV("headers", headersMap)

		// Body
		bodyBytes, _ := io.ReadAll(r.Body)
		addKV("body", &String{Value: string(bodyBytes)})

		// Call the Go-Mix handler
		result := rt.CallFunction(handler, reqMap)

		// Handle the response
		if result.GetType() == ErrorType {
			http.Error(w, result.ToString(), http.StatusInternalServerError)
			return
		}

		statusCode := 200
		responseBody := ""

		// Check if result is a map (structured response) or just a string/other (body)
		if resMap, ok := result.(*Map); ok {
			// Check for status
			if s, ok := resMap.Pairs["status"]; ok {
				if sInt, ok := s.(*Integer); ok {
					statusCode = int(sInt.Value)
				}
			}
			// Check for body
			if b, ok := resMap.Pairs["body"]; ok {
				responseBody = b.ToString()
			}
			// Check for headers
			if h, ok := resMap.Pairs["headers"]; ok {
				if hMap, ok := h.(*Map); ok {
					for _, k := range hMap.Keys {
						w.Header().Set(k, hMap.Pairs[k].ToString())
					}
				}
			}
		} else {
			responseBody = result.ToString()
		}

		w.WriteHeader(statusCode)
		w.Write([]byte(responseBody))
	}
}

// httpPost performs a POST request.
// Syntax: post_http(url, content_type, body)
func httpPost(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: post_http expects 3 arguments (url, content_type, body)")
	}
	url := args[0].ToString()
	contentType := args[1].ToString()
	bodyContent := args[2].ToString()

	resp, err := http.Post(url, contentType, strings.NewReader(bodyContent))
	if err != nil {
		return createError("ERROR: post_http failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return createError("ERROR: failed to read response body: %v", err)
	}

	return &String{Value: string(body)}
}

// httpPut performs a PUT request.
// Syntax: put_http(url, content_type, body)
func httpPut(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 3 {
		return createError("ERROR: put_http expects 3 arguments (url, content_type, body)")
	}
	url := args[0].ToString()
	contentType := args[1].ToString()
	bodyContent := args[2].ToString()

	req, err := http.NewRequest("PUT", url, strings.NewReader(bodyContent))
	if err != nil {
		return createError("ERROR: failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", contentType)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return createError("ERROR: put_http failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return createError("ERROR: failed to read response body: %v", err)
	}

	return &String{Value: string(body)}
}

// httpDelete performs a DELETE request.
// Syntax: delete_http(url)
func httpDelete(rt Runtime, writer io.Writer, args ...GoMixObject) GoMixObject {
	if len(args) != 1 {
		return createError("ERROR: delete_http expects 1 argument (url)")
	}
	url := args[0].ToString()

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return createError("ERROR: failed to create request: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return createError("ERROR: delete_http failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return createError("ERROR: failed to read response body: %v", err)
	}

	return &String{Value: string(body)}
}
