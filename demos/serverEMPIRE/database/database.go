package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Person struct {
	Id   int
	Name string
}

var Db *gorm.DB

func Dbinit() {
	Db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	Db.AutoMigrate(&Person{})

	Db.Create(&Person{Id: 320323, Name: "Terry"})

	var user Person
	Db.First(&user)
	fmt.Println(user)
	//https://gorm.io/docs/
}
