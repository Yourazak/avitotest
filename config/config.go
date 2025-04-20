package config

import (
	"avitotes/variable"
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
	AuthTokenModerator string
	AuthTokenClient    string
}

func NewConfig() *Config {
	err := godotenv.Load() //поиск env файла
	if err != nil {
		log.Println(variable.Msg_err_env)
		return nil
	}
	return &Config{
		Db: DbConfig{
			DsnDb: os.Getenv("DSN_DB"),
		},
		Auth: Authconfig{
			AuthTokenModerator: os.Getenv("TOKEN_MODETATOR"),
			AuthTokenClient:    os.Getenv("SECRET"),
		},
	}
}
