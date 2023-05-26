package lib

type Post struct {
	Title       string   `json:"title"`
	Link        string   `json:"link"`
	Description string   `json:"description"`
	imageUrl    string   `default:"nil"`
	Tags        []string `json:"tags"`
}

//func (m Module) get() {
//
//}
