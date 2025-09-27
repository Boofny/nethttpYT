package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleOther(w http.ResponseWriter, r *http.Request) {
    log.Println("Received a non domain request")

    // Example response data
    resp := map[string]string{
        "hello": "Hello, stranger...",
    }

    // Set Content-Type header for JSON
    w.Header().Set("Content-Type", "application/json")

    // Encode and write JSON to response
    if err := json.NewEncoder(w).Encode(resp); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

// func handleOther(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Received a non domain request")
// 	w.Write([]byte("Hello, stranger..."))
// }

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at my domain")
	w.Write([]byte("Hello, Domain name!"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", handleOther)
	router.HandleFunc("dreamsofcode.foo/", handle)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port",server.Addr)
	server.ListenAndServe()
}
