package category

import (
	"awesome-blog/api/backend/handler"
	"awesome-blog/api/backend/service/category"
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	categoryService category.CategoryService
}
type ResponseResult handler.ResponseResult

// CategoryCount godoc
// @Summary 获取全部分类的数量
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/count [get]
func (c *CategoryHandler) CategoryCount(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取分类数量成功",
		Data: nil,
	}
	categoryCount, err := c.categoryService.GetCategoryCount()
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取分类数量失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = categoryCount
	g.JSON(
		http.StatusOK, result)
}

// CategoryList godoc
// @Summary 分页获取分类
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/list [get]
func (c *CategoryHandler) CategoryList(g *gin.Context) {

	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取分类列表成功",
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
	categoryList, err := c.categoryService.GetCategoryList(page, pageSize)

	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取分类列表失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = categoryList
	g.JSON(
		http.StatusOK, result)
}

// AllCategory godoc
// @Summary 获取所有分类
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/all [get]
func (c *CategoryHandler) AllCategory(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取分类列表成功",
		Data: nil}
	categoryList, err := c.categoryService.GetAllCategory()
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取分类列表失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	result.Data = categoryList
	g.JSON(
		http.StatusOK, result)
}

// EditCategory godoc
// @Summary 修改分类信息
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/edit [Post]
func (c *CategoryHandler) EditCategory(g *gin.Context) {
	var request EditCategoryRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "更新分类成功",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError   // 请求数据有误
		result.Msg = utils.GetFormError(err) // 获取表单错误信息
		g.JSON(http.StatusOK, result)        // 返回 json
		return
	}
	category, err := c.categoryService.GetCategoryById(request.Id)
	if err != nil {
		global.Logger.Sugar().Error(err.Error())
		result.Code = handler.ServerError
		result.Msg = "获取分类信息失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	category.Name = request.Name
	category.Description = request.Description
	_, err = c.categoryService.UpdateCategory(category)

	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "更新分类失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	//result.Data = category
	g.JSON(
		http.StatusOK, result)
}

// DeleteCategory godoc
// @Summary 删除某个分类，删除前先判断是否有属于这个分类的文章
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/{id} [delete]
func (c *CategoryHandler) DeleteCategory(g *gin.Context) {
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "删除分类成功",
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
	flag, err := c.categoryService.DeleteCategoryById(id)
	if flag == false || err != nil {
		global.Logger.Sugar().Error(err.Error())
		result.Code = handler.ServerError
		result.Msg = "删除失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// AddCategory godoc
// @Summary 删除某个分类，删除前先判断是否有属于这个分类的文章
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/add [addPost]
func (c *CategoryHandler) AddCategory(g *gin.Context) {
	var request AddCategoryRequest
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
	if request.Name == "" {
		result.Code = handler.RequestError // 请求数据有误
		result.Msg = "分类名不准为空"             // 获取表单错误信息
		g.JSON(http.StatusOK, result)      // 返回 json
		return
	}
	_, err := c.categoryService.AddCategory(request.Name, request.Description)
	if err != nil {
		global.Logger.Sugar().Error("Mysql error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "添加分类失败"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}

// MultiDelCategories godoc
// @Summary 批量删除分类
// @Tags Category
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/category/multiDelete/{ids} [delete]
func (c *CategoryHandler) MultiDelCategories(g *gin.Context) {
	ids := g.Param("ids") // 获取 ids
	if ids == "" {
		g.JSON(http.StatusOK, ResponseResult{
			Code: handler.RequestError,
			Msg:  "请勾选要删除的条目",
			Data: nil,
		})
		return
	}
	if err := c.categoryService.MultiDeleteByIds(ids); err != nil {
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
