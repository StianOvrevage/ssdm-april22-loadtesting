package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func Health(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "I'm OK!")
}

func main() {
	// Get optional configurations
	listen := ":8080"
	if os.Getenv("LISTEN") != ""{
		listen = os.Getenv("LISTEN")
	}

    // Configure and start webserver
    http.HandleFunc("/health", Health)
    fmt.Printf("Starting webserver on %v\n", listen)
    log.Fatal(http.ListenAndServe(listen, nil))
}