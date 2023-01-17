package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	AWS_ACCESS_KEY_ID     = "AKIAUDSMZPA6FOJTM5MJ"
	AWS_SECRET_ACCESS_KEY = "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"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", HandleRoot)
	fmt.Println("Starting a simple go server on port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
