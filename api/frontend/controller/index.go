package controller

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"awesome-blog/models"
	"net/http"
	"strconv"
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
	indexTemplate.WriteData(w, initialize.BuildViewData("Blog", postPage))
}
