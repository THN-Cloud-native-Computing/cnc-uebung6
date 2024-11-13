package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Service B!")
}

func main() {
	http.HandleFunc("/service-b", handler)
	log.Println("Service B listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

