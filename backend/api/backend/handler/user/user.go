package user

import (
	"awesome-blog/api/backend/handler"
	"awesome-blog/api/backend/service/user"
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService user.UserService
}
type ResponseResult handler.ResponseResult

// UserInfo godoc
// @Summary 根据token获取用户详细信息
// @Tags User
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/user/info [post]
func (c *UserHandler) UserInfo(g *gin.Context) {
	var request UserInfoRequest

	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "获取用户信息成功",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError   // 请求数据有误
		result.Msg = utils.GetFormError(err) // 获取表单错误信息
		g.JSON(http.StatusOK, result)        // 返回 json
		return
	}
	token := request.Token
	fmt.Println("token", token)
	userId, err := global.RDb.Get(global.Ctx, token).Result()
	fmt.Println("userId", userId)
	if err != nil {
		global.Logger.Sugar().Error("Redis error: ", err.Error())
		result.Code = handler.ServerError
		result.Msg = "Redis错误"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	id, _ := strconv.Atoi(userId)
	user, _ := c.userService.GetUserById(id)
	result.Data = user
	g.JSON(
		http.StatusOK, result)
}
