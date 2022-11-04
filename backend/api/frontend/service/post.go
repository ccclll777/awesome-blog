package service

import (
	"awesome-blog/models"
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
)

type PostService struct {
	postData     models.Post
	categoryData models.Category
	tagData      models.Tag
}

func (c *PostService) GetPostById(post_id int) models.PostDatail {
	post, err := c.postData.GetPostById(post_id)
	if err != nil {
		fmt.Errorf("获取失败", err)
	}
	tags, _ := c.tagData.GetTagListByPost(post.ID)
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	category, _ := c.categoryData.GetCategoryById(post.CategoryId)
	markdown := []byte(post.MDContent)
	var buf bytes.Buffer
	if err := goldmark.Convert(markdown, &buf); err != nil {
		fmt.Errorf("出错")
	}
	postDatail := models.PostDatail{post, tagNames, category.Name, buf.String()}
	return postDatail
}
