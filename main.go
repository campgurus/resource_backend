package main
import (
	"fmt"
	"github.com/jinzhu/gorm"
	"resource-api/Config"
	"resource-api/Models"
	"resource-api/Routes"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Client{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
