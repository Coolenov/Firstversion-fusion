package pkg

type Post struct {
	Title          string   `json:"title"`
	Link           string   `json:"link"`
	Description    string   `json:"description"`
	ImageUrl       string   `json:"imageUrl"`
	Source         string   `json:"source"`
	Tags           []string `json:"tags"`
	PublishingTime int64    `json:"publishingTime"`
}

type Content struct {
	Id              int64
	Title           string
	Link            string
	Description     string
	Image_url       string
	Source          string
	Publishing_time int64
}

type PostTags struct {
	Id      int64
	Tag_id  int64
	Post_id int64
}

type Tag struct {
	Id   int64
	Text string
}

type Sourse struct {
	Name string
}
