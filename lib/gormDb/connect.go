package gormDb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DBgorm *gorm.DB

func DbConnect() {
	var err error
	dsn := os.Getenv("DB_URL")
	DBgorm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	fmt.Println("Connected to database")
}
