package habr

import (
	"NaxProject/lib"
	"fmt"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"strings"
	"time"
)

func HabrScraper() []lib.Post {
	var fp = gofeed.NewParser()
	//var feed, _ = fp.ParseURL("https://habr.com/ru/rss/all/all/")
	var feed, _ = fp.ParseURL("https://habr.com/ru/rss/all/all/?limit=100")

	var posts []lib.Post
	for i := 0; i < 100; i++ {
		//var Tags []lib.Tag
		//Tags = makeArrTags(feed.Items[i].Categories)
		item := lib.Post{
			Title:       feed.Items[i].Title,
			Link:        feed.Items[i].Link,
			Description: sanitizeText(feed.Items[i].Description),
			//Tags:            Tags,
			Tags:            feed.Items[i].Categories,
			Source:          "Habr_ru",
			Publishing_time: formPubTime(feed.Items[i].Published),
		}
		posts = append(posts, item)
	}
	return posts
}

func sanitizeText(str string) string {
	p := bluemonday.StripTagsPolicy()
	result := p.Sanitize(str)
	result = strings.ReplaceAll(result, "Читать далее", "")
	return result
}

func formPubTime(timeStr string) int64 {
	timeString := timeStr
	timeLayout := "Mon, 02 Jan 2006 15:04:05 MST"
	timeObj, err := time.Parse(timeLayout, timeString)
	if err != nil {
		fmt.Println("Ошибка при парсинге времени:", err)
	}
	timestamp := timeObj.Unix()
	return timestamp
}

//
//func makeArrTags(tags []string) []lib.Tag {
//	var Tags []lib.Tag
//	for _, tag := range tags {
//		tagObj := lib.Tag{
//			Model: gorm.Model{},
//			Text:  strings.ToLower(tag),
//		}
//		Tags = append(Tags, tagObj)
//	}
//	return Tags
//}
