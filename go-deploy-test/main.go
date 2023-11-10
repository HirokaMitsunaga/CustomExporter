package main

import (
	"fmt"
	"log"
	"net/http"
)
 
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello World!")
	log.Println("Hello World! has been sent to a client.")
}
 
func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8888", nil))
}