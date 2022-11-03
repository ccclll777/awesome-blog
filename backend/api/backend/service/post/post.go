package post

import (
	"awesome-blog/models"
)

type PostService struct {
	postData models.Post
	tagData  models.Tag
}

func (e *PostService) AddPost(title, description, keywords, content, mdContent, TagIds string, author_id, is_published, category_id int) (models.Post, error) {
	post := models.Post{Title: title, Description: description, Keywords: keywords,
		Content: content, MDContent: mdContent, AuthorId: author_id,
		IsPublished: is_published, CategoryId: category_id}
	//count,err := e.postData.GetPostCountByTitle(title)
	//if count != 0{
	//
	//}
	err := e.postData.AddPost(post, TagIds)
	return post, err
}
func (e *PostService) UpdatePost(title, description, keywords, content, mdContent, TagIds string, post_id, author_id, is_published, category_id int) (models.Post, error) {
	post := models.Post{ID: post_id, Title: title, Description: description, Keywords: keywords,
		Content: content, MDContent: mdContent, AuthorId: author_id,
		IsPublished: is_published, CategoryId: category_id}
	err := e.postData.UpdatePost(post, TagIds)
	return post, err
}
func (e *PostService) GetPostCount() (int, error) {
	total, err := e.postData.GetPostCount()
	return total, err
}
func (e *PostService) GetPostList(page, pageSize int) ([]models.Post, error) {
	postList, err := e.postData.GetPostList(page, pageSize)

	return postList, err
}
func (e *PostService) GetPostById(postId int) (models.Post, error) {
	post, err := e.postData.GetPostById(postId)
	return post, err
}

func (e *PostService) GetTagByPostId(postId int) ([]models.Tag, error) {
	tagList, err := e.tagData.GetTagListByPost(postId)
	return tagList, err
}
func (e *PostService) PublishPost(postId int) error {
	err := e.postData.PublishPost(postId)
	return err
}
func (e *PostService) StopPublishPost(postId int) error {
	err := e.postData.StopPublishPost(postId)
	return err
}

func (e *PostService) DeletePost(postId int) error {
	err := e.postData.DeletePost(postId)
	return err
}
func (e *PostService) UnDeletePost(postId int) error {
	err := e.postData.UnDeletePost(postId)
	return err
}
