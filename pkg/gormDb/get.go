package gormDb

import (
	"github.com/Coolenov/FusionAPI-monorepo/lib"
)

func GetContents() []pkg.Content {
	var contents []pkg.Content
	DBgorm.Table("posts").Find(&contents)
	return contents
}

func GetTagid(tag_text string) int64 {
	var tag pkg.Tag
	DBgorm.Table("tags").Where("text = ?", tag_text).First(&tag)
	return tag.Id

}

func GetPostids(tagId int64) []int64 {
	var postTags []pkg.PostTags
	DBgorm.Table("posts_tags").Where("tag_id = ?", tagId).Find(&postTags)

	var postIds []int64
	for _, postTag := range postTags {
		postIds = append(postIds, postTag.Post_id)
	}
	return postIds
}

func GetContentsByIds(postIds []int64) []pkg.Content {
	var contents []pkg.Content

	for _, postId := range postIds {
		var content pkg.Content
		DBgorm.Table("posts").Where("id = ?", postId).First(&content)
		contents = append(contents, content)
	}
	return contents
}

func GetContentById(post_id int64) pkg.Content {
	var content pkg.Content
	DBgorm.Table("posts").Where("id = ?", post_id).First(&content)
	return content
}

func GetAllTags() []pkg.Tag {
	var tags []pkg.Tag
	DBgorm.Table("tags").Find(&tags)
	return tags
}

func GetLastContentBySource(source string) pkg.Content {
	var content pkg.Content
	DBgorm.Table("posts").Where("source =?", source).Last(&content)
	return content
}

func GetSourceByid(post_id int64) string {
	var content pkg.Content
	DBgorm.Table("posts").Where("id = ?", post_id).Find(&content)
	return content.Source
}

func GetPreviousContentByid(post_id int64) pkg.Content {
	var content pkg.Content
	source := GetSourceByid(post_id)
	DBgorm.Table("posts").Order("id desc").Where("source = ? AND id < ?", source, post_id).Find(&content)
	return content
}

func GetNextContentByid(post_id int64) pkg.Content {
	var content pkg.Content
	source := GetSourceByid(post_id)
	DBgorm.Table("posts").Where("source = ? AND id > ?", source, post_id).First(&content)
	return content
}

func GetAllSources() []string {
	var sources []string
	DBgorm.Table("posts").Distinct().Pluck("source", &sources)
	return sources
}
