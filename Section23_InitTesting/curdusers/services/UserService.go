package services

import (
	"curdusers/configs"
	"curdusers/models"
)

type IUserService interface {
	CreateUser(*models.User) error
}

type UserRepository struct {
	Func IUserService
}

var userRepository IUserService

func init() {
	ur := &UserRepository{}
	ur.Func = ur

	userRepository = ur
}

func GetUserRepository() IUserService {
	return userRepository
}

func SetUserRepository(ur IUserService) {
	userRepository = ur
}

func (u *UserRepository) CreateUser(user *models.User) error {
	err := configs.DBMysql.Save(&user)
	if err != nil {
		return err.Error
	}

	return nil
}
