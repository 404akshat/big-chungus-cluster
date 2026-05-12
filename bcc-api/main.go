package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "200 OK")
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
}
