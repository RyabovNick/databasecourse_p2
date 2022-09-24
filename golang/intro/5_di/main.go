package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "Jack")
	http.ListenAndServe(":8080", http.HandlerFunc(GreetHandler))
}
