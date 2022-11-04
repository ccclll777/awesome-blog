package handler

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/global"
	"awesome-blog/models"
	"net/http"
	"strconv"
)

/*
主页所需的数据
*/
/*
主页所需的数据
*/
var indexPostService service.IndexService

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//要渲染那一页的模版
	indexTemplate := global.Template.Index
	if err := r.ParseForm(); err != nil {
		indexTemplate.WriteError(w, err)
	}
	page, err := strconv.Atoi(r.Form.Get("page"))
	if err != nil {
		page = 1
	}
	//search := r.Form.Get("search")
	category := r.Form.Get("category")
	tag := r.Form.Get("tag")
	//fmt.Println("search", search)
	//fmt.Println("category", category)
	//fmt.Println("tag", tag)
	var postPage models.PostPageResult
	if category != "" {
		postPage = indexPostService.GetPostListByCategory(category, page, global.Cfg.Blog.PageSize)
	} else if tag != "" {
		postPage = indexPostService.GetPostListByTag(tag, page, global.Cfg.Blog.PageSize)
	} else {
		postPage = indexPostService.GetPostList(page, global.Cfg.Blog.PageSize)
	}

	indexTemplate.WriteData(w, global.BuildViewData("Blog", postPage))
}
