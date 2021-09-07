package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Social-Media-app/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func handleRequest() {
	myRoute := mux.NewRouter().StrictSlash(true)
	myRoute.HandleFunc("/", homeHandler).Methods("GET")
	myRoute.HandleFunc("/users", user.AllUsers).Methods("GET")
	myRoute.HandleFunc("/profiles", user.AllProfiles).Methods("GET")
	myRoute.HandleFunc("/user/{user_id}/profile", user.SpecificProfile).Methods(http.MethodGet)
	myRoute.HandleFunc("/user", user.Newuser).Methods("POST")
	myRoute.HandleFunc("/user/delete/{user_id}", user.Deluser).Methods("DELETE")
	myRoute.HandleFunc("/user/create/profile/{user_id}", user.UpdateUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRoute))
}

func main() {
	//Use this function only first time.
	//initialiseDB()

	fmt.Println("Initiating server")
	handleRequest()
}

func initialiseDB() {
	dsn := "pk:pk@tcp(9.30.95.8:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	//sample examples

	h1 := []user.Hobbies{
		{Hname: "chess"},
		{Hname: "cook"},
	}

	h2 := []user.Hobbies{
		{Hname: "roam"},
		{Hname: "sleep"},
	}

	l1 := []user.Languages{
		{Lname: "eng"},
		{Lname: "kan"},
	}

	l2 := []user.Languages{
		{Lname: "mara"},
		{Lname: "mara"},
	}
	users := []user.User{
		{
			Name:    "pk",
			Email:   "pk@gmail.com",
			Add:     "hn297",
			Profile: user.Profile{Hobbies: h1, Languages: l1},
		},
		{
			Name:    "jk",
			Email:   "jk@gmail.com",
			Add:     "mnt76",
			Profile: user.Profile{Hobbies: h2, Languages: l1},
		},
		{
			Name:    "vb",
			Email:   "vb@gmail.com",
			Add:     "jhjh",
			Profile: user.Profile{Hobbies: h2, Languages: l2},
		},
	}
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.Profile{})
	db.AutoMigrate(&user.Languages{})
	db.AutoMigrate(&user.Hobbies{})

	result := db.Create(&users)
	if result.Error != nil {
		panic("messed up")
	}
	fmt.Println("Writing to the database is completed")

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")

}
