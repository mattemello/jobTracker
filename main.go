package main

import (
	"fmt"
	"github.com/mattemello/jobTracker/tmp/server"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start the Server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.HandlerMainPage(&w)
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
