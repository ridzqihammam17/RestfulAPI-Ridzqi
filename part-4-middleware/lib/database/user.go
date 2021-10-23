package database

import (
	"restfulapi/ridzqi/config"
	"restfulapi/ridzqi/middlewares"
	"restfulapi/ridzqi/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUser(id int) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return user, e
	}

	return user, nil
}

func RegisterUser(user models.User) (models.User, error) {
	if e := config.DB.Save(&user).Error; e != nil {
		return user, e
	}

	return user, nil
}

func UpdateUser(id int, newUser models.User) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return user, e
	}

	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = newUser.Password

	if e := config.DB.Save(&user).Error; e != nil {
		return user, e
	}

	return user, nil
}

func DeleteUser(id int) (models.User, error) {
	var user models.User

	if e := config.DB.First(&user, id).Error; e != nil {
		return user, e
	}

	if e := config.DB.Delete(&user).Error; e != nil {
		return user, e
	}

	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error

	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
