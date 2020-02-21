package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func MyServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}
func main() {
	http.HandleFunc("/hello", MyServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
