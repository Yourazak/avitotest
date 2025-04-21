package main

import (
	"avitotes/config"
	"avitotes/database"
	"avitotes/internal/pvz"
	reception "avitotes/internal/receptions"
	"avitotes/internal/server"
	"avitotes/internal/users"
	"avitotes/pkg/repos"
	"log"
)

func StartApp() {
	//подргужаем файл config
	conf := config.NewConfig()

	//работаем с базой данных
	db := CreateDb(conf)

	//создаем репозиторий на основе бд
	allRepos := CreateRepository(db)

	//запускаем сервер передавая туда репозиторий
	server.ServerStart(conf, allRepos)
}

func CreateDb(conf *config.Config) *database.Db {
	//подключение к базе данных
	db := database.NewDb(conf)

	//Создаем таблицу в бд если она не создана
	err := db.CreateTableUser()
	if err != nil {
		log.Fatal("Не удалось создать таблицу users:", err)
		return nil
	}
	//создаем таблицу в бд для PVZ если она не создана
	err = db.CreateTablePVZ()
	if err != nil {
		log.Fatal("Не удалось создать таблицу pvz:", err)
		return nil
	}
	//создаем таблицу в бд для Reception(приемки) если она не создана
	err = db.CreateTableReception()
	if err != nil {
		log.Fatal("Не удалось создать таблицу reception:", err)
		return nil
	}
	return db
}

func CreateRepository(db *database.Db) *repos.AllRepository {
	//репозиторий для User
	userRepository := users.NewUserRepo(db)

	//репозиторий для PVZ
	pvzRepository := pvz.NewPVZRepo(db)

	//репозиторий для Reception
	receptionRepository := reception.NewReceptionRepo(db)

	return repos.NewAllRepository(userRepository, pvzRepository, receptionRepository)
}
