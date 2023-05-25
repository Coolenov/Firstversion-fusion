package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	counter := 0

	for counter != 100 {
		timer := time.NewTimer(10 * time.Second)

		posts := getPosts()
		savePosts(posts)
		fmt.Println("Parsed ", counter)
		if counter != 99 {
			<-timer.C
			counter++
		}

	}
}
