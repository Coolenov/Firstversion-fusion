package lib

type Post struct {
	Title          string   `json:"title"`
	Link           string   `json:"link"`
	Description    string   `json:"description"`
	ImageUrl       string   `json:"imageUrl"`
	Source         string   `json:"source"`
	Tags           []string `json:"tags"`
	PublishingTime int64    `json:"publishingTime"`
}

type Contents struct {
	data []Content `json:"data"`
}

type Content struct {
	sourceName string `json:"sourceName"`
	postData   []Post `json:"postdata"`
}

type Response struct {
	Articles []Post `json:"articles"`
}
