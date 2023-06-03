package main

import (
	"NaxProject/api/controllers"
	"NaxProject/initialize"
	"NaxProject/lib/gormDb"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnv()
	gormDb.DbConnect()
}

func main() {
	r := gin.Default()

	r.GET("/contents/all/all", controllers.GetAllContents)
	r.GET("/tags/all/all", controllers.GetAllTags)
	r.GET("/sources", controllers.GetAllSources)

	r.POST("/t", controllers.GetContentsByTag)
	r.POST("/content/id", controllers.GetContentById)
	r.POST("/content/source", controllers.GetLastContentBySource)
	r.POST("/next", controllers.GetNextContent)
	r.POST("/previous", controllers.GetPreviousContent)

	r.Run(":9000")
}
