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
	log.Fatal(http.ListenAndServe(":10000", r))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func ReturnAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllPosts")
	var tags = []string{"Процессоры", "симуляция"}
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Naxproject")
	if err != nil {
		panic(err.Error())
	}
	posts := lib.GetPosts(tags, db)
	json.NewEncoder(w).Encode(posts)
}
