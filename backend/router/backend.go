package router

import (
	"awesome-blog/api/backend/handler/auth"
	"awesome-blog/api/backend/handler/category"
	"awesome-blog/api/backend/handler/post"
	"awesome-blog/api/backend/handler/tag"
	"awesome-blog/api/backend/handler/user"
	"awesome-blog/middlewares"
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
		authApiRouter.POST("/auth/logout", authHandler.Logout)
	}
	userHandler := user.UserHandler{}
	userApiRouter := router.Group(rootPath)
	{
		userApiRouter.POST("/user/info", userHandler.UserInfo)
	}
	categoryHandler := category.CategoryHandler{}
	categoryApiRouter := router.Group(rootPath)
	{
		categoryApiRouter.GET("/category/count", middlewares.AuthAdminMiddleware(), categoryHandler.CategoryCount)
		categoryApiRouter.GET("/category/list", middlewares.AuthAdminMiddleware(), categoryHandler.CategoryList)
		categoryApiRouter.GET("/category/all", middlewares.AuthAdminMiddleware(), categoryHandler.AllCategory)
		categoryApiRouter.POST("/category/edit", middlewares.AuthAdminMiddleware(), categoryHandler.EditCategory)
		categoryApiRouter.DELETE("/category/:id", middlewares.AuthAdminMiddleware(), categoryHandler.DeleteCategory)
		categoryApiRouter.POST("/category/add", middlewares.AuthAdminMiddleware(), categoryHandler.AddCategory)
		categoryApiRouter.DELETE("/category/multiDelete/:ids", middlewares.AuthAdminMiddleware(), categoryHandler.MultiDelCategories)
	}
	tagHandler := tag.TagHandler{}
	tagApiRouter := router.Group(rootPath)
	{
		tagApiRouter.GET("/tag/count", middlewares.AuthAdminMiddleware(), tagHandler.TagCount)
		tagApiRouter.GET("/tag/list", middlewares.AuthAdminMiddleware(), tagHandler.TagList)
		tagApiRouter.GET("/tag/all", middlewares.AuthAdminMiddleware(), tagHandler.AllTag)
		//tagApiRouter.POST("/category/edit", middlewares.AuthAdminMiddleware(), categoryHandler.EditCategory)
		tagApiRouter.DELETE("/tag/:id", middlewares.AuthAdminMiddleware(), tagHandler.DeleteTag)
		tagApiRouter.POST("/tag/add", middlewares.AuthAdminMiddleware(), tagHandler.AddTag)
		tagApiRouter.DELETE("/tag/multiDelete/:ids", middlewares.AuthAdminMiddleware(), tagHandler.MultiDelTags)
		//tagApiRouter.GET("/tag/postTag", middlewares.AuthAdminMiddleware(), tagHandler.PostTagList)

	}

	postHandler := post.PostHandler{}
	postApiRouter := router.Group(rootPath)
	{
		postApiRouter.GET("/post/count", middlewares.AuthAdminMiddleware(), postHandler.PostCount)
		postApiRouter.GET("/post/list", middlewares.AuthAdminMiddleware(), postHandler.PostList)
		postApiRouter.GET("/post/fetchPostById", middlewares.AuthAdminMiddleware(), postHandler.PostById)
		postApiRouter.POST("/post/add", middlewares.AuthAdminMiddleware(), postHandler.AddPost)
		postApiRouter.POST("/post/update", middlewares.AuthAdminMiddleware(), postHandler.UpdatePost)
		postApiRouter.POST("/post/publish", middlewares.AuthAdminMiddleware(), postHandler.PublishPost)
		postApiRouter.POST("/post/stopPublish", middlewares.AuthAdminMiddleware(), postHandler.StopPublishPost)
		postApiRouter.POST("/post/delete", middlewares.AuthAdminMiddleware(), postHandler.DeletePost)
		postApiRouter.POST("/post/undelete", middlewares.AuthAdminMiddleware(), postHandler.UnDeletePost)

	}
}
