package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Man struct {
	ID    int
	Name  string
	Woman Woman
	Bros  []Bro
}
type Woman struct {
	ManId int
	ID    int
	Name  string
}
type Bro struct {
	ManId int
	Id    int
	Name  string
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	//mapping to db
	db.AutoMigrate(&Man{}, &Woman{}, &Bro{})

	// man := Man{
	// 	ID:   1,
	// 	Name: "ManA",
	// 	Woman: Woman{
	// 		Name: "WifeA",
	// 	},
	// 	Bros: []Bro{
	// 		{Id: 1, Name: "bro1"},
	// 		{Id: 2, Name: "bro1"},
	// 	},
	// }
	tmp := Bro{
		Id:   2,
		Name: "bro2",
	}
	db.Create(&tmp)
	// db.Debug().Create(&man)

	// Create
	// db.Create(&Product{Name: "P2", Price: 200})

	// // Read
	// var product Product
	// db.First(&product, "id = ?", "13")
	// all := []Product{}
	// db.Find(&all, 13)
	// fmt.Printf("all: %+v", all)

	// // Update
	// db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&Product{}).Update("Price", 400)
	// // Update
	// db.Where("price=?", "200").Updates(Product{Price: 300})
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Name": "F42"})

	// // Delete
	// db.Unscoped().Where("name=?", "F42").Delete(&[]Product{})
}
