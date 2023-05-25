package scrapers

import (
	"Habr_gofeed/lib"
	"github.com/mmcdole/gofeed"
)

func HabrScraper(url string) []lib.Post {
	var fp = gofeed.NewParser()
	var feed, _ = fp.ParseURL(url)

	var posts []lib.Post
	for i := 0; i < 20; i++ {
		item := lib.Post{
			Title:       feed.Items[i].Title,
			Link:        feed.Items[i].Link,
			Description: feed.Items[i].Description,
			Tags:        feed.Items[i].Categories,
		}
		posts = append(posts, item)
	}
	return posts
}
