package router

import (
	"awesome-blog/api/frontend/controller"
	"awesome-blog/common/initialize"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/blog", controller.Index)
	http.HandleFunc("/categories", controller.Category)
	http.HandleFunc("/tags", controller.Tag)
	http.HandleFunc("/post", controller.Post)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(initialize.SystemCfg.CurrentDir+"/public"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(initialize.SystemCfg.DocumentAssetsDir))))
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir(initialize.SystemCfg.CurrentDir+"/images"))))

}
