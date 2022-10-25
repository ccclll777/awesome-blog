package controller

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/initialize"
	"net/http"
	"strconv"
)

var postService service.PostService

func Post(w http.ResponseWriter, r *http.Request) {
	articleTemplate := initialize.Template.Article

	if err := r.ParseForm(); err != nil {
		articleTemplate.WriteError(w, err)
	}
	post_id := r.Form.Get("id")
	id, _ := strconv.Atoi(post_id)
	postDatail := postService.GetPostById(id)
	articleTemplate.WriteData(w, initialize.BuildViewData("Post", postDatail))
}
