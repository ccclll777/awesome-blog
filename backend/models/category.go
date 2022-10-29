package models

import (
	"awesome-blog/common/global"
	"time"
)

/*
博客分类
*/
type Category struct {
	Id          int       `gorm:"id;primaryKey" ` // 自增ID
	Name        string    `gorm:"name" `          // 专题名
	Description string    `gorm:"description" `   // 专题描述
	CreatedAt   time.Time `gorm:"created_at"`     // 创建时间
	IsDelete    int       `gorm:"is_delete"`      // 是否被删除
	//ParentId    int       `gorm:"parent_id"`      // 父节点的ID
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
	if err := global.Db.Table(e.TableName()).Where("id = ?", id).Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
func (e *Category) GetCategoryByName(category_name string) (Category, error) {
	var category Category
	if err := global.Db.Table(e.TableName()).Where("name = ?", category_name).Find(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
func (e *Category) GetAllCategory() ([]Category, error) {
	var categories []Category
	if err := global.Db.Table(e.TableName()).Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
func (e *Category) GetCategoryCount() (int, error) {
	var total int64
	if err := global.Db.Table(e.TableName()).Count(&total).Error; err != nil {
		return 0, err
	}
	return int(total), nil
}
func (e *Category) GetCategoryList(page, pageSize int) ([]Category, error) {
	offset := (page - 1) * pageSize
	queryBuider := global.Db.Limit(pageSize).Offset(offset).Order("id asc").Table(e.TableName())
	var categoryList []Category
	if err := queryBuider.Find(&categoryList).Error; err != nil {
		return nil, err
	}
	return categoryList, nil
}
func (e *Category) DeleteCategoryById(id int) (Category, error) {
	var category Category
	if err := global.Db.Table(e.TableName()).Where("id = ?", id).Delete(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
func (e *Category) UpdateCategory(category Category) (Category, error) {
	if err := global.Db.Table(e.TableName()).Where("id = ?", category.Id).Updates(Category{Name: category.Name, Description: category.Description}).Error; err != nil {
		return category, err
	}
	return category, nil
}

func (e *Category) AddCategory(category Category) (Category, error) {
	if err := global.Db.Table(e.TableName()).Create(&category).Error; err != nil {
		return category, err
	}
	return category, nil
}
