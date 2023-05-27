package main

import (
	"NaxProject/internal/apiDataCollector"
	"NaxProject/lib/database"
	"fmt"
	"time"
)

func main() {

	for {
		db := database.DbConnect()
		links, err := database.GetScrapersUrl(db)
		if err != nil {
			db.Close()
			fmt.Println("Cant get scrapers URL!!!\n Trying more...", err)
			continue
		}
		fmt.Println(links)
		for _, link := range links {
			apiDataCollector.GetAndSaveScrapersPosts(link)
			err := database.ChangeLastRequestByLink(link, db)
			if err != nil {
				fmt.Println(err)
			}
		}
		db.Close()
		fmt.Println("for is finished")
		time.Sleep(5 * time.Second)
	}
}
