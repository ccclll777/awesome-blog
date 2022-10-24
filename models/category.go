package models

import (
	"awesome-blog/common/initialize"
	"fmt"
	"time"
)

/*
博客分类
*/
type Category struct {
	Id          int       `gorm:"id;primaryKey" ` // 自增ID
	Name        string    `gorm:"name" `          // 专题名
	Description string    `gorm:"desc" `          // 专题描述
	CreatedAt   time.Time `gorm:"created_at"`     // 创建时间
}
type CategoryJson struct {
	Name     string
	Quantity int
	Posts    []PostJson
}

//type Categories []Category

func (Category) TableName() string {
	return "category"
}
func (e *Category) GetCategoryById(id int) (Category, error) {
	var category Category
	if err := initialize.Db.Table(e.TableName()).Where("id = ?", id).Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
func (e *Category) GetAllCategory() ([]Category, error) {
	var categories []Category
	if err := initialize.Db.Table(e.TableName()).Find(&categories).Error; err != nil {
		return nil, err
	}
	fmt.Println(categories)
	return categories, nil
}
