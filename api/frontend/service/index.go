package service

import (
	"awesome-blog/common/initialize"
	"awesome-blog/models"
	"fmt"
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
	fmt.Printf("post=%+v\n", postList[0])
	total := postData.GetPostCount()
	fmt.Printf("total", total)
	var postJsons []models.PostJson
	//查询分类信息和标签信息
	wg := sync.WaitGroup{}
	for _, post := range postList {
		wg.Add(1)
		go func() {
			//var category models.Category
			//category, _ = categoryData.GetCategoryById(post.CategoryId)
			//tags, _ := tagData.GetTagListByPost(post.ID)
			//var tagStrs []string
			//for _, tag := range tags {
			//	tagStrs = append(tagStrs, tag.Name)
			//}
			//fmt.Println(tagStrs)
			//var user models.User
			//user, _ = userData.GetUserById(post.AuthorId)
			var postJson models.PostJson
			postJson = c.GetPostJson(post)
			postJsons = append(postJsons, postJson)
			wg.Done()
		}()
	}
	wg.Wait()
	//fmt.Println(postJsons)
	result := models.Pagination(&postJsons, page, initialize.Cfg.Blog.PageSize, int(total))
	return result
}
func (c *IndexService) GetPostJson(post models.Post) models.PostJson {
	var category models.Category
	category, _ = categoryData.GetCategoryById(post.CategoryId)
	tags, _ := tagData.GetTagListByPost(post.ID)
	var tagStrs []string
	for _, tag := range tags {
		tagStrs = append(tagStrs, tag.Name)
	}
	fmt.Println(tagStrs)
	var user models.User
	user, _ = userData.GetUserById(post.AuthorId)
	return models.PostJson{post.Title,
		post.UpdatedAt,
		post.Description,
		tagStrs, user.Username, post.MusicId, "", "", category.Name}
}
