// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// http://play.golang.org/p/LY5P96Xrbl

// https://gist.github.com/jmoiron/e9f72720cef51862b967#file-01-curl-go
// Sample code provided by Jason Moiron
//
// ./example2 http://www.goinggo.net/feeds/posts/default

// Sample program to show how to write a simple version of curl using
// the io.Reader and io.Writer interface support.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func exit(code int, a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
	os.Exit(code)
}

// init is called before main.
func init() {
	if len(os.Args) != 2 {
		exit(2, "Usage: ./example2 <url>")
	}
}

// main is the entry point for the application.
func main() {
	// resp here is a response, and resp.Body is an io.Reader
	resp, err := http.Get(os.Args[1])
	if err != nil {
		exit(1, "http protocol error:", err)
	}
	// Close the ReadCloser when we're done with it.
	// We don't need to check the error, since
	// Close errors on Readers are meaningless.
	defer resp.Body.Close()

	// io.Copy(dst io.Writer, src io.Reader) (int64, error)
	// Copies from the Body to Stdout, returning any Read or Write error.
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		exit(1, "stream copy error:", err)
	}
}
