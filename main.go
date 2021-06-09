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
	user := User{
		Name:    "pk",
		Email:   "pk@gmail.com",
		Add:     "hn297",
		Profile: Profile{Hobbies: []Hobbies{{Hname: "cycle"}, {Hname: "football"}}, Languages: []Languages{{Lname: "eng"}, {Lname: "mar"}}},
	}
	// users := []User{
	// 	{
	// 		Name:    "pk",
	// 		Email:   "pk@gmail.com",
	// 		Add:     "hn297",
	// 		Profile: Profile{Hobbies: "cycling", Languages: "hindi"},
	// 	},
	// 	{
	// 		Name:    "jk",
	// 		Email:   "jk@gmail.com",
	// 		Add:     "mnt76",
	// 		Profile: Profile{Hobbies: "chess", Languages: "englis"},
	// 	},
	// 	{
	// 		Name:    "vb",
	// 		Email:   "vb@gmail.com",
	// 		Add:     "jhjh",
	// 		Profile: Profile{Hobbies: "reading", Languages: "marathi"},
	// 	},
	// }

	// fmt.Println(users)
	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Profile{})
	// result := db.Create(&users)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }
	// db.Model(&Person).Related(&Skill{})
	// db.AutoMigrate(&Person{})
	// db.AutoMigrate(&Skill{})
	// sk := []Skill{{Sname: "coding"}, {Sname: "talking"}, {Sname: "walking"}}
	// p := Person{First: "rajat", Age: 22, Skills: sk}

	db.AutoMigrate(&Profile{})
	db.AutoMigrate(&Languages{})
	db.AutoMigrate(&Hobbies{})
	db.AutoMigrate(&User{})
	// l1 := []Languages{{Lname: "eng"},
	// 	{Lname: "mm"},
	// }

	// p := Profile{Languages: l1}
	result := db.Create(&user)
	if result.Error != nil {
		panic("messed up")
	}
	fmt.Println("Writing to the database is completed")
	// fmt.Println(result.RowsAffected)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "hello world")

}
