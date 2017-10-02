package main

import (
	"fmt"
	"html"
	"net/http"
)

// AddNumber takes 2 integer and return sum if both values
func AddNumber(a int, b int) int {
	return a + b
}

// FooHandler is mock handler
func FooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "foo %T %q", r.URL.Path, html.EscapeString(r.URL.Path))
}

// IndexHandler handles root HTTP path
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello index")
}
