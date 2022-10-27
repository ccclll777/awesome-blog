package router

import (
	"awesome-blog/api/backend/handler/auth"
	"awesome-blog/api/backend/handler/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
}

func (a *ApiRouter) InitAdminRouter(rootPath string, router *gin.Engine) {
	fmt.Println("正在加载路由")
	authHandler := auth.AuthHandler{}
	authApiRouter := router.Group(rootPath)
	{
		authApiRouter.POST("/auth/login", authHandler.Login)
	}
	userHandler := user.UserHandler{}
	userApiRouter := router.Group(rootPath)
	{
		userApiRouter.POST("/user/info", userHandler.UserInfo)
	}

}
