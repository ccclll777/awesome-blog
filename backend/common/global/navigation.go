package global

/*
初始化导航栏
*/
type Nav struct {
	Title string
	Path  string
}

func InitExtraNav() []Nav {
	var navigations []Nav
	navigations = append(navigations, Nav{"About", Cfg.Navigation.About})
	navigations = append(navigations, Nav{"Work", Cfg.Navigation.Work})
	return navigations
}
