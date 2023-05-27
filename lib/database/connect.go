package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const dbUrl = "root:root@tcp(127.0.0.1:3306)/NaxProject"

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	return db
}
