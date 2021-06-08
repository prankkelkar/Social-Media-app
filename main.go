package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", homeHandler).Methods("GET")
	myRoute.HandleFunc("/users", AllUsers).Methods("GET")
	myRoute.HandleFunc("/user/{name}/{email}", Newuser).Methods("POST")
	myRoute.HandleFunc("/users/{name}", Deluser).Methods("DELETE")
	myRoute.HandleFunc("/user/{name}/{email}", Deluser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRoute))
}

func main() {
	fmt.Println("hello: initiating server")
	handleRequest()
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")

}
