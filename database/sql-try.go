package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func sqlconnect() {

	db, err := sql.Open("mysql", "pk:pk@tcp(9.30.95.8:3306)/test")

	if err != nil {
		panic(err.Error())
	}
	fmt.Println("hello: initiating server")

	defer db.Close()
	//handleRequest()
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO egg VALUES ('from code')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
}
