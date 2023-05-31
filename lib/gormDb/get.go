package gormDb

import (
	"NaxProject/initialize"
	"NaxProject/lib"
	"fmt"
)

func GetScraperLinks() []string {
	var links []string
	var scrapers []lib.Scraper
	result := initialize.DB.Find(&scrapers)
	if result.Error != nil {
		fmt.Println("Error, cant find scrapers in table Scrapers", result.Error)
	}

	for _, scraper := range scrapers {
		links = append(links, scraper.Link)
	}
	return links

}
