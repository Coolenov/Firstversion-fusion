package initialize

import (
	"database/sql"
	"fmt"
	"os"
)

// var DB *gorm.DB
var DB *sql.DB

func DBConnect() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to database")
}

//func DbConnect() {
//	var err error
//	dsn := os.Getenv("DB_URL")
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatal("Failed to connect to database")
//	}
//	fmt.Println("Connected to database")
//}
