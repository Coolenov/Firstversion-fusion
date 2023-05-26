package internal

import (
	"NaxProject/lib"
	"NaxProject/scrapers"
	"database/sql"
)

const (
	habrUrl = "https://habr.com/ru/rss/all/all/"

	dbUrl = "root:password@tcp(127.0.0.1:3306)/Nax"
)

func GetScraperPosts() []lib.Post {
	var habrPosts = scrapers.HabrScraper(habrUrl)

	finalPosts := append(habrPosts)
	return finalPosts
}

func SaveScraperPosts(posts []lib.Post) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	for _, post := range posts {
		if !lib.CheckPostExistByLink(post.Link, db) {
			postId := lib.AddPostIntoPostsTable(post, db)
			for _, tag := range post.Tags {
				var tagId int64
				if !lib.CheckTagExist(tag, db) {
					tagId = lib.AddTagIntoTagsTable(tag, db)
				} else {
					tagId = lib.GetTagId(tag, db)
				}
				lib.AddIntoPostTagsTable(postId, tagId, db)
			}

		}
	}

}
