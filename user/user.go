package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Social-Media-app/database"
	"github.com/gorilla/mux"
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
	ProfileID uint   `gorm:"primaryKey"`
}

type Languages struct {
	Lname     string `json:"Lname" gorm:"primaryKey; not null"`
	ProfileID uint   `gorm:"primaryKey"`
}
type Profile struct {
	gorm.Model
	Hobbies   []Hobbies   `json:"Hobbies" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Languages []Languages `json:"Languages" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UserID    uint
}

type User struct {
	gorm.Model
	Name    string  `json:"Name"`
	Email   string  `json:"Email" gorm:"unique"`
	Add     string  `json:"Address"`
	Profile Profile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func AllUsers(w http.ResponseWriter, r *http.Request) {

	//Get data from mysql
	db := database.GetCon()
	var users []User
	// Get all records
	result := db.Preload("Profile").Preload("Profile.Languages").Preload("Profile.Hobbies").Find(&users)
	// SELECT * FROM users;
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Your request could not be processed"}`))
		panic("Issue with DB")
	}

	//Encode the list in json and print on the screen.
	json.NewEncoder(w).Encode(users)
}

func AllProfiles(w http.ResponseWriter, r *http.Request) {
	//Get data from mysql
	db := database.GetCon()

	var profiles []Profile
	// Get all records
	result := db.Preload("Languages").Preload("Hobbies").Find(&profiles)
	// SELECT * FROM users;
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"Your request could not be processed"}`))
		panic("Issue with DB")
	}

	//Encode the list in json and print on the screen.
	json.NewEncoder(w).Encode(profiles)
}

func SpecificProfile(w http.ResponseWriter, r *http.Request) {
	//Frame an object of type user and add it to the database
	var u User
	id := -1
	var err error
	pathParams := mux.Vars(r)
	if val, ok := pathParams["user_id"]; ok {
		id, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	db := database.GetCon()
	db.Preload("Profile").Preload("Profile.Languages").Preload("Profile.Hobbies").First(&u, id)
	if u.Profile.ID != 0 {
		json.NewEncoder(w).Encode(&u.Profile)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Profile is not present for the user yet."))
	}
}

func Newuser(w http.ResponseWriter, r *http.Request) {
	//Decode the json
	//Frame an object of type user and add it to the database
	var u User
	json.NewDecoder(r.Body).Decode(&u)
	db := database.GetCon()
	db.Create(&u)
	//fmt.Println(u)
	fmt.Fprintf(w, "Creating new user")
	json.NewEncoder(w).Encode(&u)

}

func Deluser(w http.ResponseWriter, r *http.Request) {
	id := -1
	var err error
	path_params := mux.Vars(r)
	if val, ok := path_params["user_id"]; ok {
		id, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}
	if id != -1 {
		db := database.GetCon()
		db.Delete(&User{}, id)
		fmt.Fprintf(w, "User deleted")
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	var p Profile
	var u User
	json.NewDecoder(r.Body).Decode(&p)
	db := database.GetCon()

	id := -1
	var err error
	path_params := mux.Vars(r)
	if val, ok := path_params["user_id"]; ok {
		id, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}
	if id != -1 {
		db.Preload("Profile").Preload("Profile.Languages").Preload("Profile.Hobbies").First(&u, id)
		//update Profile
		u.Profile = p
		db.Save(&u)
		json.NewEncoder(w).Encode(&u)
	}
}
