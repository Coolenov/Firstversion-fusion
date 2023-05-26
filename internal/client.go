package internal

import (
	"NaxProject/lib"
	"NaxProject/scrapers"
)

const (
	habrUrl = "https://habr.com/ru/rss/all/all/"
)

func GetScraperPosts() []lib.Post {
	var habrPosts = scrapers.HabrScraper(habrUrl)

	finalPosts := append(habrPosts)
	return finalPosts
}

func SaveScraperPosts(posts []lib.Post) {
	db := lib.DbConnect()
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
