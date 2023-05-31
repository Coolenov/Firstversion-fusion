package database

import (
	"NaxProject/lib"
	"database/sql"
	"fmt"
)

func GetTagsIdFromDbByTags(tags []string, db *sql.DB) []int64 {
	var tagId int64
	var tagIds []int64
	for _, i := range tags {
		row := db.QueryRow("SELECT id FROM tags WHERE tagText=?", i).Scan(&tagId)
		if row != nil {
			continue
		}
		tagIds = append(tagIds, tagId)
	}
	return tagIds
}

func GetPostIdFromDb(tagIds []int64, db *sql.DB) []int64 {
	var postId int64
	var postIds []int64
	for _, i := range tagIds {
		row := db.QueryRow("SELECT post_id FROM postsTags WHERE tag_id=?", i).Scan(&postId)
		if row != nil {
			continue
		}
		postIds = append(postIds, postId)
	}
	return postIds
}

func GetPostsByTags(PostTagsFromUser []string, db *sql.DB) []lib.Post {
	var result []lib.Post
	tagIds := GetTagsIdFromDbByTags(PostTagsFromUser, db)
	if tagIds != nil {
		postIds := GetPostIdFromDb(tagIds, db)
		var t, d, l, s string //t - title, d - description, l - link
		for _, i := range postIds {
			err := db.QueryRow("SELECT title,description,link,source FROM posts WHERE id=?", i).Scan(&t, &d, &l, &s)
			if err != nil {
				continue
			}
			res := lib.Post{
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

func GetAllPosts(db *sql.DB) []lib.Post {
	var posts []lib.Post
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
		post := lib.Post{
			Title:       t,
			Link:        l,
			Description: d,
			Image_url:   "",
			Tags:        nil,
		}
		posts = append(posts, post)
	}
	return posts
}

func GetPostBySource(sourceName string, db *sql.DB) []lib.Post {
	var posts []lib.Post
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
		post := lib.Post{
			Title:       t,
			Link:        l,
			Description: d,
			Image_url:   "",
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

func GetScrapersUrl(db *sql.DB) ([]string, error) {
	var urls []string
	rows, err := db.Query("SELECT link FROM scrapers WHERE CURRENT_TIMESTAMP > last_request + timeout;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var url string
		err := rows.Scan(&url)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return urls, err
}

func GetTagIdByTag(tag string, db *sql.DB) int64 {
	var tagId int64
	row := db.QueryRow("SELECT id FROM tags WHERE text=?", tag).Scan(&tagId)
	if row != nil {
		fmt.Println(row)
	}
	return tagId

}
