package internal

import (
	"NaxProject/lib"
	"NaxProject/scrapers/habr"
)

const (
	habrUrl = "https://habr.com/ru/rss/all/all/"
)

func GetScraperPosts() []lib.Post {
	var habrPosts = habr.HabrScraper()

	finalPosts := append(habrPosts)
	return finalPosts
}
