package middlewares

import (
	"awesome-blog/api/backend/handler"
	"awesome-blog/common/global"
	"awesome-blog/common/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 管理员授权
func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := handler.ResponseResult{ // 封装返回体内容
			Code: handler.Forbidden, // 状态码
			Msg:  "",                // 提示信息
			Data: nil,               // 数据
		}
		token := c.Request.Header.Get("X-Token")
		fmt.Println("token", token)
		if token != "" {
			decodedClaims := utils.VerifyToken(token, global.Cfg.Server.TokenSecretKey)
			//fmt.Println("decodedClaims.IsAdmin", decodedClaims)
			if decodedClaims != nil && decodedClaims.IsAdmin == 1 {
				c.Next()
				c.Abort()
				return
			}
			result.Msg = "当前为管理员界面，无权访问！"
			//fmt.Println(result.Msg)
			c.JSON(http.StatusForbidden, result)
			c.Abort()
			return
		} else {
			result.Msg = "未授权！"
			c.JSON(http.StatusUnauthorized, result)
		}
		c.Abort()
		return
	}
}

// 用户授权
func AuthUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Token")
		result := handler.ResponseResult{ // 封装返回体内容
			Code: handler.Forbidden, // 状态码
			Msg:  "",                // 提示信息
			Data: nil,               // 数据
		}
		if token != "" {
			decodedClaims := utils.VerifyToken(token, global.Cfg.Server.TokenSecretKey)
			if decodedClaims != nil {
				c.Set("userId", decodedClaims.UserId)
				c.Next()
				c.Abort()
				return
			}
			result.Msg = "请求未携带 Token，无权访问！"
			c.JSON(http.StatusForbidden, result)
			c.Abort()
			return
		} else {
			result.Msg = "未授权！"
			c.JSON(http.StatusUnauthorized, result)
		}
		c.Abort()
		return
	}
}
