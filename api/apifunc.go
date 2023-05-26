package api

import (
	"NaxProject/lib"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", homePage)
	r.HandleFunc("/posts", ReturnAllPosts)
	r.HandleFunc("/contents", ReturnContents)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPosts")
	//var tags = []string{"android"}
	db := lib.DbConnect()
	defer db.Close()
	posts := lib.GetAllPosts(db)
	json.NewEncoder(w).Encode(posts)
}

func MakeContentBySourceName(sourceName string, db *sql.DB) Content {
	posts := lib.GetPostBySource(sourceName, db)
	result := Content{
		sourceName: sourceName,
		postData:   posts,
	}
	//fmt.Println(result.postData)
	//fmt.Println(result.sourceName)
	fmt.Println(result)
	return result
}

func ReturnContents(w http.ResponseWriter, r *http.Request) {
	db := lib.DbConnect()
	defer db.Close()
	var contents []map[string]interface{}
	uniqSources := lib.GetUniqSources(db)
	for _, oneSource := range uniqSources {
		posts := lib.GetPostBySource(oneSource, db)
		myMap := make(map[string]interface{})
		myMap["data"] = posts
		myMap["service_name"] = oneSource

		contents = append(contents, myMap)
	}
	//fmt.Println(contents)
	json.NewEncoder(w).Encode(contents)
}
