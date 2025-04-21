package users

import (
	"avitotes/internal/dto/payload"
	"context"
)

type UserService interface {
	Register(ctx context.Context, email, password, role string) (*User, error)
	Login(ctx context.Context, email, password string) (*payload.TokenResponse, error)
	GetToken(role string) (string, error)
}
