package main

import (
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/health", healthCheck)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
