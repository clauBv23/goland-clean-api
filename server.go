package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func main() {	
	r := mux.NewRouter().StrictSlash(true)
	port:=":8000"

	r.HandleFunc("/",func(response http.ResponseWriter, request *http.Request){
		fmt.Fprintln(response, "Up and running  ")})

	r.HandleFunc("/posts", getPosts).Methods("GET")

	r.HandleFunc("/post", addPost).Methods("POST")

	log.Println("server running on port", port)
	log.Fatal(http.ListenAndServe(port, r))
}

