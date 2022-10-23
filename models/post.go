package models

import (
	"awesome-blog/common/initialize"
	"fmt"
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
	Title       string    `json:"title"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Tags        []string  `json:"tags"`
	Author      string    `json:"author"`
	MusicId     string    `json:"musicId"`
	Path        string
	ShortUrl    string
	Category    string
}

func (Post) TableName() string {
	return "post"
}
func (e *Post) GetPostList(page, pageSize int) []Post {
	offset := (page - 1) * pageSize
	queryBuider := initialize.Db.Limit(pageSize).Offset(offset).Order("updated_at asc").Table(e.TableName())
	var postList []Post
	err := queryBuider.Find(&postList)
	if err.Error != nil {
		fmt.Errorf("查询博客列表失败: ", err)
		return nil
	}
	return postList
}
func (e *Post) GetPostCount() int {
	var total int64
	initialize.Db.Table(e.TableName()).Count(&total)
	return int(total)
}
