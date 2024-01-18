// Package declaration indicating that this Go file is part of the 'main' package.
package main

// Importing necessary packages from the standard library.
import (
	"fmt"      // Package for formatted I/O.
	"html"     // Package for HTML escaping.
	"log"      // Package for logging.
	"net/http" // Package for HTTP server functionality.
)

// The main function, which is the entry point of the program.
func main() {
	// Registering a handler function for the root path ("/").
	http.HandleFunc("/", rootHandler)

	// Registering a handler function for the "/hello" path.
	http.HandleFunc("/hello", helloHandler)

	// Starting the HTTP server on port 8080 and handling requests using the default ServeMux.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler function for the root path ("/").
func rootHandler(w http.ResponseWriter, r *http.Request) {
	// Responding with a formatted string containing the escaped URL path.
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// Handler function for the "/hello" path.
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Responding with a simple "Hello" message.
	fmt.Fprintf(w, "Hello")
}
