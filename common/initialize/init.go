package initialize

import (
	"gorm.io/gorm"
)

/*
项目所需的全局变量
*/
var Db *gorm.DB
var SystemCfg SystemConfig
var Template HtmlTemplate
var Navigations []Nav
var Cfg Config

func Init() {
	SystemCfg = LoadSystemConfig()
	LoadConfig()
	//创建mysql链接
	Db = InitMysql()
	Template, _ = InitHtmlTemplate(SystemCfg.CurrentDir + "/templates/frontend")
	Navigations = InitExtraNav()
}
