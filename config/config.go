package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Db   DbConfig
	Auth Authconfig
}

type DbConfig struct {
	DsnDb string
}
type Authconfig struct {
	AuthToken string
}

func NewConfig() *Config {
	err := godotenv.Load() //поиск env файла
	if err != nil {
		log.Println("Неудачная загрузка env-файла")
		return nil
	}
	return &Config{
		Db: DbConfig{
			DsnDb: os.Getenv("DSN_DB"),
		},
		Auth: Authconfig{
			AuthToken: os.Getenv("SECRET"),
		},
	}
}
