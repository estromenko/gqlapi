package models

// UserWithToken ...
type UserWithToken struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
