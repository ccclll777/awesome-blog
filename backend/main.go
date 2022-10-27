package main

import (
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"awesome-blog/middlewares"
	"awesome-blog/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	global.Init()
	//博客界面的路由
	router.FrontendRouter()
	if err := http.ListenAndServe(":"+global.Cfg.Server.Port, nil); err != nil {
		fmt.Println("ServeErr:", err)
	}
	gin.SetMode(global.Cfg.Server.Mode)
	r := gin.Default()
	r.Use(middlewares.Cors()) //跨域
	router.InitRouter(r)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()
	utils.ShutdownGin(srv, time.Second*5)

}
