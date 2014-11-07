package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	log.Println("Running web server at 8080")
	http.ListenAndServe(":8080", nil)
}