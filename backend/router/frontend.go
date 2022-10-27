package router

import (
	handler2 "awesome-blog/api/frontend/handler"
	"awesome-blog/common/global"
	"net/http"
)

func FrontendRouter() {
	http.HandleFunc("/", handler2.IndexHandler)
	http.HandleFunc("/blog", handler2.IndexHandler)
	http.HandleFunc("/categories", handler2.CategoryHandler)
	http.HandleFunc("/tags", handler2.TagHandler)
	http.HandleFunc("/post", handler2.PostHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(global.SystemCfg.CurrentDir+"/public"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(global.SystemCfg.DocumentAssetsDir))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(global.SystemCfg.CurrentDir+"/images"))))

}
