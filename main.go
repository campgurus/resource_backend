package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"resource-api/Config"

	"resource-api/Models"
	"resource-api/Routes"
)

var err error

func main() {

	db, err := gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()


	db.AutoMigrate(&Models.Client{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
