package service

import (
	"awesome-blog/common/initialize"
	"awesome-blog/models"
	"sync"
)

/*
主页所需的数据
*/
var postData models.Post
var categoryData models.Category
var tagData models.Tag
var userData models.User

type IndexService struct {
}

func (c *IndexService) GetPost(id int) models.Post {
	var post models.Post
	initialize.Db.First(&post, id)
	return post
}

func (c *IndexService) GetPostList(page, pageSize int) models.PostPageResult {
	var postList []models.Post
	//var err error
	postList, _ = postData.GetPostList(page, pageSize)
	total, _ := postData.GetPostCount()
	var postJsons []models.PostJson
	//查询分类信息和标签信息
	wg := sync.WaitGroup{}
	for i := 0; i < len(postList); i++ {
		post := postList[i]
		wg.Add(1)
		go func() {
			var postJson models.PostJson
			postJson = c.GetPostJson(&post)
			postJsons = append(postJsons, postJson)
			wg.Done()
		}()
	}
	wg.Wait()
	result := models.Pagination(&postJsons, page, initialize.Cfg.Blog.PageSize, int(total))
	return result
}

func (c *IndexService) GetPostListByTag(tag_name string, page, pageSize int) models.PostPageResult {
	var postList []models.Post
	//var err error
	tag, _ := tagData.GetTagByName(tag_name)
	postList, _ = postData.GetPostsByTagId(tag.Id, page, pageSize)
	total, _ := postData.GetPostCountByTag(tag.Id)
	var postJsons []models.PostJson
	//查询分类信息和标签信息
	wg := sync.WaitGroup{}
	for i := 0; i < len(postList); i++ {
		post := postList[i]
		wg.Add(1)
		go func() {
			var postJson models.PostJson
			postJson = c.GetPostJson(&post)
			postJsons = append(postJsons, postJson)
			wg.Done()
		}()
	}
	wg.Wait()
	result := models.Pagination(&postJsons, page, initialize.Cfg.Blog.PageSize, int(total))
	return result
}
func (c *IndexService) GetPostListByCategory(category_name string, page, pageSize int) models.PostPageResult {
	var postList []models.Post
	//var err error
	category, _ := categoryData.GetCategoryByName(category_name)
	postList, _ = postData.GetPostsByCategoryId(category.Id, page, pageSize)
	total, _ := postData.GetPostCountByCategory(category.Id)
	var postJsons []models.PostJson
	//查询分类信息和标签信息
	wg := sync.WaitGroup{}
	for i := 0; i < len(postList); i++ {
		post := postList[i]
		wg.Add(1)
		go func() {
			var postJson models.PostJson
			postJson = c.GetPostJson(&post)
			postJsons = append(postJsons, postJson)
			wg.Done()
		}()
	}
	wg.Wait()
	result := models.Pagination(&postJsons, page, initialize.Cfg.Blog.PageSize, int(total))
	return result
}
func (c *IndexService) GetPostJson(post *models.Post) models.PostJson {
	var category models.Category
	category, _ = categoryData.GetCategoryById(post.CategoryId)
	tags, _ := tagData.GetTagListByPost(post.ID)
	var tagStrs []string
	for _, tag := range tags {
		tagStrs = append(tagStrs, tag.Name)
	}
	var user models.User
	user, _ = userData.GetUserById(post.AuthorId)
	return models.PostJson{post.ID, post.Title,
		post.UpdatedAt,
		post.Description,
		tagStrs, user.Username, post.MusicId, category.Name}
}
