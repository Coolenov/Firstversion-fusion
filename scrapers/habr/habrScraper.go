package habr

import (
	"NaxProject/lib"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"strings"
	"time"
)

func HabrScraper() []lib.Post {
	var fp = gofeed.NewParser()
	var feed, _ = fp.ParseURL("https://habr.com/ru/rss/all/all/")

	var posts []lib.Post
	for i := 0; i < 20; i++ {
		item := lib.Post{
			Title:          feed.Items[i].Title,
			Link:           feed.Items[i].Link,
			Description:    sanitizeText(feed.Items[i].Description),
			Tags:           feed.Items[i].Categories,
			Source:         "Habr_ru",
			PublishingTime: formPubTime(feed.Items[i].Published),
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
	timeLayout := "02 Jan 2006 15:04:05 MST"
	timeObj, err := time.Parse(timeLayout, timeString)

	if err != nil {
		return 0
	}
	timestamp := timeObj.Unix()
	return timestamp
}
