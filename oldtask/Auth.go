package oldtask

//
//import (
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"net/url"
//)
//
//type authResponse struct {
//	Token string `json:"token"`
//}
//
//func main() {
//	token, err := GetAuthToken("client") //пытаемся вытянуть токен
//	if err != nil {                      // если ошибка
//		fmt.Println("Ошибка:", err) //выводим ошибку
//		return                      //возвращаем
//	}
//	fmt.Println("Токен:", token) //если все ноомально , то выводим токен
//}
//
//func GetAuthToken(userType string) (string, error) {
//	if userType != "client" && userType != "moderator" { //если тип != клиент или модератор
//		return "", fmt.Errorf("неверный тип пользователя: %s", userType) //то возвращаем и выводим ошибку
//	}
//	baseURL := "https://api.avito.ru/dummyLogin" //ручка
//	params := url.Values{}
//	params.Add("user_type", userType) //добавление пользователя
//	fullURL := baseURL + "?" + params.Encode()
//
//	resp, err := http.Get(fullURL)
//	if err != nil {
//		return "", fmt.Errorf("ошибка запроса: %v", err)
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		return "", fmt.Errorf("сервер вернул ошибку: %d", resp.StatusCode)
//	}
//	var data authResponse
//	if err = json.NewDecoder(resp.Body).Decode(&data); err != nil {
//		return "", fmt.Errorf("ошибка парсинга JSON: %v", err)
//	}
//	return data.Token, nil
//}
