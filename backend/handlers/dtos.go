package handlers

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
