package post

import "awesome-blog/models"

type PostRequest struct {
	PostId      int    `json:"post_id"`
	Title       string `json:"title" `
	Description string `json:"description"`  //摘要/描述
	Keywords    string `json:"keywords"`     //关键词
	Content     string `json:"content" `     //HTML内容
	MDContent   string `json:"md_content" `  //markdown 内容
	AuthorId    int    `json:"author_id"`    //作者
	IsPublished int    `json:"is_published"` // 是否已经发布 还是只是草稿
	CategoryId  int    `json:"category_id"`  //所属分类
	TagIds      string `json:"tag_ids"`      //标签列表 最多三个
}

type PostResponse struct {
	Post    models.Post
	TagList []models.Tag
}
type UpdatePostRequest struct {
	Id int `json:"post_id" `
}
