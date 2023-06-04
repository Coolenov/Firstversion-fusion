package gormDb

//func AddPostsToDataBase(posts []lib.Post) {
//	db := initialize.DB
//	for _, post := range posts {
//		result := db.FirstOrCreate(&post, lib.Post{Title: post.Title})
//		if result.Error != nil {
//			fmt.Println("Error. Post already exist or cannot add new post", result.Error)
//			return
//		}
//	}
//}
//
//func AddTagsToDataBase(posts []lib.Post) {
//
//	for _, post := range posts {
//		for _, tag := range post.Tags {
//			if err := initialize.DB.Create(&tag).Error; err != nil {
//				if err != nil {
//					fmt.Println(err)
//				}
//				continue
//			}
//		}
//
//	}
//}
//
////func AddTagsToDataBase(posts []lib.Post) {
////	db := initialize.DB
////	for _, post := range posts {
////		for _, tag := range post.Tags {
////			var existingTag lib.Tag
////			if err := db.Where("text = ?", tag.Text).First(&existingTag).Error; err != nil {
////				if gorm.IsRecordNotFoundError(err) {
////					if err := db.Create(&tag).Error; err != nil {
////						continue
////					}
////				} else {
////					continue
////				}
////			}
////		}
////	}
////}
//
//func AddPostWithTags(post lib.Post) error {
//	for i := range post.Tags {
//		var existingTag lib.Tag
//		if err := initialize.DB.Where("text = ?", post.Tags[i].Text).First(&existingTag).Error; err != nil {
//			if errors.Is(err, gorm.ErrRecordNotFound) {
//				if err := initialize.DB.Create(&post.Tags[i]).Error; err != nil {
//					return err
//				}
//			} else {
//				return err
//			}
//		} else {
//			post.Tags[i] = existingTag
//		}
//	}
//
//	if err := initialize.DB.Create(&post).Error; err != nil {
//		return err
//	}
//
//	return nil
//}
