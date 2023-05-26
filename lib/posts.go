package lib

type Post struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	ImageUrl    string   `default:"nil"`
	Source      string   `json:"source"`
	Tags        []string `json:"tags"`
}

//func (m Module) get() {
//
//}
