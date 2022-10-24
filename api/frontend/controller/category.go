package controller

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"fmt"
	"net/http"
)

var categoryService service.CategoryService

func Category(w http.ResponseWriter, r *http.Request) {

	categoriesTemplate := initialize.Template.Categories
	//fmt.Println("Category")
	result := categoryService.GetGroupByCategory(initialize.Cfg.Blog.CategoryDisplayQuantity)
	fmt.Println(result[0].Posts)
	categoriesTemplate.WriteData(w, initialize.BuildViewData("Categories", result))
}
