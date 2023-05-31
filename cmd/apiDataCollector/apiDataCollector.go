package main

import (
	"NaxProject/initialize"
	"NaxProject/internal/apiDataCollector"
)

func init() {
	initialize.LoadEnv()
	initialize.DBConnect()
}

func main() {

	//for {
	//	db := database.DbConnect()
	//	links, err := database.GetScrapersUrl(db)
	//	if err != nil {
	//		db.Close()
	//		fmt.Println("Cant get scrapers URL!!!\n Trying more...", err)
	//		continue
	//	}
	//	fmt.Println(links)
	//	for _, link := range links {
	//		apiDataCollector.GetAndSaveScrapersPosts(link)
	//		err := database.ChangeLastRequestByLink(link, db)
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//	db.Close()
	//	fmt.Println("for is finished")
	//	time.Sleep(5 * time.Second)
	//}

	//for {
	//	links := gormDb.GetScraperLinks()
	//	for _, link := range links {
	//		apiDataCollector.GetAndSavePosts(link)
	//	}
	//	time.Sleep(10 * time.Second)
	//}

	apiDataCollector.GetAndSaveScrapersPosts("http://127.0.0.1:8000/get/news")
}
