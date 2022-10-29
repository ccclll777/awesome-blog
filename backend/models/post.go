package models

import (
	"awesome-blog/common/global"
	"time"
)

// 数据库，博客数据映射模型
type Post struct {
	ID          int       `gorm:"primaryKey"` // ID, store自行控制
	Title       string    `gorm:"title" `
	Description string    `gorm:"description"` //摘要/描述
	Count       int       `gorm:"count"`       // 评论数量
	Content     string    `gorm:"content" `    //markdown 内容
	AuthorId    int       `gorm:"author_id"`   //作者
	IsDraft     bool      `gorm:"is_draft" `   // 是否是草稿
	DeletedAt   time.Time `gorm:"deleted_at"`  // 删除时间
	UpdatedAt   time.Time `gorm:"updated_at"`  // 更新时间
	CreatedAt   time.Time `gorm:"created_at"`  // 创建时间
	IsDelete    int       `gorm:"is_delete"`   // 是否删除
	MusicId     string    `gorm:"music_id"`
	Url         string    `gorm:"url"`
	CategoryId  int       `gorm:"category_id"` //所属分类
	Category    string
}
type PostJson struct {
	Id          int
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Author      string    `json:"author"`
	MusicId     string    `json:"musicId"`
	Category    string
}
type PostDatail struct {
	Post
	Tags     []string
	Category string
	Body     string
}

func (Post) TableName() string {
	return "post"
}
func (e *Post) GetPostById(post_id int) (Post, error) {
	var post Post
	if err := global.Db.Table(e.TableName()).Where("id =?", post_id).Find(&post).Error; err != nil {
		return post, err
	}
	return post, nil
}

func (e *Post) GetPostList(page, pageSize int) ([]Post, error) {
	offset := (page - 1) * pageSize
	queryBuider := global.Db.Limit(pageSize).Offset(offset).Order("created_at asc").Table(e.TableName())
	var postList []Post
	if err := queryBuider.Find(&postList).Error; err != nil {
		return nil, err
	}
	return postList, nil
}
func (e *Post) GetPostCount() (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) GetPostCountByTag(tag_id int) (int, error) {
	var total int64
	if err := global.Db.Table("post_tag").Where("tag_id =?", tag_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) GetPostCountByCategory(category_id int) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("category_id =?", category_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Post) GetPostsByCategoryId(category_id int, page, postQuantity int) ([]Post, error) {
	offset := (page - 1) * postQuantity
	var postList []Post
	if err := global.Db.Table(e.TableName()).Where("category_id =?", category_id).Limit(postQuantity).Offset(offset).Order("created_at asc").Find(&postList).Error; err != nil {
		return nil, err
	}
	return postList, nil
}
func (e *Post) GetPostsByTagId(tag_id int, page, postQuantity int) ([]Post, error) {
	offset := (page - 1) * postQuantity
	var posts []Post
	if err := global.Db.Table(e.TableName()).
		Where("id in ( select post_id as id from post_tag where tag_id = ?)", tag_id).Limit(postQuantity).Offset(offset).Order("created_at asc").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (e *Post) GetPostCountByCategoryId(category_id int) (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Where("category_id =?", category_id).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
