package auth

// 登录请求
type LoginRequest struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}

// 登录响应
type LoginResponse struct {
	Username string `json:"username"`
	UserId   int    `json:"userId"`
	Token    string `json:"token"`
	IsAdmin  int    `json:"isAdmin"`
	Email    string `json:"email"`
}

// 登录请求
type LogoutRequest struct {
	Token string `json:"token"`
}
