package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	//https://blog.csdn.net/fwhezfwhez/article/details/120372792
	cfg := cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	cfg.AllowAllOrigins = true
	return cors.New(cfg)
	//return func(c *gin.Context) {
	//	method := c.Request.Method
	//
	//	// 可将将* 替换为指定的域名
	//	c.Header("Access-Control-Allow-Origin", "*")
	//	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
	//	c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	//	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
	//	c.Header("Access-Control-Allow-Credentials", "true")
	//	fmt.Println("cors")
	//	if method == "OPTIONS" {
	//		c.AbortWithStatus(http.StatusNoContent)
	//	}
	//	defer func() {
	//		if err := recover(); err != nil {
	//			log.Printf("Panic info is: %v", err)
	//		}
	//	}()
	//
	//	c.Next()
	//}
}
