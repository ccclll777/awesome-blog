package service

import (
	"awesome-blog/models"
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
)

type PostService struct {
}

func (c *PostService) GetPostById(post_id int) models.PostDatail {
	post, err := postData.GetPostById(post_id)
	if err != nil {
		fmt.Errorf("获取失败", err)
	}
	tags, _ := tagData.GetTagListByPost(post.ID)
	var tagNames []string
	for _, tag := range tags {
		tagNames = append(tagNames, tag.Name)
	}
	category, _ := categoryData.GetCategoryById(post.CategoryId)
	markdown := []byte(post.Content)
	var buf bytes.Buffer
	if err := goldmark.Convert(markdown, &buf); err != nil {
		fmt.Errorf("出错")
	}
	postDatail := models.PostDatail{post, tagNames, category.Name, buf.String()}
	return postDatail
}
