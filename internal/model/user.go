package model

type User struct {
	ID       string
	Username string
	Password string
}

type UserRequest struct {
	Username string `json:"username"`
}

type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}
