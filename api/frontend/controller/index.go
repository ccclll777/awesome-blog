package controller

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"awesome-blog/models"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

/*
主页所需的数据
*/
var indexPostService service.IndexService

func Index(w http.ResponseWriter, r *http.Request) {
	//要渲染那一页的模版
	indexTemplate := initialize.Template.Index
	if err := r.ParseForm(); err != nil {
		indexTemplate.WriteError(w, err)
	}
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}
	var postPage models.PostPageResult
	postPage = indexPostService.GetPostList(page, initialize.Cfg.Blog.PageSize)
	//posts := indexPostService.GetPostList(page, 2)
	//fmt.Println("posts", posts)
	//search := r.Form.Get("search")
	//category := r.Form.Get("category")
	//tag := r.Form.Get("tag")
	//fmt.Println("search", search)
	//fmt.Println("category", category)
	//fmt.Println("tag", tag)

	fmt.Println("result.Page", postPage.Page)
	fmt.Println("result.Description", postPage.List[0].Description)
	fmt.Println("result.Tags", postPage.List[0].Tags)
	fmt.Println("result.Category", postPage.List[0].Category)
	fmt.Println("result.Title", postPage.List[0].Title)
	fmt.Println("result.MusicId", postPage.List[0].MusicId)
	fmt.Println("result.TotalPage", postPage.TotalPage)
	fmt.Println("result.Total", postPage.Total)
	fmt.Println("result.Date", time.Time(postPage.List[0].Date))
	indexTemplate.WriteData(w, initialize.BuildViewData("Blog", postPage))
}
