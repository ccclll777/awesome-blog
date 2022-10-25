package main

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"awesome-blog/models"
	"awesome-blog/router"
	"fmt"
	"net/http"
)

// User 博客账户
// var indexPostService service.IndexService
var categoryService service.CategoryService

// var categoryData models.Category
// var tagData models.Tag
var postData models.Post

func main() {
	initialize.Init()
	router.Router()
	//post, _ := postData.GetPostById(1)
	//fmt.Println(post)
	//var categories []models.Categories
	//categories, _ := categoryData.GetAllCategory()
	//fmt.Println(categories)
	//result := categoryService.GetCategoryList(initialize.Cfg.Blog.CategoryDisplayQuantity)
	//fmt.Println("result", result)
	//fmt.Println("config.Cfg", config.Cfg.Version)
	//fmt.Printf("Version：v%v \n", config.Cfg.Version)
	//fmt.Printf("ListenAndServe On Port %v \n", config.Cfg.Port)
	//fmt.Printf("UpdateArticle's GitHookUrl: %v   Secret:  %v \n", config.Cfg.GitHookUrl, config.Cfg.WebHookSecret)
	if err := http.ListenAndServe(":"+initialize.Cfg.Server.Post, nil); err != nil {
		fmt.Println("ServeErr:", err)
	}

}
