package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	fmt.Println("Hello world!")
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Printf("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
