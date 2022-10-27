package auth

import "awesome-blog/models"

type AuthService struct {
	userData models.User
}

// 查询用户
func (c *AuthService) GetUser(username string) (models.User, error) {
	var user models.User
	user, err := c.userData.GetUserByUsername(username)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// 更新用户
func (c *AuthService) UpdateUser(user *models.User) error {
	return c.userData.Update(user)
}
