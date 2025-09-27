package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dreamsofcode-io/nethttp/middleware"
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

func handle(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a request at my domain")
	w.Write([]byte("Hello, Domain name!"))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/test", handleOther)
	router.HandleFunc("dreamsofcode.foo/", handle)

	// e.Use(middleware.CORS())
	//use needs to be the the option then middleware.CORS will pass default values
	//i think we can imagen the stack as the use function and the 
	//middleware.CORS or middleware.CustomCORS as the logging and allow and logging 
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.AllowCors,//this is allowing cors
	)

	server := http.Server{
		Addr:    ":8081",
		Handler: stack(router),
	}

	fmt.Println("Server listening on port",server.Addr)
	server.ListenAndServe()
}
