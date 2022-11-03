package router

import (
	"awesome-blog/common/global"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter(r *gin.Engine) {

	err := global.InitLogger(
		global.Cfg.LoggerCfg.FileName,
		global.Cfg.LoggerCfg.Level,
		global.Cfg.LoggerCfg.Format,
		global.Cfg.LoggerCfg.MaxSize,
		global.Cfg.LoggerCfg.MaxBackups,
		global.Cfg.LoggerCfg.MaxAge,
	)
	if err != nil {
		log.Panicln("初始化日志失败：", err.Error())
	}
	//流量控制中间件
	//middlewares.InitBucket(time.Second*time.Duration(global.Cfg.Server.LimitTime), global.Cfg.Server.LimitCap)

	// 加载路由
	apiRouter := ApiRouter{}
	apiRouter.InitAdminRouter("/api/v1", r)
}
