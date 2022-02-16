package repository

import (
	"hinccvi/go-template/dao/gorm"
	"hinccvi/go-template/model"
	"hinccvi/go-template/request"
)

func CreateUser(user request.User) (error, int64) {
	userModel := model.User{
		Name:     user.Name,
		Age:      user.Age,
		Position: user.Position,
	}
	result := gorm.DB.Create(&userModel)

	return result.Error, result.RowsAffected
}

func GetAllUser() []model.User {
	var user []model.User
	gorm.DB.Find(&user)
	return user
}

func GetUserById(id int) model.User {
	var user model.User
	gorm.DB.First(&user, id)
	return user
}

func UpdateUser(user model.User, updateUser request.UpdateUser) error {
	if updateUser.Name != "" {
		user.Name = updateUser.Name
	}
	if updateUser.Position != "" {
		user.Position = updateUser.Position
	}
	if updateUser.Age != 0 {
		user.Age = updateUser.Age
	}
	result := gorm.DB.Save(&user)
	return result.Error
}

func DeleteUser(user model.User) error {
	result := gorm.DB.Delete(&user)
	return result.Error
}