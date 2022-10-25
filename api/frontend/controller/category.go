package controller

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"net/http"
)

var categoryService service.CategoryService

func Category(w http.ResponseWriter, r *http.Request) {

	categoriesTemplate := initialize.Template.Categories
	result := categoryService.GetGroupByCategory(initialize.Cfg.Blog.CategoryDisplayQuantity)
	categoriesTemplate.WriteData(w, initialize.BuildViewData("Categories", result))
}
