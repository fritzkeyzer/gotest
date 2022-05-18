package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handleEndpoint)
	http.ListenAndServe(":8090", nil)
}

func handleEndpoint(w http.ResponseWriter, req *http.Request){
	log.Println("req received")
	fmt.Fprintf(w, "hello\n")
}