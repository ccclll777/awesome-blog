package category

import (
	"awesome-blog/models"
	"strconv"
	"strings"
)

type CategoryService struct {
	categoryData models.Category
	postData     models.Post
}

func (e *CategoryService) GetCategoryCount() (int, error) {
	total, err := e.categoryData.GetCategoryCount()
	return total, err
}
func (e *CategoryService) GetCategoryList(page, pageSize int) ([]models.Category, error) {
	categoryList, err := e.categoryData.GetCategoryList(page, pageSize)
	return categoryList, err
}
func (e *CategoryService) GetAllCategory() ([]models.Category, error) {
	categoryList, err := e.categoryData.GetAllCategory()
	return categoryList, err
}

// 删除分类，删除分类之前需要判断这一类有没有文章
func (e *CategoryService) DeleteCategoryById(id int) (bool, error) {
	count, _ := e.postData.GetPostCountByCategoryId(id)
	if count == 0 {
		_, err := e.categoryData.DeleteCategoryById(id)
		return true, err
	}
	return false, nil
}
func (e *CategoryService) GetCategoryById(id int) (models.Category, error) {
	category, err := e.categoryData.GetCategoryById(id)
	return category, err
}
func (e *CategoryService) UpdateCategory(category models.Category) (models.Category, error) {
	category, err := e.categoryData.UpdateCategory(category)
	return category, err
}
func (e *CategoryService) AddCategory(name string, description string) (models.Category, error) {
	category := models.Category{Name: name, Description: description}
	category, err := e.categoryData.AddCategory(category)
	return category, err
}

func (e *CategoryService) MultiDeleteByIds(ids string) error {
	idList := strings.Split(ids, ",") // 根据 , 分割成字符串数组
	for i := 0; i < len(idList); i++ {
		id, _ := strconv.Atoi(idList[i])
		flag, _ := e.DeleteCategoryById(id)
		if flag == false {
			continue
		}
	}
	return nil
}
