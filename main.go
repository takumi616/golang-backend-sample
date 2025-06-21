package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	})

	port := os.Getenv("APP_PORT")
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
