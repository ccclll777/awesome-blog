package handler

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/global"
	"net/http"
)

var tagService service.TagService

func TagHandler(w http.ResponseWriter, r *http.Request) {
	tagsTemplate := global.Template.Tags
	result := tagService.GetGroupByTag(global.Cfg.Blog.TagDisplayQuantity)
	tagsTemplate.WriteData(w, global.BuildViewData("Tags", result))
}
