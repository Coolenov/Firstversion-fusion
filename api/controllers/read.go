package controllers

import (
<<<<<<< Updated upstream
	"NaxProject/lib"
	"NaxProject/lib/gormDb"
=======
>>>>>>> Stashed changes
	"fmt"
	"github.com/Coolenov/FusionAPI-monorepo/lib"
	"github.com/Coolenov/FusionAPI-monorepo/lib/gormDb"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllContents(c *gin.Context) {
	var contents []lib.Content
	contents = gormDb.GetContents()
	c.JSON(200, gin.H{
		"contents": contents,
	})
}

func GetContentsByTag(c *gin.Context) {
	tag := c.PostForm("tag")
	tag_id := gormDb.GetTagid(tag)
	post_ids := gormDb.GetPostids(tag_id)
	contents := gormDb.GetContentsByIds(post_ids)
	c.JSON(200, contents)
}

func GetContentById(c *gin.Context) {
	id_str := c.PostForm("id")
	post_id, err := strconv.ParseInt(id_str, 10, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	content := gormDb.GetContentById(post_id)
	c.JSON(200, content)
}

func GetAllTags(c *gin.Context) {
	tags := gormDb.GetAllTags()

	c.JSON(200, tags)
}

func GetLastContentBySource(c *gin.Context) {
	source := c.PostForm("source")
	content := gormDb.GetLastContentBySource(source)
	if content.Id != 0 {
		c.JSON(200, gin.H{
			"error":   false,
			"content": content,
		})
	} else {
		c.JSON(400, gin.H{"error": true, "message": "Post not found"})
	}
}

func GetNextContent(c *gin.Context) {
	post_id := c.PostForm("id")
	post_id_int, err := strconv.ParseInt(post_id, 10, 64)
	content := gormDb.GetNextContentByid(post_id_int)
	if err != nil {
		fmt.Println(err)
	}
	if content.Id != 0 {
		c.JSON(200, gin.H{
			"error":   false,
			"content": content,
		})
	} else {
		c.JSON(400, gin.H{"error": true, "message": "Post not found"})
	}
}

func GetPreviousContent(c *gin.Context) {
	post_id := c.PostForm("id")
	post_id_int, err := strconv.ParseInt(post_id, 10, 64)
	content := gormDb.GetPreviousContentByid(post_id_int)
	if err != nil {
		fmt.Println(err)
	}
	if content.Id != 0 {
		c.JSON(200, gin.H{
			"error":   false,
			"content": content,
		})
	} else {
		c.JSON(400, gin.H{"error": true, "message": "Post not found"})
	}
}

func GetAllSources(c *gin.Context) {
	sources := gormDb.GetAllSources()
	c.JSON(200, gin.H{
		"error":   false,
		"content": sources,
	})
}
