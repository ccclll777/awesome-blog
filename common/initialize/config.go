package initialize

// 组合全部配置模型

type Config struct {
	Server     Server     `mapstructure:"server"`
	Mysql      Mysql      `mapstructure:"mysql"`
	Blog       Blog       `mapstructure:"blog"`
	Navigation Navigation `mapstructure:"navigations"`
}

// 和根目录config.yaml映射
type Server struct {
	Post string `mapstructure:"post"`
}
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

// 博客配置
type Blog struct {
	PageSize                int      `mapstructure:"pageSize"`
	DescriptionLen          int      `mapstructure:"descriptionLen"`
	Author                  string   `mapstructure:"author"`
	AppRepository           string   `mapstructure:"AppRepository"`
	AppName                 string   `mapstructure:"AppName"`
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
