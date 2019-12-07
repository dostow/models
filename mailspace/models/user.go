package models

type User struct {
	Email     string `json:"Email"`
	CreatedAt string `json:"created_at"`
	ID        string `json:"id"`
	Username  string `json:"username"`
}
