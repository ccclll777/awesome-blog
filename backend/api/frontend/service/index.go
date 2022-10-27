package service

import (
	"awesome-blog/common/global"
	"awesome-blog/models"
	"sync"
)

/*
主页所需的数据
*/

type IndexService struct {
	postData     models.Post
	categoryData models.Category
	tagData      models.Tag
	userData     models.User
}

func (c *IndexService) GetPost(id int) models.Post {
	var post models.Post
	global.Db.First(&post, id)
	return post
}

func (c *IndexService) GetPostList(page, pageSize int) models.PostPageResult {
	var postList []models.Post
	//var err error
	postList, _ = c.postData.GetPostList(page, pageSize)
	total, _ := c.postData.GetPostCount()
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
	result := models.Pagination(&postJsons, page, global.Cfg.Blog.PageSize, int(total))
	return result
}

func (c *IndexService) GetPostListByTag(tag_name string, page, pageSize int) models.PostPageResult {
	var postList []models.Post
	//var err error
	tag, _ := c.tagData.GetTagByName(tag_name)
	postList, _ = c.postData.GetPostsByTagId(tag.Id, page, pageSize)
	total, _ := c.postData.GetPostCountByTag(tag.Id)
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
	result := models.Pagination(&postJsons, page, global.Cfg.Blog.PageSize, int(total))
	return result
}
func (c *IndexService) GetPostListByCategory(category_name string, page, pageSize int) models.PostPageResult {
	var postList []models.Post
	//var err error
	category, _ := c.categoryData.GetCategoryByName(category_name)
	postList, _ = c.postData.GetPostsByCategoryId(category.Id, page, pageSize)
	total, _ := c.postData.GetPostCountByCategory(category.Id)
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
	result := models.Pagination(&postJsons, page, global.Cfg.Blog.PageSize, int(total))
	return result
}
func (c *IndexService) GetPostJson(post *models.Post) models.PostJson {
	var category models.Category
	category, _ = c.categoryData.GetCategoryById(post.CategoryId)
	tags, _ := c.tagData.GetTagListByPost(post.ID)
	var tagStrs []string
	for _, tag := range tags {
		tagStrs = append(tagStrs, tag.Name)
	}
	var user models.User
	user, _ = c.userData.GetUserById(post.AuthorId)
	return models.PostJson{post.ID, post.Title,
		post.UpdatedAt,
		post.Description,
		tagStrs, user.Username, post.MusicId, category.Name}
}
