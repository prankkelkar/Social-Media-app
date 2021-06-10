package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetCon() *gorm.DB {
	dsn := "pk:pk@tcp(9.30.95.8:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}
