package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started!")

	http.ListenAndServe(":8081", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("new request")
		fmt.Fprintf(w, "Hello world")
	}))
}
