package models

import (
	"awesome-blog/common/initialize"
	"time"
)

type Tag struct {
	Id          int       `gorm:"id;primaryKey" ` // 自增ID
	Name        string    `gorm:"name" `          // 专题名
	Description string    `gorm:"desc" `          // 专题描述
	CreatedAt   time.Time `gorm:"created_at"`     // 创建时间
}
type TagJson struct {
	Name     string
	Quantity int
	Posts    []PostJson
}

// type Categories []Category
func (Tag) TableName() string {
	return "tag"
}
func (e *Tag) GetTagById(id int) (Tag, error) {
	var tag Tag
	if err := initialize.Db.Table(e.TableName()).Where("id = ?", id).First(&tag).Error; err != nil {
		return tag, err
	}
	return tag, nil
}
func (e *Tag) GetTagByName(tag_name string) (Tag, error) {
	var tag Tag
	if err := initialize.Db.Table(e.TableName()).Where("name = ?", tag_name).First(&tag).Error; err != nil {
		return tag, err
	}
	return tag, nil
}
func (e *Tag) GetTagListByPost(post_id int) ([]Tag, error) {
	var tags []Tag
	if err := initialize.Db.Table(e.TableName()).
		Where("id in ( select tag_id as id from post_tag where post_id = ?)", post_id).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
func (e *Tag) GetAllTag() ([]Tag, error) {
	var tags []Tag
	if err := initialize.Db.Table(e.TableName()).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
