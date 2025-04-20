package req

import (
	"encoding/json"
	"io"
	"net/http"
)

// функция которая проверяется корректность тела Request
func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Функция которая выводит в формем Json ответ в Response
func JsonResponse(w *http.ResponseWriter, data any) {
	(*w).Header().Set("Content-Type", "application/json")
	json.NewEncoder(*w).Encode(data)
}

func Decode[T any](body io.ReadCloser) (*T, error) {
	//создаем структуру куда будем класть наши данные
	var payload T
	//декодируем её из JSON --> в структуру
	err := json.NewDecoder(body).Decode(&payload)
	//если не удалось декодировать
	if err != nil {
		return nil, err
	}
	//если удалось
	return &payload, nil
}
