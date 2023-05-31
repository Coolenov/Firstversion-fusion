package main

import (
	"NaxProject/api/controllers"
	"NaxProject/initialize"
	"github.com/gin-gonic/gin"
)

func init() {
	initialize.LoadEnv()
	initialize.DbConnect()
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.ShowHomePage)
	r.GET("/all", controllers.GetAllContents)
	r.Run() // listen and serve
}
