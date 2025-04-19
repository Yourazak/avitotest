package users

import (
	"avitotes/config"
	payload "avitotes/interal/dto/payloadError"
	"avitotes/pkg/jwt"
	"fmt"
	"log"
)

type UserService struct {
	UserRepository *UserRepo
	Config         *config.Config
}

func NewUserService(repo *UserRepo, conf *config.Config) *UserService {
	return &UserService{
		UserRepository: repo,
		Config:         conf,
	}
}

// Функция регистрации пользователя
func (service *UserService) Register(email, password, role string) (*User, error) {
	//реализовать валидацию email
	//если пользователь с таким email уже есть, то ошибка
	user := NewUser(email, password, role)
	createdUser, err := service.UserRepository.Creat(user)
	if err != nil {
		log.Println("CreateUser:не удалось создать пользователя123")
		return nil, err
	}
	return createdUser, nil
}

// функция авторизации пользователя
func (service *UserService) Login(email, password string) (*payload.TokenResponse, error) {
	user, err := service.UserRepository.FindUserByEmailPass(email, password)
	if err != nil {
		log.Println("Login", err)
		return nil, err
	}
	tokenStr, err := service.GetToken(user.Role)
	if err != nil {
		log.Println("Login", err)
		return nil, err
	}
	req := &payload.TokenResponse{
		Token: tokenStr,
	}
	return req, nil
}

// достать токен соответствующей роли
func (service *UserService) GetToken(role string) (string, error) {
	if role != "client" && role != "moderator" {
		return "", nil
	}
	var secret string
	switch role {
	case "client":
		secret = service.Config.Auth.AuthTokenClient
	case "moderator":
		secret = service.Config.Auth.AuthTokenModerator
	}
	jwtToken := jwt.NewJWT(secret)
	tokenStr, err := jwtToken.Create(role)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenStr, nil
}
