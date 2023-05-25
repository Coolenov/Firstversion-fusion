package lib

type Post struct {
	Title       string
	Link        string
	Description string
	imageUrl    string `default:"nil"`
	Tags        []string
}

//func (m Module) get() {
//
//}
