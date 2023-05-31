package controllers

import "github.com/gin-gonic/gin"

func ShowHomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"Hello": "World",
	})
}
