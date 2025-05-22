package main

import (
	"fmt"
	"net/http"
)

// the main is the entry point of the Go program, where the execution of the application begins.
func main() {
	fmt.Println("The serve is launched!")

	http.Handle("/ping", NewPing())
	http.ListenAndServe(":80", nil)
}
