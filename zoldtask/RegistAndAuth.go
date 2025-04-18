package zoldtask

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"strings"
//)
//
//type RegisterRequest struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//	UserType string `json:"user_type"`
//}
//type LoginRequest struct {
//	Email    string `json:"email"`
//	Password string `json:"password"`
//}
//type AuthResponse struct {
//	Token string `json:"token"`
//}
//
//func RegisterUser(email, password, userType string) (string, error) {
//	if userType != "client" && userType != "moderator" {
//		return "", fmt.Errorf("неверный тип пользователя")
//	}
//	if !strings.Contains(email, "@") {
//		return "", fmt.Errorf("email должен содержать @")
//	}
//	if len(password) < 8 {
//		return "", fmt.Errorf("пароль слишком короткий")
//	}
//
//	requestBody := RegisterRequest{
//		Email:    email,
//		Password: password,
//		UserType: userType,
//	}
//
//	//конвертер в json
//	jsonBody, err := json.Marshal(requestBody)
//	if err != nil {
//		return "", fmt.Errorf("ошибка создания JSON:%v", err)
//	}
//
//	//отправляем post запрос
//	resp, err := http.Post("http://api.avito.ru/register",
//		"application/json",
//		bytes.NewBuffer(jsonBody),
//	)
//	if err != nil {
//		return "", fmt.Errorf("ошибка запроса: %v", err)
//	}
//	defer resp.Body.Close()
//	//проверяем статус
//	switch resp.StatusCode {
//	case http.StatusOK:
//		var response AuthResponse
//		if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
//			return "", fmt.Errorf("ошибка парсинга JSON: %v", err)
//		}
//		return response.Token, nil
//	case http.StatusBadRequest:
//		return "", fmt.Errorf("неверные данные(проверьте email и пароль)")
//	case http.StatusConflict:
//		return "", fmt.Errorf("пользователь уже существует")
//	default:
//		return "", fmt.Errorf("ошибка сервера: %d", resp.StatusCode)
//	}
//}
