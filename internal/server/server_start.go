package server

import (
	"avitotes/config"
	"avitotes/interal/pvz"
	"avitotes/interal/reception"
	"avitotes/interal/users"
	"net/http"
)

func ServerStart(conf *config.Config, resp *users.UserRepo) {

	//создаем пустой router
	router := http.NewServeMux()

	//подключаем к router ручки для User
	ConnectHandlerForUser(router, conf, reps.UserRepo)

	//подключаем к router ручки для PVZ
	ConnectHandlerForPvz(router, conf, reps.PvzRepo)

	//подключаем к ruoter ручки для Reception
	ConnectHandlerForReception(router, conf, reps.ReceptionRepo)

	//передаем server наши ручки
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	//запускаем server
	server.ListenAndServe()
}

func ConnectHandlerForReception(router *http.ServeMux, conf *config.Config, reps *pvzReceptionRepo) {
	//подключаем сторонние service
	receptionService := reception.NewReceptionService(reps, conf)

	//подключаем receptionHandDependency для Reception
	receptionHandDepend := reception.ReceptionHandlerDependency{
		receptionService,
		conf,
	}

	//подключаем ручки для Reception к router
	receptions.NewReceptionHandler(router, &receptionHandDepend)
}

func ConnectHandlerForPvz(router *http.ServeMux, conf *config.Config, reps *pvz.PVZRepo) {
	//подключаем сторонние service
	pvzService := pvzNewPvzService(reps, conf)

	//подлючаем pvzHandDependency для PVZ
	pvzHandDepend := pvz.HandlerDependency{
		pvzService,
		conf,
	}

	//подключаем ручки PVZ к router
	pvz.NewPvzHandler(router, &pvzHandDepend)
}

func ConnectHandlerForUser(router *http.ServeMux, conf *config.Config, reps *users.UserRepo) {
	//подключаем сторонние service
	userService := user.NewUserSrvice(reps, conf)

	//подключаем userHandlerDependency для User
	userHandDepend := users.UserHandlerDependency{
		userService,
		conf,
	}

	//подключаем ручки USERS к router
	users.NewUserHandler(router, &userHandDepend)
}
