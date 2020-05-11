package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	_, _ = io.WriteString(res, "cat cat cat")
}

func main() {

	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	_ = http.ListenAndServe(":8080", nil)
}

// this is similar to this:
// https://play.golang.org/p/X2dlgVSIrd
// ---and this---
// https://play.golang.org/p/YaUYR63b7L
