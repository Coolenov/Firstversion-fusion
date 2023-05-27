package database

import (
	"NaxProject/lib"
	"database/sql"
	"fmt"
)

func AddTagIntoTagsTable(tag string, db *sql.DB) int64 {
	res, err := db.Exec("INSERT INTO tags(tagText) VALUES(?)", tag)
	if err != nil {
		panic(err.Error())
	}
	tagId, err := res.LastInsertId()
	return tagId

}

func AddPostIntoPostsTable(post lib.Post, db *sql.DB) int64 {
	res, err := db.Exec("INSERT INTO posts(title,description,link,imageUrl,source) VALUES(?,?,?,?,?)",
		post.Title,
		post.Description,
		post.Link,
		post.ImageUrl,
		post.Source)
	if err != nil {
		panic(err.Error())
	}
	postId, err := res.LastInsertId()
	fmt.Print()
	return postId

}

func AddIntoPostTagsTable(postId int64, tagId int64, db *sql.DB) {
	_, err := db.Exec("INSERT INTO postsTags(tag_id,post_id) VALUES(?,?)",
		tagId,
		postId)
	if err != nil {
		panic(err.Error())
	}
}
