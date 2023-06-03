package controllers

//func ContentCreate(c *gin.Context) {
//
//	//create post
//	content := models.Content{
//		Title:           "123testTitle",
//		Description:     "test_description",
//		Link:            "git.hub.url",
//		Source:          "hub.url",
//		Image_url:       "asdasd",
//		Publishing_time: 0,
//	}
//	result := initialize.DB.Table("posts").Create(&content)
//	if result.Error != nil {
//		c.Status(400)
//		fmt.Println(result.Error)
//		return
//	}
//	c.JSON(200, gin.H{
//		"content": content,
//	})
//
//}
