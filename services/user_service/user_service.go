package user_service

import (
	"errors"
	user_dtos "goshort/dtos/user_dto"
	"goshort/models"
	"goshort/utils"

	"goshort/repositories/user_repository"

	"gopkg.in/go-playground/validator.v9"
)

type UserService struct {
	validate *validator.Validate
}

var instance *UserService

func GetInstance() *UserService {
	if instance == nil {
		instance = &UserService{
			validate: validator.New(),
		}
	}
	return instance
}

// methods
func (us *UserService) CreateUser(user *user_dtos.UserDTO_Registration) (*user_dtos.UserDTO_Info, error) {
	errValidateUser := us.validate.Struct(user)

	if errValidateUser != nil {
		return nil, errValidateUser
	}

	userExists := user_repository.GetInstance().GetUserByEmail(user.Email)

	if userExists != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, errHash := utils.HashPassword(user.Password)

	if errHash != nil {
		return nil, errors.New("error hashing password")
	}

	userToDb := &models.User{
		Username: user.Username,
		Password: hashedPassword,
		Email:    user.Email,
	}

	newUser, err := user_repository.GetInstance().CreateUser(userToDb)

	if err != nil {
		return nil, err
	}

	return newUser, nil
}
