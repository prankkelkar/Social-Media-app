package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "root:pk@tcp(9.47.67.255:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("hello: initiating server")

	defer db.Close()
	//handleRequest()
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO egg VALUES ('from code')")

	defer insert.Close()
	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")

}
