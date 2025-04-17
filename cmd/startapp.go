package main

import (
	"avitotes/config"
	"fmt"
)

func StartApp() {
	//подргужаем файл конфиг
	conf := config.NewConfig()
	fmt.Println(conf)
}
