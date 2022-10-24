package controller

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"net/http"
)

var tagService service.TagService

func Tag(w http.ResponseWriter, r *http.Request) {
	tagsTemplate := initialize.Template.Tags
	result := tagService.GetGroupByTag(initialize.Cfg.Blog.TagDisplayQuantity)
	tagsTemplate.WriteData(w, initialize.BuildViewData("Tags", result))
}
