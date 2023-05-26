package lib

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const dbUrl = "root:root@tcp(127.0.0.1:3306)/NaxProject"

func DbConnect() *sql.DB {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func CheckTagExist(tag string, db *sql.DB) bool {
	var tagId int64
	row := db.QueryRow("SELECT id FROM tags WHERE tagText=?", tag).Scan(&tagId)
	if row != nil {
		return false
	}
	return true
}

func GetTagId(tag string, db *sql.DB) int64 {
	var tagId int64
	row := db.QueryRow("SELECT id FROM tags WHERE tagText=?", tag).Scan(&tagId)
	if row != nil {
		return 0
	}
	return tagId
}

func AddTagIntoTagsTable(tag string, db *sql.DB) int64 {
	res, err := db.Exec("INSERT INTO tags(tagText) VALUES(?)", tag)
	if err != nil {
		panic(err.Error())
	}
	tagId, err := res.LastInsertId()
	return tagId

}

func CheckPostExistByLink(link string, db *sql.DB) bool {
	var postId int64
	row := db.QueryRow("SELECT id FROM posts WHERE link=?", link).Scan(&postId)
	if row != nil {
		return false
	}
	return true
}

func AddPostIntoPostsTable(post Post, db *sql.DB) int64 {
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

func GetPostsByTags(PostTagsFromUser []string, db *sql.DB) []Post {
	var result []Post
	tagIds := GetTagIdFromDb(PostTagsFromUser, db)
	if tagIds != nil {
		postIds := getPostIdFromDb(tagIds, db)
		var t, d, l, s string //t - title, d - description, l - link
		for _, i := range postIds {
			err := db.QueryRow("SELECT title,description,link,source FROM posts WHERE id=?", i).Scan(&t, &d, &l, &s)
			if err != nil {
				continue
			}
			res := Post{
				Title:       t,
				Link:        l,
				Description: d,
				Source:      s,
			}
			result = append(result, res)
		}
		return result
	}
	return result
}
func GetAllPosts(db *sql.DB) []Post {
	var posts []Post
	rows, err := db.Query("SELECT title,description,link FROM posts")
	if err != nil {
		fmt.Println("cannot take data from table(posts)", err)
	}
	//defer rows.Close()
	for rows.Next() {
		var t, d, l string //t - title, d - description, l - link
		err := rows.Scan(&t, &d, &l)
		if err != nil {
			fmt.Println("Ошибка чтения строки")
			continue
		}
		post := Post{
			Title:       t,
			Link:        l,
			Description: d,
			ImageUrl:    "",
			Tags:        nil,
		}
		posts = append(posts, post)
	}
	return posts
}

func GetPostBySource(sourceName string, db *sql.DB) []Post {
	var posts []Post
	rows, err := db.Query("SELECT title,description,link FROM posts WHERE source=?", sourceName)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var t, d, l string //t - title, d - description, l - link
		err := rows.Scan(&t, &d, &l)
		if err != nil {
			fmt.Println(err)
			continue
		}
		post := Post{
			Title:       t,
			Link:        l,
			Description: d,
			ImageUrl:    "",
			Tags:        nil,
			Source:      sourceName,
		}
		posts = append(posts, post)
	}
	return posts
}

func GetUniqSources(db *sql.DB) []string {
	var result []string

	rows, err := db.Query("SELECT DISTINCT source FROM posts")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var source string
		if err := rows.Scan(&source); err != nil {
			fmt.Println(err)
		}
		result = append(result, source)
	}
	return result
}
