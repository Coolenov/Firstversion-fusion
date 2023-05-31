package lib

type Post struct {
	Title           string   `json:"title"`
	Link            string   `json:"link"`
	Description     string   `json:"description"`
	Image_url       string   `json:"imageUrl"`
	Source          string   `json:"source"`
	Tags            []string `json:"tags"`
	Publishing_time int64    `json:"publishingTime"`
}

//type Post struct {
//	gorm.Model
//	Title           string
//	Link            string
//	Description     string
//	Image_url       string
//	Source          string
//	Tags            []Tag `gorm:"many2many:post_tags;"`
//	Publishing_time int64
//}
//
//type Tag struct {
//	gorm.Model
//	Text string
//}
//
//type Scraper struct {
//	gorm.Model
//	Name string
//	Link string
//}

//type Response struct {
//	Articles []Post `json:"articles"`
//}
