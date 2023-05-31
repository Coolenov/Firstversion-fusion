package api

import (
	"NaxProject/lib"
	"NaxProject/scrapers/habr"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//func HandleRequests() {
//	r := mux.NewRouter()
//	r.HandleFunc("/get/news", ReturnAllPosts)
//	log.Fatal(http.ListenAndServe(":9980", r))
//}
//func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
//	fmt.Println("Endpoint Hit: returnAllPosts")
//	//var tags = []string{"android"}
//	db := database.DbConnect()
//	defer db.Close()
//	posts := habr.HabrScraper()
//	json.NewEncoder(w).Encode(posts)
//}

func ReturnAllPosts(c *gin.Context) {

	var posts []lib.Post

	posts = habr.HabrScraper()

	c.JSON(200, posts)
}
