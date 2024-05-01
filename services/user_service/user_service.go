package user_service

import (
	"goshort/dtos/user_dto"
	"goshort/models"
	"goshort/utils"
	"goshort/utils/http_status"
	"goshort/utils/jwt"
	"sync"

	"goshort/repositories/user_repository"

	"gopkg.in/go-playground/validator.v9"
)

type UserService struct {
	validate *validator.Validate
}

var (
	instance *UserService
	once     sync.Once
)

func GetInstance() *UserService {
	once.Do(func() {
		instance = &UserService{
			validate: validator.New(),
		}
	})
	return instance
}

// methods
func (us *UserService) CreateUser(user *user_dto.UserDTO_Registration) (*user_dto.UserDTO_Info_Token, *http_status.HTTPStatus) {
	err := us.validate.Struct(user)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	userExists, _ := user_repository.GetInstance().GetUserByEmail(user.Email)

	if userExists != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusConflict, Message: "user already exists"}
	}

	hashedPassword, errHash := utils.HashPassword(user.Password)

	if errHash != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: errHash.Error()}
	}

	userToDb := &models.User{
		Username: user.Username,
		Password: hashedPassword,
		Email:    user.Email,
	}

	new_user, err := user_repository.GetInstance().CreateUser(userToDb)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	new_token, err := jwt.GetInstance().GenerateToken(new_user)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	new_user_with_token := &user_dto.UserDTO_Info_Token{
		Username: new_user.Username,
		Email:    new_user.Email,
		Token:    new_token,
	}

	return new_user_with_token, &http_status.HTTPStatus{Code: http_status.StatusCreated, Message: "user created"}
}

func (us *UserService) GetUserByEmail(email string) (*user_dto.UserDTO_Info_Token, *http_status.HTTPStatus) {
	user, err := user_repository.GetInstance().GetUserByEmail(email)

	if err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: err.Error()}
	}

	if user == nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusNotFound, Message: "user not found"}
	}

	new_token, new_token_err := jwt.GetInstance().GenerateToken(user)

	if new_token_err != nil {
		return nil, &http_status.HTTPStatus{Code: http_status.StatusInternal, Message: new_token_err.Error()}
	}

	userToken := &user_dto.UserDTO_Info_Token{
		Username: user.Username,
		Email:    user.Email,
		Token:    new_token,
	}

	return userToken, &http_status.HTTPStatus{Code: http_status.StatusOK, Message: "user found"}
}
