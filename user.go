package main

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all users endpoint hit")
}

func Newuser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "New users endpoint hit")
}

func Deluser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "del users endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update user users endpoint hit")
}
