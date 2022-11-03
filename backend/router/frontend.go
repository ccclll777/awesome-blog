package router

import (
	"awesome-blog/api/frontend/handler"
	"awesome-blog/common/global"
	"net/http"
)

func FrontendRouter() {
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/blog", handler.IndexHandler)
	http.HandleFunc("/categories", handler.CategoryHandler)
	http.HandleFunc("/tags", handler.TagHandler)
	http.HandleFunc("/post", handler.PostHandler)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(global.SystemCfg.CurrentDir+"/public"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(global.SystemCfg.DocumentAssetsDir))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(global.SystemCfg.CurrentDir+"/images"))))

}
