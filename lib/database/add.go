package database

import (
	"FusionAPI/lib"
	"database/sql"
)

func AddTagIntoTagsTable(tag string, db *sql.DB) int64 {
	res, err := db.Exec("INSERT INTO tags(text) VALUES(?)", tag)
	if err != nil {
		panic(err.Error())
	}
	tagId, err := res.LastInsertId()
	return tagId

}

func AddPostIntoPostsTable(post lib.Post, db *sql.DB) int64 {
	res, err := db.Exec("INSERT INTO posts(source,title,description,link,image_url,publishing_time) VALUES(?,?,?,?,?,?)",
		post.Source,
		post.Title,
		post.Description,
		post.Link,
		post.ImageUrl,
		post.PublishingTime)
	if err != nil {
		panic(err.Error())
	}
	postId, err := res.LastInsertId()
	return postId

}

func AddIntoPostTagsTable(postId int64, tagId int64, db *sql.DB) {
	_, err := db.Exec("INSERT INTO posts_tags(tag_id,post_id) VALUES(?,?)",
		tagId,
		postId)
	if err != nil {
		panic(err.Error())
	}
}
