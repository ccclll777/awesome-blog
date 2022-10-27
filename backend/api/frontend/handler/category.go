package handler

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/global"
	"net/http"
)

var categoryService service.CategoryService

func CategoryHandler(w http.ResponseWriter, r *http.Request) {

	categoriesTemplate := global.Template.Categories
	result := categoryService.GetGroupByCategory(global.Cfg.Blog.CategoryDisplayQuantity)
	categoriesTemplate.WriteData(w, global.BuildViewData("Categories", result))
}
