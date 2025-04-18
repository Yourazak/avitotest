package server

import (
	"avitotes/config"
	"avitotes/interal/users"
	"net/http"
)

func ServerStart(conf *config.Config, resp *users.UserRepo) {

	//создаем пустой router
	router := http.NewServeMux()

	//подключаем сторонние service
	userService := users.NewUserService(resp)

	//подключаем userHandDependency для User
	usersHandDepend := users.UserHandDependency{
		userService,
		conf,
	}

	//Подключаем ручки USERS к router
	users.NewUserHandler(router, &userHandDepend)

	//Передаем server наши ручки
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//запускаем server
	server.ListenAndServe()
}
