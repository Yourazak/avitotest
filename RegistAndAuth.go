package main

type RegisterReuest struct {
	Email string `json:"email"`
	Password string `json:"password"`
	User_Type string `json:"user_type"`
}
type LoginRequest struct {

}