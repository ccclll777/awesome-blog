package user

import "awesome-blog/models"

type UserService struct {
	userData models.User
}

// 查询用户
func (c *UserService) GetUserById(userId int) (models.User, error) {
	var user models.User
	user, err := c.userData.GetUserById(userId)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
