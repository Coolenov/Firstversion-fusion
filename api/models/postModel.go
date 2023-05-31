package models

type Content struct {
	Title           string
	Description     string
	Link            string
	Source          string
	Image_url       string
	Publishing_time int64
}

type Contents struct {
	Head string
	Body Content
}

type PostTags struct {
	Tag_id  int64
	Post_id int64
}

type Tags struct {
	tag_text string
}
