package users

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Password string    `json:"password"`
}
