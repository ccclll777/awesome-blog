package handler

import (
	"awesome-blog/api/frontend/service"
	"awesome-blog/common/global"
	"net/http"
	"strconv"
)

var postService service.PostService

func PostHandler(w http.ResponseWriter, r *http.Request) {
	articleTemplate := global.Template.Article

	if err := r.ParseForm(); err != nil {
		articleTemplate.WriteError(w, err)
	}
	post_id := r.Form.Get("id")
	id, _ := strconv.Atoi(post_id)
	postDatail := postService.GetPostById(id)
	articleTemplate.WriteData(w, global.BuildViewData("Post", postDatail))
}
