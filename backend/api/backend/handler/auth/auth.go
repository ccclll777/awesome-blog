package auth

import (
	"awesome-blog/api/backend/handler"
	"awesome-blog/api/backend/service/auth"
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type AuthHandler struct {
	authService auth.AuthService
}
type ResponseResult handler.ResponseResult

// Login godoc
// @Summary 根据用户名和密码登录
// @Tags Auth
// @version 1.0
// @Accept json
// @Produce json
// @Param LoginRequest body LoginRequest true "登录表单"
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/auth/login [post]
func (c *AuthHandler) Login(g *gin.Context) {
	var request LoginRequest
	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "登录成功",
		Data: nil,
	}

	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError    // 请求数据有误
		result.Msg = utils.GetFormError(err)  // 获取表单错误信息
		g.JSON(http.StatusBadRequest, result) // 返回 json
		return
	}
	user, _ := c.authService.GetUser(request.Username)
	if user.Username == "" { // 用户不存在
		result.Code = handler.RequestError
		result.Msg = "不存在该用户"
		g.JSON(http.StatusOK, result)
		return
	}
	if request.Password != user.Password { // 密码错误
		result.Code = handler.RequestError
		result.Msg = "密码错误"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	decodedClaims := utils.VerifyToken(user.Token, global.Cfg.Server.TokenSecretKey)

	if decodedClaims == nil {

		jwtClaims := jwt.NewWithClaims(
			jwt.SigningMethodHS256, jwt.MapClaims{
				"userId":   strconv.FormatInt(int64(user.ID), 10),
				"username": user.Username,
				"iat":      time.Now().Unix(),
				"isAdmin":  user.IsAdmin,
			})
		token := utils.GenerateToken(jwtClaims, global.Cfg.Server.TokenSecretKey)
		user.Token = token
		//err := c.authService.UpdateUser(&user)

		//if err != nil {
		//	global.Logger.Sugar().Error("error: ", err.Error())
		//	result.Code = handler.ServerError
		//	result.Msg = "服务器端错误"
		//	g.JSON(http.StatusOK, result) // 返回 json
		//	return
		//}
		//插入redis 用token查询用户的userId 设置过期时间为24h
		err := global.RDb.Set(global.Ctx, user.Token, user.ID, time.Hour*24).Err()
		if err != nil {
			global.Logger.Sugar().Error("Redis error: ", err.Error())
			result.Code = handler.ServerError
			result.Msg = "Redis错误"
			g.JSON(http.StatusOK, result) // 返回 json
			return
		}
	}
	userResponse := LoginResponse{user.Username, user.ID, user.Token, user.IsAdmin, user.Email}
	result.Data = userResponse
	g.JSON(
		http.StatusOK, result)
}

// UserInfo godoc
// @Summary 退出登陆，清空token信息
// @Tags User
// @version 1.0
// @Accept application/json
// @Success 100 object  ResponseResult
// @Failure 103/104 object ResponseResult 失败
// @Router /api/v1/user/logout [post]
func (c *AuthHandler) Logout(g *gin.Context) {
	var request LogoutRequest

	result := ResponseResult{ // 定义 api 返回信息结构
		Code: handler.Success,
		Msg:  "退出登录",
		Data: nil,
	}
	if err := g.ShouldBind(&request); err != nil {
		result.Code = handler.RequestError   // 请求数据有误
		result.Msg = utils.GetFormError(err) // 获取表单错误信息
		g.JSON(http.StatusOK, result)        // 返回 json
		return
	}
	token := request.Token
	_, err := global.RDb.Del(global.Ctx, token).Result()
	if err != nil {
		result.Code = handler.ServerError
		result.Msg = "Redis错误"
		g.JSON(http.StatusOK, result) // 返回 json
		return
	}
	g.JSON(
		http.StatusOK, result)
}
