package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/service-b")
	if err != nil {
		http.Error(w, "Error calling Service B", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading response from Service B", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Response from Service B: %s", string(body))
}

func main() {
	http.HandleFunc("/service-a", handler)
	log.Println("Service A listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

