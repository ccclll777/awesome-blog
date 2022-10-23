package initialize

import (
	"fmt"
	"github.com/spf13/viper"
)

// 使用viper读取配置信息
func LoadConfig() {
	path := SystemCfg.CurrentDir
	fmt.Println("path", path)
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("读取配置文件失败: %w \n", err))
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(fmt.Errorf("无法将配置文件解析为结构体%w \n", err))
	}
}
