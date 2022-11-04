package global

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

/*
项目所需的全局变量
*/
var Db *gorm.DB
var RDb *redis.Client
var SystemCfg SystemConfig
var Template HtmlTemplate
var Navigations []Nav
var Cfg Config
var Ctx = context.Background() //goroutine 的上下文

func Init() {
	SystemCfg = LoadSystemConfig()
	LoadConfig()
	//创建mysql链接
	Db = InitMysql()
	RDb = InitRedis()
	Template, _ = InitHtmlTemplate(SystemCfg.CurrentDir + "/templates/frontend")
	Navigations = InitExtraNav()

}
