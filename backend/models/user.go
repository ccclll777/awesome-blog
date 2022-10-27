// Package model provides ...
package models

import (
	"awesome-blog/common/global"
	"fmt"
)

// User 博客账户
type User struct {
	ID        int    `gorm:"id;primaryKey"`
	Username  string `gorm:"username"`   // 用户名
	Password  string `gorm:"password"`   // 密码
	Email     string `gorm:"email"`      // 邮件地址
	PhoneN    string `gorm:"phone_n" `   // 手机号
	Address   string `gorm:"address" `   // 地址信息
	LoginIP   string `gorm:"login_ip"`   // 最近登录IP
	CreatedAt string `gorm:"created_at"` // 创建时间
	IsAdmin   int    `gorm:"is_admin"`   // 创建时间
	Token     string `gorm:"token"`      // 创建时间
}

func (User) TableName() string {
	return "user"
}
func (e *User) GetUserById(id int) (User, error) {
	var user User
	if err := global.Db.Table(e.TableName()).Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// GetByUsername 根据用户名
func (e *User) GetUserByUsername(username string) (User, error) {
	var user User
	fmt.Println("username", user)
	if err := global.Db.Table(e.TableName()).Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// 根据token查询用户
func (e *User) GetUserByToken(username string) (User, error) {
	var user User
	fmt.Println("username", user)
	if err := global.Db.Table(e.TableName()).Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// 更新用户
func (e *User) Update(user *User) error {
	return global.Db.Table(e.TableName()).Save(&user).Error
}
