package main

import (
	"NaxProject/initialize"
	"NaxProject/lib"
	"log"
)

func init() {
	initialize.LoadEnv()
	initialize.DbConnect()
}

func main() {
	err := initialize.DB.AutoMigrate(&lib.Post{})
	//err = initialize.DB.AutoMigrate(&lib.Tag{})
	//err = initialize.DB.AutoMigrate(lib.Scraper{})
	if err != nil {
		log.Fatal(err)
	}
}
