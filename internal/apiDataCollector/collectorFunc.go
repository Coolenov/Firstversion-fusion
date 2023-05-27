package apiDataCollector

import (
	"NaxProject/lib"
	"NaxProject/lib/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAndSaveScrapersPosts(link string) {
	posts := GetScrapedData(link)
	saveScraperPosts(posts)
}

func GetScrapedData(scraper_link string) []lib.Post {
	var posts []lib.Post
	client := &http.Client{}

	req, err := http.NewRequest("GET", scraper_link, nil)
	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return posts
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return posts
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении ответа:", err)
		return posts
	}
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)
		return posts
	}
	resp.Body.Close()
	return posts
}

func saveScraperPosts(posts []lib.Post) {
	db := database.DbConnect()
	defer db.Close()
	for _, post := range posts {
		if !database.CheckPostExistByLink(post.Link, db) {
			postId := database.AddPostIntoPostsTable(post, db)
			for _, tag := range post.Tags {
				var tagId int64
				if !database.CheckTagExist(tag, db) {
					tagId = database.AddTagIntoTagsTable(tag, db)
				} else {
					tagId = database.GetTagIdByTag(tag, db)
				}
				database.AddIntoPostTagsTable(postId, tagId, db)
			}
		}
	}
}
