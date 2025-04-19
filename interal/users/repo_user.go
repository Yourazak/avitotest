package users

import (
	"avitotes/database"
	"fmt"
	"log"
)

type UserRepoDeps struct {
	Database *database.Db
}

type UserRepo struct {
	Database *database.Db
}

func NewUserRepo(database *database.Db) *UserRepo {
	return &UserRepo{
		Database: database,
	}
}

func (repo *UserRepo) Creat(user *User) (*User, error) {
	query := `INSERT INTO users(email,password,role) VALUES ($1,&2,&3) RETURNING id`
	err := repo.Database.MyDb.QueryRow(query, user.Email, user.Password, user.Role).Scan(&user.Id)
	if err != nil {
		log.Println("Create", err)
		return nil, fmt.Errorf("Ошибка при создании пользователя: %w", err)
	}
	return user, nil
}

func (repo *UserRepo) FindUserByEmailPass(email, password string) (*User, error) {

	//запрос
	query := `SELECT id,email,password, role FROM users where email = $1 and password = $2`
	//из бд дастаем данные и иницилизируем структуру
	var user User
	err := repo.Database.MyDb.QueryRow(query, email, password).Scan(&user.Id, &user.Email, &user.Password, &user.Role)
	if err != nil {
		log.Println("FindUserByEmailPass:", err)
		return nil, fmt.Errorf("Нет пользователя с логином и паролем %s %s", email, password)
	}
	return &user, nil
}
