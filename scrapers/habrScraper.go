package scrapers

import (
	"NaxProject/lib"
	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"strings"
)

func HabrScraper(url string) []lib.Post {
	var fp = gofeed.NewParser()
	var feed, _ = fp.ParseURL(url)

	var posts []lib.Post
	for i := 0; i < 20; i++ {
		item := lib.Post{
			Title:       feed.Items[i].Title,
			Link:        feed.Items[i].Link,
			Description: cleanText(feed.Items[i].Description),
			Tags:        feed.Items[i].Categories,
			Source:      "Habr_ru",
		}
		posts = append(posts, item)
	}
	return posts
}

func cleanText(str string) string {
	p := bluemonday.StripTagsPolicy()
	result := p.Sanitize(str)
	result = strings.ReplaceAll(result, "Читать далее", "")
	return result
}
