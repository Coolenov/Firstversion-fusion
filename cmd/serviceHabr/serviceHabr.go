package main

import (
	"NaxProject/initialize"
	"NaxProject/scrapers/habr/api"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnv()
}
func main() {

	r := gin.Default()
	r.GET("/get/news", api.ReturnAllPosts)
	r.Run()
}
