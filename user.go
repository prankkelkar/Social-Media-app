package main

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type MultiString []string

type CreditCard struct {
	gorm.Model
	Number  string
	CuserID uint
}

type Hobbies struct {
	Hname     string `gorm:"primaryKey; not null"`
	ProfileID int    `gorm:"primaryKey; not null"`
}

type Languages struct {
	Lname     string `gorm:"primaryKey; not null"`
	ProfileID int    `gorm:"primaryKey; not null"`
}
type Profile struct {
	gorm.Model
	Hobbies   []Hobbies   `json:"Hobbies"`
	Languages []Languages `json:"Languages"`
	UserID    int
}

type Skill struct {
	Sname    string `gorm:"primaryKey; not null"`
	PersonID int    `gorm:"primaryKey; not null"`
}

type User struct {
	gorm.Model
	Name    string `json:"Name"`
	Email   string `json:"Email"`
	Add     string `json:"Address"`
	Profile Profile
}

type Person struct {
	gorm.Model
	First  string
	Skills []Skill
	Age    int
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	//Way to go
	//Using gorm get the list of users and store it in a slice
	// users := []User{
	// 	{"pk", "pk@gmail.com", "hn297", Profile{Hobbies: []string{"cycle", "run"}, Languages: []string{"hindi", "marathi"}}},
	// 	{"gk", "gk@gmail.com", "hn757", Profile{Hobbies: []string{"chess", "walk"}, Languages: []string{"hindi", "marathi"}}},
	// }

	// //Encode the list in json and print on the screen.
	// json.NewEncoder(w).Encode(users)
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
