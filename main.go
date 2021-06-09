package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
	dsn := "pk:pk@tcp(9.30.95.8:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	// h1 := []Hobbies{
	// 	{Hname: "chess"},
	// 	{Hname: "cook"},
	// }

	// h2 := []Hobbies{
	// 	{Hname: "roam"},
	// 	{Hname: "sleep"},
	// }

	// l1 := []Languages{
	// 	{Lname: "chess"},
	// 	{Lname: "cook"},
	// }

	// l2 := []Languages{
	// 	{Lname: "roam"},
	// 	{Lname: "sleep"},
	// }

	// users := []User{
	// 	{
	// 		Name:    "pk",
	// 		Email:   "pk@gmail.com",
	// 		Add:     "hn297",
	// 		Profile: Profile{Hobbies: h1, Languages: l1},
	// 	},
	// 	{
	// 		Name:    "jk",
	// 		Email:   "jk@gmail.com",
	// 		Add:     "mnt76",
	// 		Profile: Profile{Hobbies: h2, Languages: l1},
	// 	},
	// 	{
	// 		Name:    "vb",
	// 		Email:   "vb@gmail.com",
	// 		Add:     "jhjh",
	// 		Profile: Profile{Hobbies: h2, Languages: l2},
	// 	},
	// }

	// fmt.Println(users)
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Profile{})
	// db.AutoMigrate(&Languages{})
	//db.AutoMigrate(&Hobbies{})
	// result := db.Create(&users)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	fmt.Println("Writing to the database is completed")
	// fmt.Println(result.RowsAffected)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")

}
