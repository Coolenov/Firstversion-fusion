package lib

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//ts := "root:root@tcp(127.0.0.1:3306)/Naxproject"
//
//func dbConnect(url string) *sql.DB {
//	db, err := sql.Open("mysql", url)
//	if err != nil {
//		panic(err.Error())
//	}
//	return db
//}

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

func GetTagIdFromDb(PostTagsFromUser []string, db *sql.DB) []int64 {
	var tagId int64
	var tagIds []int64
	for _, i := range PostTagsFromUser {
		row := db.QueryRow("SELECT id FROM tags WHERE tagText=?", i).Scan(&tagId)
		if row != nil {
			continue
		}
		tagIds = append(tagIds, tagId)
	}
	return tagIds
}

func getPostIdFromDb(tagIds []int64, db *sql.DB) []int64 {
	var postId int64
	var postIds []int64
	for _, i := range tagIds {
		row := db.QueryRow("SELECT post_id FROM postsTags WHERE tag_id=?", i).Scan(&postId)
		if row != nil {
			continue
		}
		postIds = append(postIds, postId)
		//fmt.Println(postIds[0])
	}
	return postIds
}

func GetPosts(PostTagsFromUser []string, db *sql.DB) []Post {
	var result []Post
	tagIds := GetTagIdFromDb(PostTagsFromUser, db)
	if tagIds != nil {
		postIds := getPostIdFromDb(tagIds, db)

		var t, d, l string //t - title, d - description, l - link
		for _, i := range postIds {
			err := db.QueryRow("SELECT title,description,link FROM posts WHERE id=?", i).Scan(&t, &d, &l)
			if err != nil {
				continue
			}
			res := Post{
				Title:       t,
				Link:        l,
				Description: d,
			}
			result = append(result, res)
		}
		return result
	}
	return result
}
