package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Profile struct {
	Hobbies   []string `json:"Hobbies"`
	Languages []string `json:"Languages"`
}

type User struct {
	///gorm.Model
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Add     string `json:"Address"`
	Profile `gorm:"embedded"`
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	//Way to go
	//Using gorm get the list of users and store it in a slice
	users := []User{
		{"pk", "pk@gmail.com", "hn297", Profile{Hobbies: []string{"cycle", "run"}, Languages: []string{"hindi", "marathi"}}},
		{"gk", "gk@gmail.com", "hn757", Profile{Hobbies: []string{"chess", "walk"}, Languages: []string{"hindi", "marathi"}}},
	}

	//Encode the list in json and print on the screen.
	json.NewEncoder(w).Encode(users)
}

func Newuser(w http.ResponseWriter, r *http.Request) {
	//Decode the json
	//Frame an object of type user and add it to the database
	fmt.Fprintf(w, "New users endpoint hit")

}

func Deluser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "del users endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update user users endpoint hit")
}
