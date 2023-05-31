package controllers

import (
	"NaxProject/api/models"
	"NaxProject/initialize"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetAllContents(c *gin.Context) {
	var contents models.Content

	//contents := models.Contents{
	//	Head: ,
	//	Body: models.Content{},
	//}
	//
	result := initialize.DB.Table("posts").Find(&contents)
	if result.Error != nil {
		c.Status(400)
		fmt.Println(result.Error)
		return
	}
	c.JSON(200, gin.H{
		"contents": contents,
	})
}
