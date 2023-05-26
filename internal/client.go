package internal

import (
	"NaxProject/lib"
	"NaxProject/scrapers"
	"database/sql"
)

const (
	habrUrl = "https://habr.com/ru/rss/all/all/"

	dbUrl = "root:root@tcp(127.0.0.1:3306)/Naxproject"
)

func GetScraperPosts() []lib.Post {
	var habrPosts = scrapers.HabrScraper(habrUrl)

	finalPosts := append(habrPosts)
	return finalPosts
}

func SavePosts(posts []lib.Post) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	for _, post := range posts {
		postId := lib.AddPostIntoPostsTable(post, db)
		if postId != 0 {
			for _, tag := range post.Tags {
				tagId := lib.AddTagIntoTagsTable(tag, db)
				lib.AddIntoPostTagsTable(postId, tagId, db)
			}
		}
	}

}
