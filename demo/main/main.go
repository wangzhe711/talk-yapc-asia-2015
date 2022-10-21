package main

import (
	"log"
	"net/http"
	"talk-demo/demo"
)

func main() {
	log.Printf("Starting on port 8080")
	http.HandleFunc("/hi", demo.HandleHi)
	// why 0.0.0.0?
	// Basically, that tells the server to listen on all networks, not just to localhost connections,
	// otherwise the server is only accessible from
	// within the container (which is the local host for the server running there)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
