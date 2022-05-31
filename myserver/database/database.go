package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Dbinit() {
	Db, err = gorm.Open(sqlite.Open("./mydb.db"), &gorm.Config{})
	//https://gorm.io/docs/
}
