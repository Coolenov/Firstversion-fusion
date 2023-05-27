package database

import "database/sql"

func CheckTagExist(tag string, db *sql.DB) bool {
	var tagId int64
	row := db.QueryRow("SELECT id FROM tags WHERE tagText=?", tag).Scan(&tagId)
	if row != nil {
		return false
	}
	return true
}

func CheckPostExistByLink(link string, db *sql.DB) bool {
	var postId int64
	row := db.QueryRow("SELECT id FROM posts WHERE link=?", link).Scan(&postId)
	if row != nil {
		return false
	}
	return true
}
