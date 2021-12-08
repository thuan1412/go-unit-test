package main

import (
	"fmt"
	"net/http"
)

func hello3(name string) string {
	if name == "" {
		return "Hello, what's your name?"
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}
func hello(name string) string {
	if name == "" {
		return "Hello, what's your name?"
	} else {
		return fmt.Sprintf("Hello %s", name)
	}
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Hello world")
}

func New(a string) string {
	return a
}

func main() {
	// http.HandleFunc("/", welcome)
	// http.ListenAndServe(":3000", nil)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Println(s == nil, len(s), cap(s), s[3])
}
