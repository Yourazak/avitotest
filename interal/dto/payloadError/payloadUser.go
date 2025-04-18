package payload

type UserCreateRequest struct {
	Email    string `json:"email" validate:"requierd,email"`
	Role     string `json:"role" validate:"requierd"`
	Password string `json:"password" validate:"min=8"`
}
