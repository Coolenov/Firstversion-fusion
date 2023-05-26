package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/Naxproject")
	//if err != nil {
	//	panic(err.Error())
	//}
	//var tags = []string{"Процессоры", "симуляция"}
	//
	//fmt.Println(posts[0].Link) //пустой массив, ошибка в логике getPosts()!!!
	//defer db.Close()

	//counter := 0
	//for counter != 100 {
	//	timer := time.NewTimer(10 * time.Second)
	//
	//	posts := internal.GetPosts()
	//	internal.SavePosts(posts)
	//	fmt.Println("Parsed ", counter)
	//	if counter != 99 {
	//		<-timer.C
	//		counter++
	//	}
	//
	//}
}
