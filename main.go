package main

import "net/http"

func main() {
	mux := http.NewServeMux() // ServeMux routes requests. Empty mux = 404
	mux.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Handler: mux,     // Handler is the brain that routes all server requests
		Addr:    ":8080", // [host]:[port], :8080 means all interfaces listen on port 8080
	}
	server.ListenAndServe()
}
