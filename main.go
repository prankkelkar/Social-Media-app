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

	// users := User{
	// 	Name: "pk", Email: "pk@gmail.com", Add: "hn297", Profile: Profile{Hobbies: []string{"cycle", "run"}, Languages: []string{"hindi", "marathi"}}}
	// // people := []Person{{First: "pk", Age: 25}, {First: "sk", Age: 12}}
	db.AutoMigrate(&User{})
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
