package tag

import (
	"awesome-blog/api/backend/handler"
	"awesome-blog/api/backend/service/tag"
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TagHandler struct {
	tagService tag.TagService
}
type ResponseResult handler.ResponseResult

// TagCount godoc
// @Summary 获取全部标签的数量
// @Tags Tag
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/tag/count [get]
func (c *TagHandler) TagCount(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取标签数量成功",
		Data: nil,
	}
	tagCount, err := c.tagService.GetTagCount()
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取标签数量失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = tagCount
	g.JSON(
		http.StatusOK, result)
}

// TagList godoc
// @Summary 分页获取标签
// @Tags Tag
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/tag/list [get]
func (c *TagHandler) TagList(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取标签列表成功",
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
	tagList, err := c.tagService.GetTagList(page, pageSize)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取分类列表失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = tagList
	g.JSON(
		http.StatusOK, result)
}

// DeleteTag godoc
// @Summary 删除某个分类，删除前先判断是否有属于这个分类的文章
// @Tags Tag
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/tag/{id} [delete]
func (c *TagHandler) DeleteTag(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "删除标签成功",
		Data: nil,
	}
	id, err := strconv.Atoi(g.Param("id"))
	if err != nil {
		global.Logger.Sugar().Error("string to int conversion failed: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "请求参数有误"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	err = c.tagService.DeleteTagById(id)
	if err != nil {
		global.Logger.Sugar().Error(err.Error())
		result.Code = handler.ServerError
		result.Msg = "删除失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// AddTag godoc
// @Summary 删除某个分类，删除前先判断是否有属于这个分类的文章
// @Tags Tag
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/tag/add [post]
func (c *TagHandler) AddTag(g *gin.Context) {
	var request AddTagRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "添加分类成功",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}
	_, err := c.tagService.AddTag(request.Name, request.Description)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "添加标签失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// MultiDelTags godoc
// @Summary 批量删除分类
// @Tags Tag
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/tag/multiDelete/{ids} [delete]
func (c *TagHandler) MultiDelTags(g *gin.Context) {
	ids := g.Param("ids") // 获取 ids
	if ids == "" {
		g.JSON(http.StatusOK, ResponseResult{
			Code: handler.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}
	if err := c.tagService.MultiDeleteByIds(ids); err != nil {
		global.Logger.Sugar().Error("error: ", err.Error())
		g.JSON(http.StatusOK, ResponseResult{
			Code: handler.ServerError,
			Msg:  "服务器端错误",
			Data: nil,
		})
		return
	}
	g.JSON(http.StatusOK, ResponseResult{
		Code: handler.Success,
		Msg:  "删除成功",
		Data: nil,
	})
}
