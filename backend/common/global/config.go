package global

// 组合全部配置模型

type Config struct {
	Server     Server     `mapstructure:"server"`
	Mysql      Mysql      `mapstructure:"mysql"`
	Redis      Redis      `mapstructure:"redis"`
	Blog       Blog       `mapstructure:"blog"`
	Navigation Navigation `mapstructure:"navigations"`
	LoggerCfg  LoggerCfg  `yaml:"logger"`
}

// 和根目录config.yaml映射
type Server struct {
	Mode            string   `yaml:"mode"`              // 运行模式
	Port            string   `yaml:"port"`              // 运行端口
	BackendPort     string   `yaml:"backend_port"`      // admin接口 运行端口
	TokenExpireTime int      `yaml:"token_expire_time"` // JWT token 过期时间
	TokenSecretKey  string   `yaml:"token_secret_key"`
	AllowedRefers   []string `yaml:"allowed_refers"` // 允许的 referer
	LimitTime       int64    `yaml:"limit_time"`     // 限流时间间隔
	LimitCap        int64    `yaml:"limit_cap"`      // 间隔时间内最大访问次数

}
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}
type Redis struct {
	Url      string `mapstructure:"url"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"DB"`
}

// 博客配置
type Blog struct {
	PageSize                int      `mapstructure:"pageSize"`
	DescriptionLen          int      `mapstructure:"descriptionLen"`
	Author                  string   `mapstructure:"author"`
	AppRepository           string   `mapstructure:"AppRepository"`
	AppName                 string   `mapstructure:"AppName"`
	UtterancesRepo          string   `mapstructure:"utterancesRepo"`
	CategoryDisplayQuantity int      `mapstructure:"categoryDisplayQuantity"`
	TagDisplayQuantity      int      `mapstructure:"tagDisplayQuantity"`
	TimeLayout              string   `mapstructure:"timeLayout"`
	SiteName                string   `mapstructure:"siteName"`
	HtmlKeywords            string   `mapstructure:"htmlKeywords"`
	HtmlDescription         string   `mapstructure:"htmlDescription"`
	ThemeColor              string   `mapstructure:"themeColor"`
	ThemeOption             []string `mapstructure:"themeOption"`
}

// 导航栏配置
type Navigation struct {
	About string `mapstructure:"about"`
	Work  string `mapstructure:"work"`
}

// logger 日志
type LoggerCfg struct {
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
	Level      string `yaml:"level"`
	Format     string `yaml:"format"`
}
