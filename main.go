package main

import (
	"fmt"
	"net/http"

	"github.com/aurora1530/go-web-examples/internal"
)

func main() {
	fmt.Println("Starting server...")
	server := internal.NewServer()
	r := internal.CreateRouter(server)
	fmt.Println("Server started on port 80")
	http.ListenAndServe(":80", r)
}
