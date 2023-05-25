package lib

import (
	"database/sql"
)

func checkTagExist(tag string, db *sql.DB) int64 {
	var tagId int64
	row := db.QueryRow("SELECT id FROM tags WHERE tagText=?", tag).Scan(&tagId)
	if row != nil {
		return 0
	}
	return tagId
}

func AddTagIntoTagsTable(tag string, db *sql.DB) int64 {
	var tagId int64 = checkTagExist(tag, db)
	if tagId == 0 {
		res, err := db.Exec("INSERT INTO tags(tagText) VALUES(?)", tag)
		if err != nil {
			panic(err.Error())
		}
		tagId, err = res.LastInsertId()
		return tagId
	}
	return tagId
}

func checkPostExistByLink(link string, db *sql.DB) int64 {
	var postId int64
	row := db.QueryRow("SELECT id FROM posts WHERE link=?", link).Scan(&postId)
	if row != nil {
		return 0
	}
	return postId
}

func AddPostIntoPostsTable(post Post, db *sql.DB) int64 {
	postId := checkPostExistByLink(post.Link, db)
	if postId == 0 {
		res, err := db.Exec("INSERT INTO posts(title,description,link,imageUrl) VALUES(?,?,?,?)",
			post.Title,
			post.Description,
			post.Link,
			post.imageUrl)
		if err != nil {
			panic(err.Error())
		}
		postId, err = res.LastInsertId()
		return postId
	}
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
