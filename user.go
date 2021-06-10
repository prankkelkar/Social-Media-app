package main

import (
	"encoding/json"
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
	Hname     string `json:"Hname" gorm:"primaryKey; not null"`
	ProfileID int    `gorm:"primaryKey; not null"`
}

type Languages struct {
	Lname     string `json:"Lname" gorm:"primaryKey; not null"`
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
	Email   string `json:"Email" gorm:"unique"`
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

	//Get data from mysql
	db := GetCon()
	var users []User
	// Get all records
	result := db.Preload("Profile").Preload("Profile.Languages").Preload("Profile.Hobbies").Find(&users)
	// SELECT * FROM users;
	if result.Error != nil {
		panic("Issue with DB")
	}

	//Encode the list in json and print on the screen.
	json.NewEncoder(w).Encode(users)
}

func AllProfiles(w http.ResponseWriter, r *http.Request) {
	//Get data from mysql
	db := GetCon()

	var profiles []Profile
	// Get all records
	result := db.Preload("Languages").Preload("Hobbies").Find(&profiles)
	// SELECT * FROM users;
	if result.Error != nil {
		panic("Issue with DB")
	}

	//Encode the list in json and print on the screen.
	json.NewEncoder(w).Encode(profiles)
}

func SpecificProfile(w http.ResponseWriter, r *http.Request) {
}

func Newuser(w http.ResponseWriter, r *http.Request) {
	//Decode the json
	//Frame an object of type user and add it to the database
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	db := GetCon()
	db.Create(&u)
	//fmt.Println(u)
	fmt.Fprintf(w, "User reached me")

}

func Deluser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "del users endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update user users endpoint hit")
}
