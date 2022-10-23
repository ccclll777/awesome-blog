package main

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"awesome-blog/common/router"
	"awesome-blog/models"
	"fmt"
	"net/http"
)

// User 博客账户
var indexPostService service.IndexService
var categoryData models.Category
var tagData models.Tag

func main() {
	initialize.Init()
	router.Router()
	//fmt.Println("results", results)
	//fmt.Println("config.Cfg", config.Cfg.Version)
	//fmt.Printf("Version：v%v \n", config.Cfg.Version)
	//fmt.Printf("ListenAndServe On Port %v \n", config.Cfg.Port)
	//fmt.Printf("UpdateArticle's GitHookUrl: %v   Secret:  %v \n", config.Cfg.GitHookUrl, config.Cfg.WebHookSecret)
	if err := http.ListenAndServe(":"+initialize.Cfg.Server.Post, nil); err != nil {
		fmt.Println("ServeErr:", err)
	}

}
