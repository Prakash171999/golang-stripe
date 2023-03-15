package main

import (
	"fmt"
	"proj-mido/stripe-gateway/Config"
	"proj-mido/stripe-gateway/Models"
	"proj-mido/stripe-gateway/Routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Products{})
	r := Routes.SetupRouter()
	//running
	r.Run()
}
