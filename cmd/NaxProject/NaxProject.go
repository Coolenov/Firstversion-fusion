package main

import (
	"NaxProject/internal"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/Nax")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	counter := 0
	for counter != 100 {
		timer := time.NewTimer(10 * time.Second)
		posts := internal.GetScraperPosts()
		internal.SaveScraperPosts(posts)
		fmt.Println("Parsed ", counter)
		if counter != 99 {
			<-timer.C
			counter++
		}

	}
}
