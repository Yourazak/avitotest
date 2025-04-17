package oldtask

//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"net/http"
//	"time"
//)
//
//type PickupPoint struct {
//	ID        int       `json:"id"`
//	City      string    `json:"city"`
//	Address   string    `json:"address"`
//	CreatedAt time.Time `json:"created_at"`
//}
//type CreatePPRequest struct {
//	City    string `json:"city"`
//	Address string `json:"address"`
//}
//
//func CreatePickupPoint(token, city, address string) (*PickupPoint, error) {
//	//проврка токена == модератор
//
//	//RegisterUser(email, password, userType string) (string, error){
//	//	if userType != "moderator" {
//	//		return "", fmt.Errorf("неверный тип пользователя")
//	//	}
//	allowedCities := map[string]bool{
//		"Москва":          true,
//		"Санкт-Петербург": true,
//		"Казань":          true,
//	}
//	if !allowedCities[city] {
//		return nil, fmt.Errorf("город %s не поддерживается", city)
//	}
//
//	requestBody := CreatePPRequest{
//		City:    city,
//		Address: address,
//	}
//	jsonBody, _ := json.Marshal(requestBody)
//	req, _ := http.NewRequest("POST", "http://api.avito.ru/pickup-points", bytes.NewBuffer(jsonBody))
//	req.Header.Set("Authorization", "Bearer"+token)
//
//	client := &http.Client{}
//	resp, err := client.Do(req)
//}
//	if resp.StatusCode != htttp.StatusCreadet{
//		return nil,fmt.Errorf("ошибка создания %d",resp.StatusCode)
//}
//var pp PickupPoint
//if err := json.NewDecoder(resp.Body).Decode(&pp); err != nil{
//	return nil,err
//}
//return &pp,nil
//func main() {
//	token, _ := GetAuthToken("moderator")
//	pp, err := CreatePickupPoint(token, "Москва", "улица Пушкина,дом 10")
//	if err != nil {
//		fmt.Println("Ошибка", err)
//		return
//	}
//	fmt.Println("Создан ПВЗ ID %d: %s,%s\n", pp.ID, pp.City, pp.Address)
//}
//}
//
