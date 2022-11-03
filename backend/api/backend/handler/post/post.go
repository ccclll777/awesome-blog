package post

import (
	"awesome-blog/api/backend/handler"
	"awesome-blog/api/backend/service/post"
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"awesome-blog/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PostHandler struct {
	postService post.PostService
}
type ResponseResult handler.ResponseResult

// AddPost godoc
// @Summary 添加文章
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/add [post]
func (c *PostHandler) AddPost(g *gin.Context) {
	var request PostRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "添加文章成功",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}

	_, err := c.postService.AddPost(request.Title, request.Description,
		request.Keywords, request.Content,
		request.MDContent, request.TagIds,
		request.AuthorId, request.IsPublished, request.CategoryId)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "添加文章失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	if request.IsPublished == 0 {
		result.Msg = "添加草稿成功"
	}
	g.JSON(
		http.StatusOK, result)
}

// UpdatePost godoc
// @Summary 更新文章
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/update [post]
func (c *PostHandler) UpdatePost(g *gin.Context) {
	var request PostRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "更新文章成功",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}

	_, err := c.postService.UpdatePost(request.Title, request.Description,
		request.Keywords, request.Content,
		request.MDContent, request.TagIds, request.PostId,
		request.AuthorId, request.IsPublished, request.CategoryId)
	fmt.Println("request.TagIds", request.TagIds)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "更新文章失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// PostCount godoc
// @Summary 获取文章数量
// @Tags Post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/count [get]
func (c *PostHandler) PostCount(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取文章数量成功",
		Data: nil,
	}
	postCount, err := c.postService.GetPostCount()
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取文章数量失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = postCount
	g.JSON(
		http.StatusOK, result)
}

// PostList godoc
// @Summary 分页获取文章
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/list [get]
func (c *PostHandler) PostList(g *gin.Context) {

	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取文章列表成功",
		Data: nil,
	}
	var page int
	var err error
	page, err = strconv.Atoi(g.Query("page"))
	var pageSize int
	pageSize, err = strconv.Atoi(g.Query("pageSize"))
	if err != nil {
		global.Logger.Sugar().Error("string to int conversion failed: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "请求参数有误"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	postList, err := c.postService.GetPostList(page, pageSize)

	var postResponses []PostResponse
	for i := 0; i < len(postList); i++ {
		var tagList []models.Tag
		tagList, _ = c.postService.GetTagByPostId(postList[i].ID)
		postResponses = append(postResponses, PostResponse{Post: postList[i], TagList: tagList})
	}

	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取文章列表失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = postResponses
	g.JSON(
		http.StatusOK, result)
}

// PublishPost godoc
// @Summary 发布文章
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/publish [post]
func (c *PostHandler) PublishPost(g *gin.Context) {
	var request UpdatePostRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "文章已发布",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}
	//postId, err := strconv.Atoi(g.PostForm("post_id"))

	err := c.postService.PublishPost(request.Id)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "发布文章失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// StopPublishPost godoc
// @Summary 撤销发布文章
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/stopPublish [post]
func (c *PostHandler) StopPublishPost(g *gin.Context) {
	var request UpdatePostRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "文章已撤销",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}
	err := c.postService.StopPublishPost(request.Id)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "撤销文章失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// DeletePost godoc
// @Summary 删除文章
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/delete [post]
func (c *PostHandler) DeletePost(g *gin.Context) {
	var request UpdatePostRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "删除文章",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}
	err := c.postService.DeletePost(request.Id)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "删除文章失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// UnDeletePost godoc
// @Summary 撤销删除文章
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/undelete [post]
func (c *PostHandler) UnDeletePost(g *gin.Context) {
	var request UpdatePostRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "撤销删除成功",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}
	err := c.postService.UnDeletePost(request.Id)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "撤销失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// PostById godoc
// @Summary 根据id获取文章信息
// @Tags post
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/post/fetchPostById [post]
func (c *PostHandler) PostById(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取文章信息成功",
		Data: nil,
	}
	postId, err := strconv.Atoi(g.Query("post_id"))
	post, err := c.postService.GetPostById(postId)
	fmt.Println("post", post)
	var postResponse PostResponse
	var tagList []models.Tag
	tagList, _ = c.postService.GetTagByPostId(post.ID)
	postResponse = PostResponse{Post: post, TagList: tagList}
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取文章信息失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = postResponse
	g.JSON(
		http.StatusOK, result)
}
