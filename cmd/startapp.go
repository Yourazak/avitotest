package main

import (
	"avitotes/config"
	"avitotes/database"
	"avitotes/interal/pvz"
	"avitotes/interal/server"
	"avitotes/interal/users"
	"avitotes/pkg/repos"
	"log"
)

func StartApp() {
	//подргужаем файл config
	conf := config.NewConfig()

	//Подключаемся к базе данных
	db := CreateDb(conf)

	//создаем репозиторий на основе бд
	allRepos := CreateRepository(db)

	//запускаем сервер передавая туда репозиторий
	server.ServerStart(conf, userRepository)
}
	func CreateDb(conf *config.Config) *database.Db.DB {
		//подключение к базе данных
		db := database.NewDb(conf)

		//Создаем таблицу в бд если она не создана
		err := db.CreateTableUser()
		if err != nil {
				log.Fatal("Не удалось создать таблицу users:", err)
				return nil
		}


		//создаем таблицу в бд для PVZ если она не создана
		err := db.CreateTablePVZ()
		if err != nil{
			log.Fatal("Не удалось создать таблицу users:",err)
			return nil
		}
		return db
}

func CreateRepository(db *database.Db) *repos.AllRepository{
	//репозиторий для User
	userRepository := users.NewUserRepo(db)

	//репозиторий для PVZ
	pvzRepository := pvz.NewPVZRepo(db)

	return repos.NewAllRepository(userRepository,pvzRepository)
}