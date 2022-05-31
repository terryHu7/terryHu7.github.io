package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	Id   int
	Name string
}

func Dbinit() {
	Db, _ := gorm.Open(mysql.Open("./mydb.db"), &gorm.Config{})
	Db.AutoMigrate(&Person{})

	Db.Create(&Person{Id: 320323, Name: "Terry"})
	//https://gorm.io/docs/
}
