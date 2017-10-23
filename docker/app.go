package main

import (
	"fmt"
	"log"
	"net/http"

	as "github.com/aerospike/aerospike-client-go"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	http.Handle("/", r)
	fmt.Println("Starting up on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
	client, err := as.NewClient("127.0.0.1", 3000)
	if err != nil {
		fmt.Fprintln(w, "AS client was not created! Error: ", err)
	} else {
		fmt.Println(client)
	}
}
