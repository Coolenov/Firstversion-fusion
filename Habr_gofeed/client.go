package main

import (
	"Habr_gofeed/lib"
	"Habr_gofeed/scrapers"
	"database/sql"
)

const (
	habrUrl = "https://habr.com/ru/rss/all/all/"
)

func dbConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/Nax")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getPosts() []lib.Post {
	var habrPosts = scrapers.HabrScraper(habrUrl)

	finalPosts := append(habrPosts)
	return finalPosts
}

func savePosts(posts []lib.Post) {
	db := dbConnect()
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
