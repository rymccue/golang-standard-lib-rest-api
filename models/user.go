package models

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type PrivateUserDetails struct {
	ID       int
	Password string
	Salt     string
}
