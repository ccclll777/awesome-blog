package global

import (
	"os"
)

// 系统配置
type SystemConfig struct {
	AppName             string
	Version             float32
	CurrentDir          string
	GitHookUrl          string
	AppRepository       string
	DocumentDir         string
	DocumentAssetsDir   string
	DocumentContentDir  string
	DocumentExtraNavDir string
}

func LoadSystemConfig() SystemConfig {
	var cfg SystemConfig
	var err error
	cfg.CurrentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg.AppName = "AwesomeBlog"
	cfg.Version = 1.0
	return cfg
}
