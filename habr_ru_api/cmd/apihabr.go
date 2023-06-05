package main

import (
	"github.com/Coolenov/FusionAPI-monorepo/initialize"
	"github.com/Coolenov/FusionAPI-monorepo/lib"
	"github.com/Coolenov/FusionAPI-monorepo/services/scrapers_go/habr"
	"github.com/gin-gonic/gin"
)

func ReturnAllPosts(c *gin.Context) {

	var posts []lib.Post

	posts = habr.HabrScraper()
	c.JSON(200, posts)
}

func init() {
	initialize.LoadEnv()
}
func main() {

	r := gin.Default()
	r.GET("/get/news", ReturnAllPosts)
	r.Run(":8010")
}
