package api

import (
	"NaxProject/lib/database"
	"NaxProject/scrapers/habr"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/get/news", ReturnAllPosts)
	log.Fatal(http.ListenAndServe(":9980", r))
}
func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPosts")
	//var tags = []string{"android"}
	db := database.DbConnect()
	defer db.Close()
	posts := habr.HabrScraper()
	json.NewEncoder(w).Encode(posts)
}
