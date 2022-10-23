// Package model provides ...
package models

import "awesome-blog/common/initialize"

// User 博客账户
type User struct {
	ID        int    `gorm:"id;primaryKey"`
	Username  string `gorm:"username"`   // 用户名
	Password  string `gorm:"password"`   // 密码
	Email     string `gorm:"email"`      // 邮件地址
	PhoneN    string `gorm:"phone_n" `   // 手机号
	Address   string `gorm:"address" `   // 地址信息
	LogoutAt  string `gorm:"logout_at" ` // 登出时间
	LoginIP   string `gorm:"login_ip"`   // 最近登录IP
	LoginUA   string `gorm:"login_ua" `  // 最近登录地区
	LoginAt   string `gorm:"login_at"`   // 最近登录时间
	CreatedAt string `gorm:"created_at"` // 创建时间
}

func (User) TableName() string {
	return "user"
}
func (e *User) GetUserById(id int) (User, error) {
	var user User
	if err := initialize.Db.Table(e.TableName()).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
