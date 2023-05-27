package main

import (
	"NaxProject/internal"
	"NaxProject/lib/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.DbConnect()
	defer db.Close()
	posts := internal.GetScraperPosts()
	internal.SaveScraperPosts(posts)
}
