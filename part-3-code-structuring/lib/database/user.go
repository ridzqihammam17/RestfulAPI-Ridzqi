package database

import (
	"restfulapi/ridzqi/config"
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

func CreateUser(user models.User) (models.User, error) {
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
