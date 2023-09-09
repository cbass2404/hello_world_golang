package main

import (
	"fmt"
	"net/http"

	"github.com/cbass2404/hello_world_golang/pkg/handlers"
)

const portNumber = ":8080"

// main is the main application funciton
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
