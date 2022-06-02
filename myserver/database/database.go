package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	Id   int
	Name string
}

func Dbinit() {
	Db, _ := gorm.Open(mysql.New((mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      nil,
	})), &gorm.Config{})
	Db.AutoMigrate(&Person{})

	Db.Create(&Person{Id: 320323, Name: "Terry"})

	var user Person
	Db.First(&user)
	fmt.Println(user)
	//https://gorm.io/docs/
}
